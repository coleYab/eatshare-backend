package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port          string
	Enviroment    string
	DbConnString  string
	JWTSecret     string
	GmailPassword string
	GmailAddress  string
	JWTExpiration time.Duration
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}

func New() *Config {
	expiresIn := getEnv("JWT_EXPIRATION", "5")
	expiresInCasted, _ := strconv.Atoi(expiresIn)
	return &Config{
		Port:          fmt.Sprintf(":%v", getEnv("PORT", "8080")),
		Enviroment:    getEnv("ENV", "DEVELOPMENT"),
		DbConnString:  getEnv("DB_CONN_URI", ""),
		GmailPassword: getEnv("GMAIL_PASSWORD", ""),
		GmailAddress:  getEnv("GMAIL_ADDRESS", ""),
		JWTSecret:     getEnv("JWT_SECRET", "My awsome secret and it has to be too long"),
		JWTExpiration: time.Hour * 24 * time.Duration(expiresInCasted),
	}
}

var Cfg = New()
