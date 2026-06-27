package main

import "strings"

func ContainsKeyword(message Message, keyword string) bool {
	messageText := strings.ToLower(message.Text)
	searchKeyword := strings.ToLower(keyword)

	return strings.Contains(messageText, searchKeyword)
}

func BelongsToUser(message Message, userID int) bool {
	return message.UserID == userID
}