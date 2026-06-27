package main

import (
	"fmt"
	"sync"
	"time"
)

// ChatInput is a custom struct used to send chat data through a channel.
//
// It groups together:
// User -> the user who sends the message
// Text -> the message content
type ChatInput struct {
	User User
	Text string
}

func main() {
	fmt.Println("Simple Chat Application - Go Version")

	// Create three users for the chat application.
	user1 := NewUser(1, "Anna")
	user2 := NewUser(2, "Dmitry")
	user3 := NewUser(3, "Maria")

	chat := NewChat()

	// Create a Go channel.
	//
	// A channel is used for communication between goroutines.
	// In this program, goroutines will send ChatInput values into this channel.
	messageChannel := make(chan ChatInput)

	// Create a WaitGroup.
	//
	// A WaitGroup is used to wait for multiple goroutines to finish.
	var waitGroup sync.WaitGroup

	sampleMessages := []ChatInput{
		{User: user1, Text: "Hello everyone!"},
		{User: user2, Text: "Hi Anna, this is the Go chat app."},
		{User: user3, Text: "Go uses goroutines and channels."},
		{User: user1, Text: "We can search messages by keyword."},
		{User: user2, Text: "We can also filter by user ID."},
	}

	fmt.Println("\nSending messages...\n")

	// Loop through each sample message.
	for _, input := range sampleMessages {
		waitGroup.Add(1)


		// Start a new goroutine.
		//
		// A goroutine is a lightweight thread in Go.
		// It allows the program to run tasks concurrently.
		//
		// This simulates multiple users sending messages at the same time.
		go func(chatInput ChatInput) {
			defer waitGroup.Done()

			time.Sleep(200 * time.Millisecond)
			messageChannel <- chatInput
		}(input)
	}

	// Start another goroutine to close the channel.
	//
	// This goroutine waits until all message-sending goroutines finish.
	go func() {
		waitGroup.Wait()
		close(messageChannel)
	}()

	// Receive messages from the channel.
	//
	// This loop continues until the channel is closed.
	//
	// Each received value is a ChatInput object.
	for input := range messageChannel {
		err := chat.SendMessage(input.User, input.Text)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Message sent by", input.User.Name)
		}
	}

	// Display all messages stored in the chat history
	chat.DisplayHistory()

	// Print the total number of messages stored in the chat.
	fmt.Println("\nTotal messages stored:", chat.MessageCount())

	fmt.Println("\n--- Search Results for keyword: Go ---")
	searchResults := chat.SearchByKeyword("Go")

	if len(searchResults) == 0 {
		fmt.Println("No messages found.")
	} else {
		PrintMessages(searchResults)
	}

	// Filter messages by user ID.
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

// PrintMessages is a helper function.
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