# Simple Chat Application in Rust and Go

## Project Overview

This project implements the same Simple Chat Application in two programming languages: **Rust** and **Go**. The purpose of the project is to compare how both languages handle similar application requirements, including data structures, type systems, memory management, concurrency, error handling, modularity, readability, and programming style.

The application is a stand-alone terminal-based chat simulation. It does not use a real server, database, or network connection. Instead, it simulates multiple users sending messages locally.

## GitHub Repository

Repository Link:
https://github.com/AnnaLevinskaia/Simple-Chat-Application-in-Rust-and-Go

## Team Members

* Anna Levinskaia
* Jalshree Khanal

## Application Features

Both the Rust and Go versions include the same core features:

* Simulate multiple users
* Send messages
* Store message history
* Display messages with timestamps
* Search messages by keyword
* Filter messages by user ID
* Validate empty messages
* Demonstrate language-specific concurrency features

## Repository Structure

```text
Simple-Chat-Application-in-Rust-and-Go/
│
├── README.md
│
├── rust-simple-chat-app/
│   ├── Cargo.toml
│   ├── README.md
│   └── src/
│       ├── main.rs
│       ├── user.rs
│       ├── message.rs
│       ├── chat.rs
│       └── search.rs
│
└── go-simple-chat-app/
    ├── go.mod
    ├── README.md
    ├── main.go
    ├── user.go
    ├── message.go
    ├── chat.go
    └── search.go
```

## Rust Version

The Rust version demonstrates:

* Structs
* Enums
* Vectors
* Ownership and borrowing
* Result-based error handling
* Threads and channels

### Run Rust Application

```bash
cd rust-simple-chat-app
cargo run
```

## Go Version

The Go version demonstrates:

* Structs
* Slices
* Goroutines
* Channels
* Error returns
* WaitGroup synchronization

### Run Go Application

```bash
cd go-simple-chat-app
go run .
```

## Language Comparison Summary

Rust focuses more on memory safety, ownership, borrowing, and strict type checking. It requires more careful planning, especially when working with shared data and concurrency.

Go focuses more on simplicity, readability, and fast development. Goroutines and channels make concurrent programming easier to write and understand.

Both languages successfully implement the same chat application, but they use different programming styles and language features.

## Current Status

The application is complete in both Rust and Go. Both versions can send messages, store chat history, display timestamps, search by keyword, filter by user ID, and validate empty messages.

## Documentation

Each language folder includes its own README file with more detailed build and run instructions:

* `rust-simple-chat-app/README.md`
* `go-simple-chat-app/README.md`
