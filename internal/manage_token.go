package internal

import (
	"chatbot/models"
	"encoding/json"
	"fmt"
	"os"
)

func ParseToken() (models.Tokens, error) {
	file, err := os.Open("utils.json")
	if err != nil {
		fmt.Println(err)
		return models.Tokens{}, err
	}
	defer file.Close()
	var tokens models.Tokens
	err = json.NewDecoder(file).Decode(&tokens)
	if err != nil {
		fmt.Println(err)
		return models.Tokens{}, err
	}

	return tokens, nil
}
