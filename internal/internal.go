package internal

import (
	"time"

	"github.com/mymmrac/telego"
)

func SendMesageToChat(bot *telego.Bot, chatId telego.ChatID) error {
	for {
		message := BoringMessage()
		_, err := bot.SendMessage(&telego.SendMessageParams{
			ChatID: chatId,
			Text:   message,
		})
		if err != nil {
			return err
		}
		now := time.Now()
		nextRun := time.Date(now.Year(), now.Month(), now.Day(), 15, 0, 0, 0, now.Location())
		if now.After(nextRun) {
			nextRun = nextRun.Add(24 * time.Hour)
		}
		sleepDuration := time.Until(nextRun)

		time.Sleep(sleepDuration)

	}
}
