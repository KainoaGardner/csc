package user

import (
	"fmt"

	"github.com/KainoaGardner/csc/internal/auth"
	"github.com/KainoaGardner/csc/internal/config"
	"github.com/KainoaGardner/csc/internal/db"
	"github.com/KainoaGardner/csc/internal/types"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func SetupNewUser(postUser types.PostUser) (*types.User, error) {
	var result types.User

	result.Username = postUser.Username
	result.Email = postUser.Email
	result.CreatedTime = time.Now().UTC()

	passwordHash, err := auth.HashPassword(postUser.Password)
	if err != nil {
		return nil, err
	}
	result.PasswordHash = passwordHash

	return &result, nil
}

func SetupUserStats(newUser types.User) *types.UserStats {
	var result types.UserStats

	result.UserID = newUser.ID
	result.GameLogs = []string{}
	return &result
}

func CheckUniqueLogin(client *mongo.Client, dbConfig config.DB, user types.User) error {
	err := db.FindUserLogin(client, dbConfig, user)

	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		return err
	} else {
		return fmt.Errorf("Username or email already exists")
	}

}
