package models

type Tokens struct {
	TelegramBotToken string `json:"telegram_bot_token"`
	GptTokenEraga    string `json:"gpt_token_eraga"`
	GptTokenBaur     string `json:"gpt_token_baur"`
	Gemini           string `json:"gemini_token"`
}
