package main

import (
	"kanban-bot/infrastructure/config"
	"kanban-bot/infrastructure/flag"
	"kanban-bot/infrastructure/pg"
	"kanban-bot/infrastructure/zap"
	"kanban-bot/internal/adapter/logger"
)

func main() {
	// infrastructure
	configPath := flag.Parse()
	config := config.New(configPath)
	zap := zap.New(config.IsDebug)
	database := pg.New(config.DatabaseUrl, config.IsDebug)

	// adapter
	logger := logger.New(zap)

	// TODO: remove next line
	_, _ = logger, database
}
