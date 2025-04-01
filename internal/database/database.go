package database

import (
	"os"
	"path/filepath"

	"telebot-template/internal/database/models"
	"telebot-template/internal/database/repositories"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(dbDirPath string) error {
	/* ---------------------------------- Paths --------------------------------- */
	if err := os.MkdirAll(dbDirPath, 0700); err != nil {
		return err
	}
	dbFilePath := filepath.Join(dbDirPath, "telebot-template.db")

	/* -------------------------------- Database -------------------------------- */
	db, err := gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{
		Logger: logger.New(&log.Logger, logger.Config{
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
		}),
	})
	if err != nil {
		return err
	}

	/* ------------------------------- Migrations ------------------------------- */
	db.AutoMigrate(&models.User{})

	/* ------------------------------ Repositories ------------------------------ */
	repositories.InitUserRepo(db)

	/* -------------------------------------------------------------------------- */
	return nil
}
