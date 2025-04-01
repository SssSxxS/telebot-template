package handlers

import (
	tele "gopkg.in/telebot.v4"
)

func RegisterHelpHandlers(b *tele.Group) {
	b.Handle("/help", handleHelp)
}

func handleHelp(c tele.Context) error {
	return c.Send("Help")
}
