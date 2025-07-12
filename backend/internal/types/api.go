package types

import (
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
	Turn int    `json:"turn"`
}

type PostMoveResponse struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FEN  string             `json:"fen"`
	Move string             `json:"move"`
}

type PostPlace struct {
	Position string `json:"position"`
	Type     int    `json:"type"`
	Turn     int    `json:"turn"` //TEMP USE SENT ID INSTEAD
}

type DeletePlace struct {
	Position string `json:"position"`
	Turn     int    `json:"turn"` //TEMP USE SENT ID INSTEAD
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
}

type PostGameResponse struct {
	ID        string `json:"_id"`
	WhiteID   string `json:"whiteID"`
	BlackID   string `json:"blackID"`
	Color     string `json:"color"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Money     [2]int `json:"money"`
	State     int    `json:"state`
	PlaceLine int    `json:"placeLine"`
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
}

type PostState struct {
	State int `json:"state"`
}

type PostReady struct {
	Turn  int  `json:"turn"`
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
	Turn int  `json:"turn"`
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

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"tokenToken"`
}
