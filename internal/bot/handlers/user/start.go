package handlers

import (
	// "telebot-template/internal/database/models"
	// r "telebot-template/internal/database/repositories"

	"telebot-template/internal/bot/keyboards"

	tele "gopkg.in/telebot.v4"
)

func RegisterStartHandlers(b *tele.Group) {
	b.Handle("/start", handleStart)
}

func handleStart(c tele.Context) error {

	// Start logic here

	// sender := c.Sender()
	// user, err := r.UserRepo.GetByTelegramID(c.Sender().ID)
	// if err != nil {
	// 	return err
	// }

	kb := keyboards.GetStartKeyboard()

	return c.Send("Start", kb)
}
