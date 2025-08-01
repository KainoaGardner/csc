package types

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type APIRespone struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// game api
type PostMove struct {
	Move string `json:"move"`
}

type PostMoveResponse struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FEN  string             `json:"fen"`
	Move string             `json:"move"`
}

type PostPlace struct {
	Position string `json:"position"`
	Type     int    `json:"type"`
	Place    bool   `json:"place"`
}

type DeletePlace struct {
	Position string `json:"position"`
}

type PlaceResponse struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FEN      string             `json:"fen"`
	Position string             `json:"position"`
	Type     int                `json:"type"`
	Cost     int                `json:"cost"`
	Money    [2]int             `json:"money"`
}

type PostGame struct {
	Width     int      `json:"width"`
	Height    int      `json:"height"`
	Money     [2]int   `json:"money"`
	StartTime [2]int64 `json:"startTime"`
	PlaceLine int      `json:"placeLine"`
	Public    bool     `json:"public"`
}

type PostGameResponse struct {
	ID        string   `json:"_id"`
	WhiteID   string   `json:"whiteID"`
	BlackID   string   `json:"blackID"`
	Color     string   `json:"color"`
	Width     int      `json:"width"`
	Height    int      `json:"height"`
	Money     [2]int   `json:"money"`
	StartTime [2]int64 `json:"startTime"`
	State     int      `json:"state"`
	PlaceLine int      `json:"placeLine"`
}

type GetGameResponse struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	WhiteID       string             `bson:"whiteID" json:"whiteID"`
	BlackID       string             `bson:"blackID" json:"blackID"`
	Turn          int                `bson:"turn" json:"turn"`
	MoveCount     int                `bson:"moveCount" json:"moveCount"`
	HalfMoveCount int                `bson:"halfMoveCount" json:"halfMoveCount"`
	Winner        *int               `bson:"winner" json:"winner"`
	Reason        string             `bson:"reason" json:"reason"`
	State         int                `bson:"state" json:"state"` //0 place,1 move,2 over
	Time          [2]int64           `bson:"time" json:"time"`
	LastMoveTime  time.Time          `bson:"lastMoveTime" json:"lastMoveTime"`
	Money         [2]int             `bson:"money" json:"money"`
	Ready         [2]bool            `bson:"ready" json:"ready"`
	Draw          [2]bool            `bson:"draw" json:"draw"`
	Public        bool               `json:"public"`
}

type PostState struct {
	State int `json:"state"`
}

type PostReady struct {
	Ready bool `json:"ready"`
}

type GameOverResponse struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	WhiteID       string             `bson:"whiteID" json:"whiteID"`
	BlackID       string             `bson:"blackID" json:"blackID"`
	MoveCount     int                `bson:"moveCount" json:"moveCount"`
	HalfMoveCount int                `bson:"halfMoveCount" json:"halfMoveCount"`
	Winner        *int               `bson:"winner" json:"winner"`
	Reason        string             `bson:"reason" json:"reason"`
	State         int                `bson:"state" json:"state"` //0 place,1 move,2 over
	LastMoveTime  time.Time          `bson:"lastMoveTime" json:"lastMoveTime"`
}

type PostDrawRequest struct {
	Draw bool `json:"draw"`
}

//user api

type PostUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID          primitive.ObjectID `json:"_id,omitempty"`
	Username    string             `json:"username"`
	Email       string             `json:"email"`
	CreatedTime time.Time          `json:"createdTime"`
}

type UpdateUserStats struct {
	GamesPlayed int    `bson:"gamesPlayed"`
	GamesWon    int    `bson:"gamesWon"`
	GameLog     string `bson:"gameLog"`
}

type PostLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID          primitive.ObjectID `json:"_id"`
	AccessToken string             `json:"accessToken"`
}

type TokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type PostForgotPassword struct {
	Email string `json:"email"`
}

type PostResetPassword struct {
	Password string `json:"password"`
}

type UserStatsResponse struct {
	GamesPlayed int      `json:"gamesPlayed"`
	GamesWon    int      `json:"gamesWon"`
	GameLog     []string `json:"gameLog"`
}

type JoinableGameResponse struct {
	ID        primitive.ObjectID `json:"_id"`
	Width     int                `json:"width"`
	Height    int                `json:"height"`
	PlaceLine int                `json:"placeLine"`
	WhiteID   string             `json:"whiteID"`
	Time      [2]int64           `json:"time"`
	Money     [2]int             `json:"money"`
}

type IncomingMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type OutgoingMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type Error struct {
	Error string `json:"error"`
}
