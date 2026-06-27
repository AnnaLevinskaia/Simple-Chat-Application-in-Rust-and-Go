package main

import "time"

type Message struct {
	UserID    int
	UserName  string
	Text      string
	Timestamp time.Time
}

func NewMessage(user User, text string) Message {
	return Message{
		UserID:    user.ID,
		UserName:  user.Name,
		Text:      text,
		Timestamp: time.Now(),
	}
}