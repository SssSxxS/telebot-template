package keyboards

import (
	tele "gopkg.in/telebot.v4"
	"strconv"
)

var (
	HelpPrevBtn = tele.InlineButton{
		Unique: "help_prev",
		Text:   "« Back",
	}

	HelpNextBtn = tele.InlineButton{
		Unique: "help_next",
		Text:   "Next »",
	}

	HelpCloseBtn = tele.InlineButton{
		Unique: "help_close",
		Text:   "Close",
	}
)

func GetHelpKeyboard(pageIndex int, totalPages int) *tele.ReplyMarkup {
	// Validate page index
	if pageIndex < 0 {
		pageIndex = 0
	} else if pageIndex >= totalPages {
		pageIndex = totalPages - 1
	}

	// Create keyboard
	kb := &tele.ReplyMarkup{}
	var navRow []tele.InlineButton

	// Prev button - always show, with cyclic navigation
	prevBtn := HelpPrevBtn
	if pageIndex > 0 {
		prevBtn.Data = strconv.Itoa(pageIndex - 1)
	} else {
		// If on first page, go to last page
		prevBtn.Data = strconv.Itoa(totalPages - 1)
	}
	navRow = append(navRow, prevBtn)

	// Page number
	pageBtn := tele.InlineButton{
		Text:   strconv.Itoa(pageIndex+1) + " / " + strconv.Itoa(totalPages),
		Unique: "help_page",
	}
	navRow = append(navRow, pageBtn)

	// Next button
	nextBtn := HelpNextBtn
	if pageIndex < totalPages-1 {
		nextBtn.Data = strconv.Itoa(pageIndex + 1)
	} else {
		// If on last page, go to first page
		nextBtn.Data = strconv.Itoa(0)
	}
	navRow = append(navRow, nextBtn)

	// Add navigation row and close button
	kb.InlineKeyboard = append(kb.InlineKeyboard, navRow)
	kb.InlineKeyboard = append(kb.InlineKeyboard, []tele.InlineButton{HelpCloseBtn})

	return kb
}
