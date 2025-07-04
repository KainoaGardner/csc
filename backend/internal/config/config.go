package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB         DB
	PublicHost string
	Port       string
}

type DB struct {
	Name        string
	Uri         string
	Collections Collections
}

type Collections struct {
	Users     string
	UserStats string
	Games     string
	GameLogs  string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func checkGetenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatal("Env variable not found: ", val)
	}
	return val
}

func LoadConfig() Config {
	var result Config

	result.PublicHost = checkGetenv("PUBLIC_HOST")
	result.Port = checkGetenv("PORT")

	result.DB.Name = os.Getenv("MONGODB_DB")
	result.DB.Uri = checkGetenv("MONGODB_URI")

	result.DB.Collections.Users = checkGetenv("MONGODB_USERS_COLLECTION")
	result.DB.Collections.UserStats = checkGetenv("MONGODB_USER_STATS_COLLECTION")
	result.DB.Collections.Games = checkGetenv("MONGODB_GAMES_COLLECTION")
	result.DB.Collections.GameLogs = checkGetenv("MONGODB_GAME_LOGS_COLLECTION")

	return result
}
