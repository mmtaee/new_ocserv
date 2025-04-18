package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type Config struct {
	Host         string
	Port         string
	SecretKey    string
	AllowOrigins []string
	Databases    string
}

var cfg *Config

func Init(debug bool) {
	if debug {
		err := godotenv.Load()
		if err != nil {
			log.Printf("Error loading .env file: %v", err)
		}
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "SECRET_KEY122456"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	allowOrigins := os.Getenv("ALLOW_ORIGINS")

	databases := os.Getenv("DATABASES")
	if databases == "" {
		databases = "ocserv.db"
	}

	cfg = &Config{
		Host:         host,
		Port:         port,
		SecretKey:    secretKey,
		AllowOrigins: strings.Split(allowOrigins, ","),
		Databases:    databases,
	}

	log.Println("config initialized")
}

func Get() *Config {
	return cfg
}
