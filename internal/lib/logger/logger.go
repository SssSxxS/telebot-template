package logger

import (
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(logsDirPath string) error {
	/* -------------------------------- Log File -------------------------------- */
	if err := os.MkdirAll(logsDirPath, 0755); err != nil {
		return err
	}
	logFile, err := os.OpenFile(
		filepath.Join(logsDirPath, time.Now().Format("2006-01-02")+".log"),
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		0664,
	)
	if err != nil {
		return err
	}

	/* ----------------------------- Console Output ----------------------------- */
	zerolog.TimeFieldFormat = time.RFC3339Nano
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05.000", NoColor: false}

	/* -------------------------- Logger Initialization ------------------------- */
	multi := zerolog.MultiLevelWriter(consoleWriter, logFile)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	return nil
}
