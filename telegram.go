package utils

import (
	"github.com/0xVanfer/types"
	"github.com/imroc/req"
)

// Can only send string text.
func SendTelegramSimple(botToken string, chatId int64, text string) error {
	url := `https://api.telegram.org/bot` + botToken + `/sendMessage?chat_id=` + types.ToString(chatId) + `&text=` + text
	_, err := req.Get(url)
	return err
}
