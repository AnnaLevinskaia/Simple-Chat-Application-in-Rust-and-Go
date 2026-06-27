package main

import (
	"fmt"
	"sync"
	"time"
)

type ChatInput struct {
	User User
	Text string
}

func main() {
	fmt.Println("Simple Chat Application - Go Version")

	user1 := NewUser(1, "Anna")
	user2 := NewUser(2, "Dmitry")
	user3 := NewUser(3, "Maria")

	chat := NewChat()

	// Go-specific concurrency example using goroutines and channels.
	messageChannel := make(chan ChatInput)

	var waitGroup sync.WaitGroup

	sampleMessages := []ChatInput{
		{User: user1, Text: "Hello everyone!"},
		{User: user2, Text: "Hi Anna, this is the Go chat app."},
		{User: user3, Text: "Go uses goroutines and channels."},
		{User: user1, Text: "We can search messages by keyword."},
		{User: user2, Text: "We can also filter by user ID."},
	}

	fmt.Println("\nSending messages...\n")

	for _, input := range sampleMessages {
		waitGroup.Add(1)

		go func(chatInput ChatInput) {
			defer waitGroup.Done()

			time.Sleep(200 * time.Millisecond)
			messageChannel <- chatInput
		}(input)
	}

	go func() {
		waitGroup.Wait()
		close(messageChannel)
	}()

	for input := range messageChannel {
		err := chat.SendMessage(input.User, input.Text)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Message sent by", input.User.Name)
		}
	}

	chat.DisplayHistory()

	fmt.Println("\nTotal messages stored:", chat.MessageCount())

	fmt.Println("\n--- Search Results for keyword: Go ---")
	searchResults := chat.SearchByKeyword("Go")

	if len(searchResults) == 0 {
		fmt.Println("No messages found.")
	} else {
		PrintMessages(searchResults)
	}

	fmt.Println("\n--- Filter Results for User ID: 2 ---")
	filteredResults := chat.FilterByUser(2)

	if len(filteredResults) == 0 {
		fmt.Println("No messages found for this user.")
	} else {
		PrintMessages(filteredResults)
	}

	fmt.Println("\n--- Empty Message Test ---")
	err := chat.SendMessage(user1, "")

	if err != nil {
		fmt.Println("Validation worked:", err)
	} else {
		fmt.Println("Empty message was sent.")
	}
}

func PrintMessages(messages []Message) {
	for _, message := range messages {
		fmt.Printf(
			"[%s] User %d (%s) says: %s\n",
			message.Timestamp.Format("2006-01-02 15:04:05"),
			message.UserID,
			message.UserName,
			message.Text,
		)
	}
}