package utils

import (
	"github.com/0xVanfer/types"
	"github.com/imroc/req"
)

// Send simple telegram message.
//
// Param:
//
//	botToken: Telegram bot token, you can search for
//			  `BotFather` in telegram and generate your own bot.
//	chatId:   If you are sending a message to a user, chatId is his id, a possitive integer,
//		      like 12345678, which the user can get by starting `userinfobot`.
//	          If you are sending a message to a group, chatId is a negative integer.
//
// NOTE:
//
//	Can only send simple text. Not support `\` like `\n`.
func SendTelegramSimple(botToken string, chatId int64, text string) error {
	url := `https://api.telegram.org/bot` + botToken + `/sendMessage?chat_id=` + types.ToString(chatId) + `&text=` + text
	_, err := req.Get(url)
	return err
}
