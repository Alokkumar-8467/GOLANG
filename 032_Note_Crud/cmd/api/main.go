package main
import (
	"fmt"
	"log"
	"notes_api/internal/config"
	"notes_api/internal/db"
	"notes_api/internal/server"
)

// Now in tihs main.go we wire all the things means configure and connect each and every services.
// config -> db -> router -> run server
