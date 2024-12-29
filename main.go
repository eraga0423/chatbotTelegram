package main

import (
	"chatbot/internal"
	"fmt"
	"log"
	"os"

	"github.com/mymmrac/telego"
)

func main() {
	tokens, err := internal.ParseToken()
	bot, err := telego.NewBot(tokens.TelegramBotToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer bot.StopLongPolling()
	// var IdChatClastation int64
	// IdChatClastation = -1002396731326
	// idChat := telego.ChatID{
	// 	ID: IdChatClastation,
	// }

	// err = SendMesageToChat(bot, idChat)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	for update := range updates {
		question := "Салем!"
		if update.Message != nil {
			log.Printf("Чат ID: %d, Сообщение: %s", update.Message.Chat.ID, update.Message.Text)
			question = update.Message.Text
			fmt.Printf("четфффффффффффам: %s", question)
		}
		// answer := internal.Chatgpt8ff(question)
		answer, _ := internal.ChatGemini(question)
		_, err := bot.SendMessage(&telego.SendMessageParams{
			ChatID: telego.ChatID{
				ID: update.Message.Chat.ID,
			},
			Text: answer,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

}
