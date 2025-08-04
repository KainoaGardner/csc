package websockets

import (
	"encoding/json"
	"fmt"
	"github.com/KainoaGardner/csc/internal/config"
	"github.com/KainoaGardner/csc/internal/db"
	"github.com/KainoaGardner/csc/internal/engine"
	"github.com/KainoaGardner/csc/internal/types"
	"github.com/KainoaGardner/csc/internal/utils"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sync"
)

type GameRoom struct {
	Players map[string]*websocket.Conn //userID -> gameRoom
	Mutex   sync.Mutex
}

var GameConnections = make(map[string]*GameRoom) //gameID -> gameRoom
var GameConnectionsMutex sync.Mutex

func AddPlayerToGame(gameID string, playerID string, conn *websocket.Conn) {
	GameConnectionsMutex.Lock()
	room, ok := GameConnections[gameID]
	if !ok {
		room = &GameRoom{
			Players: make(map[string]*websocket.Conn),
		}
		GameConnections[gameID] = room
	}
	GameConnectionsMutex.Unlock()

	room.Mutex.Lock()
	room.Players[playerID] = conn
	room.Mutex.Unlock()

	log.Printf("Player %s connected to game %s", playerID, gameID)
}

func RemovePlayerFromGame(gameID string, playerID string) {
	GameConnectionsMutex.Lock()
	room, ok := GameConnections[gameID]
	GameConnectionsMutex.Unlock()
	if !ok {
		return
	}

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	conn, ok := room.Players[playerID]
	if ok {
		conn.Close()
		delete(room.Players, playerID)
	}

	if len(room.Players) == 0 {
		GameConnectionsMutex.Lock()
		delete(GameConnections, gameID)
		GameConnectionsMutex.Unlock()
	}
}

func BroadcastToGame(gameID string, msg interface{}) {
	GameConnectionsMutex.Lock()
	room, ok := GameConnections[gameID]
	GameConnectionsMutex.Unlock()
	if !ok {
		return
	}

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	for playerID, conn := range room.Players {
		err := conn.WriteJSON(msg)
		if err != nil {
			RemovePlayerFromGame(gameID, playerID)
		}
	}
}

func BroadcastToPlayer(gameID string, playerID string, msg interface{}) {
	GameConnectionsMutex.Lock()
	room, ok := GameConnections[gameID]
	GameConnectionsMutex.Unlock()
	if !ok {
		return
	}

	room.Mutex.Lock()
	defer room.Mutex.Unlock()

	conn, ok := room.Players[playerID]
	if !ok {
		return
	}

	err := conn.WriteJSON(msg)
	if err != nil {
		RemovePlayerFromGame(gameID, playerID)
	}
}

func DeleteGameRoom(gameID string) {
	GameConnectionsMutex.Lock()
	room, ok := GameConnections[gameID]
	GameConnectionsMutex.Unlock()
	if !ok {
		return
	}

	room.Mutex.Lock()
	defer room.Mutex.Unlock()
	for playerID, _ := range room.Players {
		RemovePlayerFromGame(gameID, playerID)
	}
}

func HandleMessages(gameID string, playerID string, conn *websocket.Conn, client *mongo.Client, config config.Config) {
	defer func() {
		log.Printf("Closing connection for player %s", playerID)
		RemovePlayerFromGame(gameID, playerID)
		resignCase(gameID, playerID, client, config)
		// DeleteGameRoom(gameID)
		conn.Close()
	}()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Printf("read error (player=%s game=%s): %v", playerID, gameID, err)
			resignCase(gameID, playerID, client, config)
			return
		}

		var msg types.IncomingMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			log.Printf("bad json (player=%s): %v data=%q", playerID, err, data)
			continue
		}

		var over bool
		switch msg.Type {
		case "join":
			joinCase(gameID, playerID, client, config)
		case "move":
			over = moveCase(gameID, playerID, msg, client, config)
		case "place":
			placeCase(gameID, playerID, msg, client, config)
		case "ready":
			over = readyCase(gameID, playerID, msg, client, config)
		case "draw":
			over = drawCase(gameID, playerID, msg, client, config)
		case "resign":
			over = resignCase(gameID, playerID, client, config)
		default:
		}

		if over {
			return
		}
	}
}

func broadcastError(gameID string, playerID string, err error) {
	fmt.Println("Error", err)
	data := types.Error{
		Error: err.Error(),
	}

	response := types.OutgoingMessage{
		Type: "error",
		Data: data,
	}

	BroadcastToPlayer(gameID, playerID, response)
}

func joinCase(gameID string, playerID string, client *mongo.Client, config config.Config) {
	game, err := engine.JoinGameCase(gameID, playerID, client, config)
	if err != nil {
		broadcastError(gameID, playerID, err)
		return
	}

	response := types.OutgoingMessage{
		Type: "player Joined",
	}
	BroadcastToGame(gameID, response)

	if game.State == types.PlaceState {
		data := types.PostGameResponse{
			ID:        gameID,
			WhiteID:   game.WhiteID,
			BlackID:   game.BlackID,
			Width:     game.Board.Width,
			Height:    game.Board.Height,
			Money:     game.Money,
			StartTime: game.Time,
			PlaceLine: game.Board.PlaceLine,
			State:     game.State,
		}

		response := types.OutgoingMessage{
			Type: "start",
			Data: data,
		}
		BroadcastToGame(gameID, response)
	}
}

func moveCase(gameID string, playerID string, msg types.IncomingMessage, client *mongo.Client, config config.Config) bool {
	postMove, err := utils.ParseMsgJSON[types.PostMove](msg)
	if err != nil {
		broadcastError(gameID, playerID, err)
		return false
	}

	game, fen, err := engine.MoveCase(gameID, playerID, postMove, client, config)
	if err != nil {
		broadcastError(gameID, playerID, err)
		return false
	}

	if game.State == types.OverState {
		return gameOver(game, gameID, playerID, client, config)

	} else {
		data := types.PostMoveResponse{
			ID:   game.ID,
			FEN:  fen,
			Move: postMove.Move,
		}

		response := types.OutgoingMessage{
			Type: "move",
			Data: data,
		}
		BroadcastToGame(gameID, response)

	}

	return false
}

func placeCase(gameID string, playerID string, msg types.IncomingMessage, client *mongo.Client, config config.Config) {
	postPlace, err := utils.ParseMsgJSON[types.PostPlace](msg)
	if err != nil {
		broadcastError(gameID, playerID, err)
		return
	}

	_, data, err := engine.PlaceCase(gameID, playerID, postPlace, client, config)
	if err != nil {
		broadcastError(gameID, playerID, err)
		return
	}

	response := types.OutgoingMessage{
		Type: "place",
		Data: data,
	}
	BroadcastToGame(gameID, response)

}

func readyCase(gameID string, playerID string, msg types.IncomingMessage, client *mongo.Client, config config.Config) bool {
	postReady, err := utils.ParseMsgJSON[types.PostReady](msg)
	if err != nil {
		broadcastError(gameID, playerID, err)
		return false
	}

	game, err := engine.ReadyCase(gameID, playerID, postReady, client, config)
	if err != nil {
		broadcastError(gameID, playerID, err)
		return false
	}

	if game.State == types.MoveState {
		gameLog := engine.SetupGameLog(*game)
		_, err := db.CreateGameLog(client, config.DB, gameLog)
		if err != nil {
			broadcastError(gameID, playerID, err)
			return false
		}

		data := map[string]interface{}{
			"_id":   game.ID,
			"state": game.State,
			"ready": game.Ready,
		}
		response := types.OutgoingMessage{
			Type: "move",
			Data: data,
		}
		BroadcastToGame(gameID, response)

	} else if game.State == types.OverState {
		gameLog := engine.SetupGameLog(*game)
		_, err := db.CreateGameLog(client, config.DB, gameLog)
		if err != nil {
			broadcastError(gameID, playerID, err)
			return false
		}

		return gameOver(game, gameID, playerID, client, config)
	} else {
		data := map[string]interface{}{
			"_id":   game.ID,
			"state": game.State,
			"ready": game.Ready,
		}

		response := types.OutgoingMessage{
			Type: "ready",
			Data: data,
		}
		BroadcastToGame(gameID, response)
	}
	return false
}

func drawCase(gameID string, playerID string, msg types.IncomingMessage, client *mongo.Client, config config.Config) bool {
	postDraw, err := utils.ParseMsgJSON[types.PostDrawRequest](msg)
	if err != nil {
		broadcastError(gameID, playerID, err)
		return false
	}

	game, err := engine.DrawCase(gameID, playerID, postDraw, client, config)
	if err != nil {
		broadcastError(gameID, playerID, err)
		return false
	}

	if game.State == types.OverState {
		return gameOver(game, gameID, playerID, client, config)

	} else {
		data := map[string]interface{}{
			"_id":  game.ID,
			"draw": game.Draw,
		}

		response := types.OutgoingMessage{
			Type: "draw",
			Data: data,
		}
		BroadcastToGame(gameID, response)
	}

	return false
}

func resignCase(gameID string, playerID string, client *mongo.Client, config config.Config) bool {
	game, err := engine.ResignCase(gameID, playerID, client, config)
	if err != nil {
		broadcastError(gameID, playerID, err)
		return false
	}

	log.Println("Resign Case")

	if game.State == types.OverState {
		return gameOver(game, gameID, playerID, client, config)
	}

	return false
}

func gameOver(game *types.Game, gameID string, playerID string, client *mongo.Client, config config.Config) bool {
	gameLog, err := db.FindGameLogFromGameID(client, config.DB, gameID)
	if err != nil {
		broadcastError(gameID, playerID, err)
		return false
	}
	engine.SetupFinalGameLog(*game, gameLog)
	err = db.GameLogFinalUpdate(client, config.DB, gameID, *gameLog)
	if err != nil {
		broadcastError(gameID, playerID, err)
		return false
	}

	_, err = db.DeleteGame(client, config.DB, gameID)
	if err != nil {
		broadcastError(gameID, playerID, err)
		return false
	}

	data := types.GameOverResponse{
		ID:            game.ID,
		WhiteID:       game.WhiteID,
		BlackID:       game.BlackID,
		MoveCount:     game.MoveCount,
		HalfMoveCount: game.HalfMoveCount,
		Winner:        game.Winner,
		Reason:        game.Reason,
		State:         game.State,
		LastMoveTime:  game.LastMoveTime,
	}

	response := types.OutgoingMessage{
		Type: "over",
		Data: data,
	}
	BroadcastToGame(gameID, response)

	return true
}
