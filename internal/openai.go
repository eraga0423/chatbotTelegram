package internal

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	gpt "github.com/8ff/gpt/pkg/gpt_3_5_turbo"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func Chatgpt8ff(question string) string {
	token, err := ParseToken()
	if err != nil {
		slog.Error("Error ParseToken")
		return "error token"
	}
	log.Printf("Используем токен GPT: %s\n", token.GptTokenBaur)

	api, err := gpt.Init(gpt.Params{
		API_TOKEN:    token.GptTokenBaur,
		StripNewline: true,
		Request: gpt.ChatRequest{
			Model: "gpt-3.5-turbo",
		},
	},
	)
	if err != nil {
		slog.Error("Error GPT")
		return "error gpt"
	}
	choices, err := api.Query(question)
	if err != nil {
		slog.Error("ERROR QUERy")
		return fmt.Sprintf("error:%v Query my question : %s", err, question)
	}
	result := ""
	for _, choice := range choices {
		result = choice.Message.Content
	}

	return result
}

func ChatGemini(question string) (string, error) {
	token, err := ParseToken()
	if err != nil {
		slog.Error("Error ParseToken")
		return "", err
	}
	log.Printf("Используем токен gemini: %s\n", token.Gemini)
	// url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key=" + token.Gemini
	// model := GenerativeModel("gemini-1.5-flash")
	// resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a magic backpack."))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// printResponse(resp)
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {

		log.Fatal(err)
		return "", err
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	strResp, err := GeminiMesage(resp)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return strResp, nil

}
func GeminiMesage(resp *genai.GenerateContentResponse) (string, error) {
	var result strings.Builder
	for _, choice := range resp.Candidates {

		if choice.Content != nil {

			for _, part := range choice.Content.Parts {
				if text, ok := part.(genai.Text); ok {
					// fmt.Println(part, "paaaaaaaaaaaaaaaaaaart")
					result.WriteString(string(text))
					result.WriteString("\n")

				} else {
					return "", fmt.Errorf("error part")
				}
			}

		}

	}

	return result.String(), nil
}
