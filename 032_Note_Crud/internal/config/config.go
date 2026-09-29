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
func Load() (Config, error) {

	// godotenv.Load() reads .env and sets them into the process env
	// os.getenv -> reads those values
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("Failed to load .env")
	}
	mongoURI, err := extractEnv("MONGO_URI")
	if err != nil {
		return Config{}, err
	}
