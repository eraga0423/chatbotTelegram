package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := "7261966141:AAGw7-WCqfMpHpkXZEgDNBGvX-bjN7JdpcI"
	bot, err := telego.NewBot(botToken, telego.WithDefaultLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	updates, _ := bot.UpdatesViaLongPollin(nil)

	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatID := tu.ID(update.Message.Chat.ID)
			keyboard := tu.Keyboard(
				tu.KeyboardRow(
					tu.KeyboardButton("БАТЫРМА"),
				),
			)
		().WithReplyMarkup(keyboard)
			_, _ = bot.CopyMessage(
				tu.CopyMessage(
					chatID,
					chatID,
					update.Message.MessageID,
				),
			)

		}
	}
}
