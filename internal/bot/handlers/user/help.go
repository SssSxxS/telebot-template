package handlers

import (
	"strconv"

	"telebot-template/internal/bot/keyboards"

	tele "gopkg.in/telebot.v4"
)

type HelpData struct {
	Title   string
	Content string
}

var helpPages = []HelpData{
	{
		Title:   "Basic Information",
		Content: "Welcome to the bot help! Here you will find information about available commands and functions.",
	},
	{
		Title:   "Commands",
		Content: "/start - start working with the bot\n/help - show help",
	},
	{
		Title:   "Navigation",
		Content: "Use the menu buttons to navigate through the bot. To return to the main menu, use the 'Back' button.",
	},
	{
		Title:   "Support",
		Content: "If you have any questions or problems, please contact the administrator.",
	},
}

func RegisterHelpHandlers(b *tele.Group) {
	b.Handle("/help", handleHelp)
	b.Handle(&keyboards.HelpBtn, handleHelp)
	b.Handle(&keyboards.HelpPrevBtn, handleHelpPagination)
	b.Handle(&keyboards.HelpNextBtn, handleHelpPagination)
	b.Handle(&keyboards.HelpCloseBtn, handleHelpClose)
}

func handleHelp(c tele.Context) error {
	return showHelpPage(c, 0)
}

func handleHelpPagination(c tele.Context) error {
	pageStr := c.Callback().Data
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}

	return showHelpPage(c, page)
}

func handleHelpClose(c tele.Context) error {
	return c.Delete()
}

func showHelpPage(c tele.Context, pageIndex int) error {
	page := getHelpPage(pageIndex)
	message := "<b>" + page.Title + "</b>\n\n" + page.Content
	totalPages := len(helpPages)

	kb := keyboards.GetHelpKeyboard(pageIndex, totalPages)

	if c.Callback() != nil {
		return c.Edit(message, kb)
	}
	return c.Send(message, kb)
}

func getHelpPage(pageIndex int) HelpData {
	totalPages := len(helpPages)

	if pageIndex < 0 {
		pageIndex = 0
	} else if pageIndex >= totalPages {
		pageIndex = totalPages - 1
	}

	return helpPages[pageIndex]
}
