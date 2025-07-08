package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GameLog struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	GameID  string             `bson:"gameID" json:"gameID"`
	WhiteID string             `bson:"whiteID" json:"whiteID"`
	BlackID string             `bson:"blackID" json:"blackID"`

	Date time.Time `bson:"date" json:"date"`

	MoveCount   int      `bson:"moveCount" json:"moveCount"`
	Moves       []string `bson:"moves" json:"moves"`
	BoardStates []string `bson:"boardStates" json:"boardStates"`

	Winner *int   `bson:"winner" json:"winner"`
	Reason string `bson:"reason" json:"reason"`
}
