# Go Simple Chat Application

## Project Overview

This is the Go version of the Simple Chat Application project. The application is a stand-alone, text-based chat program that simulates multiple users sending messages to each other.

The purpose of this project is to compare how the same application can be implemented in **Go** and **Rust**. This Go implementation focuses on simple syntax, readable code, goroutines, channels, slices, structs, and basic error handling.

The application runs in the terminal. It does not use a real network, server, database, or GUI. Messages are stored in memory while the program is running.

---

## Features

The Go chat application includes the following features:

* Simulates multiple users
* Sends messages from different users
* Stores message history
* Displays chat messages with timestamps
* Searches messages by keyword
* Filters messages by user ID
* Handles empty message validation
* Uses goroutines to simulate multiple users sending messages
* Uses channels for message passing
* Uses slices to store message history

---

## Go Language Features Used

### Structs

Structs are used to define the main data models:

* `User`
* `Message`
* `Chat`
* `ChatInput`

Example:

```go
type User struct {
	ID   int
	Name string
}
```

### Slices

The application uses a slice of messages to store chat history:

```go
History []Message
```

Slices are useful in Go because they can grow dynamically as new messages are added.

### Goroutines

Goroutines are used to simulate users sending messages at the same time:

```go
go func(chatInput ChatInput) {
	messageChannel <- chatInput
}(input)
```

This demonstrates Go’s built-in support for concurrency.

### Channels

Channels are used to pass messages from simulated users into the chat handler:

```go
messageChannel := make(chan ChatInput)
```

This allows messages to be sent between goroutines safely.

### Error Handling

The application uses Go’s standard error handling style. If a message is empty, the function returns an error:

```go
return errors.New("message cannot be empty")
```

---

## Project Structure

```text
go-simple-chat-app/
│
├── main.go
├── user.go
├── message.go
├── chat.go
├── search.go
├── go.mod
└── README.md
```

### File Descriptions

| File         | Description                                                                             |
| ------------ | --------------------------------------------------------------------------------------- |
| `main.go`    | Runs the application and demonstrates the chat flow                                     |
| `user.go`    | Defines the `User` struct and user creation function                                    |
| `message.go` | Defines the `Message` struct and message creation function                              |
| `chat.go`    | Handles sending messages, storing history, displaying history, searching, and filtering |
| `search.go`  | Contains helper functions for keyword search and filtering by user                      |
| `go.mod`     | Go module file                                                                          |
| `README.md`  | Project documentation and run instructions                                              |

---

## Requirements

Before running this project, make sure Go is installed.

Check Go version:

```bash
go version
```

If Go is not installed on Ubuntu/Linux, install it with:

```bash
sudo apt update
sudo apt install golang-go
```

Then check again:

```bash
go version
```

---

## How to Run the Application

From the shared project repository, go into the Go application folder:

```bash
cd go-simple-chat-app
```

If `go.mod` does not exist yet, create it:

```bash
go mod init go-simple-chat-app
```

Run the application:

```bash
go run .
```

---

## Sample Output

Example output may look like this:

```text
Simple Chat Application - Go Version

Sending messages...

Message sent by Anna
Message sent by Dmitry
Message sent by Maria
Message sent by Anna
Message sent by Dmitry

--- Chat History ---
[2026-06-27 10:15:30] User 1 (Anna) says: Hello everyone!
[2026-06-27 10:15:30] User 2 (Dmitry) says: Hi Anna, this is the Go chat app.
[2026-06-27 10:15:30] User 3 (Maria) says: Go uses goroutines and channels.
[2026-06-27 10:15:30] User 1 (Anna) says: We can search messages by keyword.
[2026-06-27 10:15:30] User 2 (Dmitry) says: We can also filter by user ID.

Total messages stored: 5

--- Search Results for keyword: Go ---
[2026-06-27 10:15:30] User 2 (Dmitry) says: Hi Anna, this is the Go chat app.
[2026-06-27 10:15:30] User 3 (Maria) says: Go uses goroutines and channels.

--- Filter Results for User ID: 2 ---
[2026-06-27 10:15:30] User 2 (Dmitry) says: Hi Anna, this is the Go chat app.
[2026-06-27 10:15:30] User 2 (Dmitry) says: We can also filter by user ID.

--- Empty Message Test ---
Validation worked: message cannot be empty
```

The exact message order may be slightly different because the program uses goroutines to simulate concurrent message sending.

---

## Testing

Basic testing was completed by running the program and checking the main requirements.

| Test Case                      | Expected Result                                | Status |
| ------------------------------ | ---------------------------------------------- | ------ |
| Multiple users send messages   | Messages are added to chat history             | Passed |
| Chat history displays messages | All messages show with timestamps              | Passed |
| Search by keyword              | Matching messages are displayed                | Passed |
| Filter by user ID              | Only messages from selected user are displayed | Passed |
| Empty message validation       | Error message is displayed                     | Passed |

---

## Known Limitations

This is a simple local chat simulation. It does not include:

* Real-time network communication
* User login or authentication
* Database storage
* Permanent message history
* GUI interface
* Real client-server chat functionality

These features were not included because the main goal of the project is to compare programming language features between Go and Rust.

---

## Conclusion

The Go version of the Simple Chat Application demonstrates how Go can be used to build a simple and readable terminal-based application. Go’s structs, slices, goroutines, channels, and error handling make the application easy to understand and quick to develop.

This implementation will be compared with the Rust version to analyze differences in readability, writability, concurrency, memory management, error handling, and overall programming style.
