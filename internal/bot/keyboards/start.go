package keyboards

import (
	tele "gopkg.in/telebot.v4"
)

var (
	HelpBtn = tele.ReplyButton{Text: "ℹ Help"}
)

func GetStartKeyboard() *tele.ReplyMarkup {
	kb := &tele.ReplyMarkup{ResizeKeyboard: true}
	kb.ReplyKeyboard = append(kb.ReplyKeyboard, []tele.ReplyButton{HelpBtn})

	return kb
}
