package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	MongoURI   string
	MongoDB    string
	ServerPort string
}
