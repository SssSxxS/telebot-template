package main

import (
	"path/filepath"

	"telebot-template/internal/bot"
	"telebot-template/internal/database"
	"telebot-template/internal/lib/logger"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	logsDirPath := filepath.Join("data", "logs")
	if err := logger.Init(logsDirPath); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize logger")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("Failed to load .env file")
	}

	dbDirPath := filepath.Join("data", "db")
	if err := database.Init(dbDirPath); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database")
	}

	bot.Run()
}
