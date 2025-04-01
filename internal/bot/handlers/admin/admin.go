package handlers

import (
	tele "gopkg.in/telebot.v4"
)

func RegisterAdminHandlers(b *tele.Group) {
	b.Handle("/admin", handleAdmin)
}

func handleAdmin(c tele.Context) error {
	return c.Send("Admin")
}
