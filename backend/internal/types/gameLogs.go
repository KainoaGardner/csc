package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GameLog struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	GameID  primitive.ObjectID `bson:"gameID" json:"gameID"`
	WhiteID string             `bson:"whiteID" json:"whiteID"`
	BlackID string             `bson:"blackID" json:"blackID"`

	Date time.Time `bson:"date" json:"date"`

	MoveCount      int      `bson:"moveCount" json:"moveCount"`
	Moves          []string `bson:"moves" json:"moves"`
	BoardStates    []string `bson:"boardStates" json:"boardStates"`
	BoardHeight    int      `bson:"boardHeight" json:"boardHeight"`
	BoardWidth     int      `bson:"boardWidth" json:"boardWidth"`
	BoardPlaceLine int      `bson:"boardPlaceLine" json:"boardPlaceLine"`

	Winner *int   `bson:"winner" json:"winner"`
	Reason string `bson:"reason" json:"reason"`
}
