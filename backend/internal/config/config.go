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
	JWT        JWT
	Email      EMAIL
}

type EMAIL struct {
	From     string
	Password string
}

type DB struct {
	Name        string
	Uri         string
	Collections Collections
}

type JWT struct {
	AccessKey          string
	RefreshKey         string
	PasswordRefreshKey string
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

	result.JWT.AccessKey = checkGetenv("JWT_ACCESS_KEY")
	result.JWT.RefreshKey = checkGetenv("JWT_REFRESH_KEY")
	result.JWT.PasswordRefreshKey = checkGetenv("JWT_PASSWORD_REFRESH_KEY")

	result.DB.Name = os.Getenv("MONGODB_DB")
	result.DB.Uri = checkGetenv("MONGODB_URI")

	result.DB.Collections.Users = checkGetenv("MONGODB_USERS_COLLECTION")
	result.DB.Collections.UserStats = checkGetenv("MONGODB_USER_STATS_COLLECTION")
	result.DB.Collections.Games = checkGetenv("MONGODB_GAMES_COLLECTION")
	result.DB.Collections.GameLogs = checkGetenv("MONGODB_GAME_LOGS_COLLECTION")

	result.Email.Password = checkGetenv("EMAIL_APP_PASSWORD")
	result.Email.From = checkGetenv("EMAIL_FROM")

	return result
}
