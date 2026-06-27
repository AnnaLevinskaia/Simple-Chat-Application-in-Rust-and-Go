package main

import (
	"errors"
	"fmt"
	"strings"
)

type Chat struct {
	History []Message
}

func NewChat() Chat {
	return Chat{
		History: []Message{},
	}
}

func (chat *Chat) SendMessage(user User, text string) error {
	if strings.TrimSpace(text) == "" {
		return errors.New("message cannot be empty")
	}

	message := NewMessage(user, text)
	chat.History = append(chat.History, message)

	return nil
}

func (chat Chat) DisplayHistory() {
	if len(chat.History) == 0 {
		fmt.Println("No messages in chat history.")
		return
	}

	fmt.Println("\n--- Chat History ---")

	for _, message := range chat.History {
		fmt.Printf(
			"[%s] User %d (%s) says: %s\n",
			message.Timestamp.Format("2006-01-02 15:04:05"),
			message.UserID,
			message.UserName,
			message.Text,
		)
	}
}

func (chat Chat) SearchByKeyword(keyword string) []Message {
	results := []Message{}

	for _, message := range chat.History {
		if ContainsKeyword(message, keyword) {
			results = append(results, message)
		}
	}

	return results
}

func (chat Chat) FilterByUser(userID int) []Message {
	results := []Message{}

	for _, message := range chat.History {
		if BelongsToUser(message, userID) {
			results = append(results, message)
		}
	}

	return results
}

func (chat Chat) MessageCount() int {
	return len(chat.History)
}