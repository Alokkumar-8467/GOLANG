package db
import (
	"context"
	"notes_api/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)
/*
cfg config.Config means the function takes one parameter named cfg, whose type is Config from your confi
