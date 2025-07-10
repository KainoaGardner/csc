package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Username      string             `bson:"username"`
	Email         string             `bson:"email"`
	EmailVerified bool               `bson:"emailVerified"`
	EmailToken    string             `bson:"emailToken"`
	PasswordHash  string             `bson:"passwordHash"`
	CreatedTime   time.Time          `bson:"createdTime"`
	Admin         bool               `bson:"admin"`
}

type UserStats struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"userID"`
	GamesPlayed int                `bson:"gamesPlayed"`
	GamesWon    int                `bson:"gamesWon"`
	GameLogs    []string           `bson:"gameLogs"`
}
