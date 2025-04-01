package bot

import (
	"os"
	"time"

	adminHandlers "telebot-template/internal/bot/handlers/admin"
	userHandlers "telebot-template/internal/bot/handlers/user"
	customMiddleware "telebot-template/internal/bot/middleware"

	"github.com/rs/zerolog/log"
	tele "gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/middleware"
)

func Run() {
	/* ------------------------------ Bot Settings ------------------------------ */
	pref := tele.Settings{
		Token: os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second,
			AllowedUpdates: []string{"message", "callback_query"}},
		ParseMode: tele.ModeHTML,
		OnError: func(err error, c tele.Context) {
			if c != nil {
				logEvent := log.Error().Err(err)
				if c.Sender() != nil {
					logEvent = logEvent.
						Int64("telegram_id", c.Sender().ID).
						Str("username", c.Sender().Username)
				}
				if c.Callback() != nil {
					logEvent = logEvent.
						Str("callback_unique", c.Callback().Unique).
						Str("callback_data", c.Callback().Data)
				} else if c.Message() != nil {
					logEvent = logEvent.Str("text", c.Text())
				}
				logEvent.Msg("Telegram bot error")
			} else {
				log.Error().Err(err).Msg("Telegram bot error without context")
			}
		},
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize bot")
	}

	/* --------------------------------- Groups --------------------------------- */
	userGroup := b.Group()
	adminGroup := b.Group()

	/* ------------------------------- Middlewares ------------------------------ */
	b.Use(middleware.AutoRespond())
	b.Use(middleware.IgnoreVia())
	b.Use(customMiddleware.IgnoreOld(time.Minute))
	b.Use(customMiddleware.UserTracker())

	adminGroup.Use(customMiddleware.IsAdmin())

	/* -------------------------------- Handlers -------------------------------- */
	userHandlers.RegisterStartHandlers(userGroup)
	userHandlers.RegisterHelpHandlers(userGroup)

	adminHandlers.RegisterAdminHandlers(adminGroup)

	/* -------------------------------------------------------------------------- */
	log.Info().Msg("TELEGRAM BOT STARTED")
	b.Start()
}
