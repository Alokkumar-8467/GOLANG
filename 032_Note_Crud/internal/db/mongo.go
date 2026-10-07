package db
import (
	"context"
	"notes_api/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)


/*
cfg config.Config means the function takes one parameter named cfg, whose type is Config from your config package. It's usually a struct holding settings like the MongoDB URI and the database name. So cfg is the settings box, and config.Config is its type.

The return values are the three things you listed:
*mongo.Client is the connection to the MongoDB server
*mongo.Database is the specific database you'll work in
error is nil if everything worked, or the problem if it didn't
