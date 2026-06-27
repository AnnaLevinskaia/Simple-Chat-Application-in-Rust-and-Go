# Rust Simple Chat Application

## Project Overview

This is the Rust version of the Simple Chat Application project. The purpose of this application is to simulate a local text-based chat system where multiple users can send messages, view chat history, search messages by keyword, and filter messages by user ID.

This application is stand-alone and runs in the terminal. It does not use a real server, network connection, or database. Messages are stored in memory while the program is running.

This Rust implementation is part of a language comparison project between **Rust** and **Go**.

---

## Features

The Rust chat application includes the following features:

* Simulates multiple users
* Sends messages from different users
* Stores message history
* Displays all chat messages with timestamps
* Searches messages by keyword
* Filters messages by user ID
* Handles empty message validation
* Demonstrates Rust modules and structured code
* Uses Rust ownership, borrowing, structs, enums, vectors, and error handling

---

## Rust Language Features Used

This project demonstrates several Rust-specific features:

### Structs

Structs are used to define the main data models:

* `User`
* `Message`
* `Chat`

### Enum

An enum is used to represent chat actions, such as:

* Send message
* Display history
* Search by keyword
* Filter by user

### Vector

The application uses `Vec<Message>` to store chat history in memory.

### Ownership and Borrowing

Rust’s ownership and borrowing system is used to safely pass user and message data between functions without causing memory issues.

### Result Error Handling

The application uses `Result<(), String>` to handle validation errors, such as empty messages.

### Threads and Channels

The application uses Rust threads and channels to simulate messages being sent from multiple users.

---

## Project Structure

```text
rust-simple-chat-app/
│
├── Cargo.toml
├── README.md
│
└── src/
    ├── main.rs
    ├── user.rs
    ├── message.rs
    ├── chat.rs
    └── search.rs
```

### File Descriptions

| File         | Description                                                         |
| ------------ | ------------------------------------------------------------------- |
| `main.rs`    | Runs the application and demonstrates the chat flow                 |
| `user.rs`    | Defines the `User` struct                                           |
| `message.rs` | Defines the `Message` struct and message display format             |
| `chat.rs`    | Handles sending messages, storing history, searching, and filtering |
| `search.rs`  | Contains helper functions for keyword search and user filtering     |
| `Cargo.toml` | Rust project configuration file and dependencies                    |

---

## Requirements

Before running this project, make sure Rust is installed.

To check if Rust is installed, run:

```bash
rustc --version
```

Also check Cargo:

```bash
cargo --version
```

If Rust is not installed, install it from the official Rust website.

---

## Dependency

This project uses the `chrono` crate for timestamps.

The dependency is listed in `Cargo.toml`:

```toml
[dependencies]
chrono = "0.4"
```

Important: the file must be named exactly:

```text
Cargo.toml
```

Rust will not recognize it if it is named `cargo.toml` because Linux is case-sensitive.

---

## How to Build and Run

From the shared project repository, go into the Rust application folder:

```bash
cd rust-simple-chat-app
```

Build the project:

```bash
cargo build
```

Run the project:

```bash
cargo run
```

---

## Sample Output

Example output may look like this:

```text
Simple Chat Application - Rust Version

Sending messages...

Message sent by Anna
Message sent by Dmitry
Message sent by Maria
Message sent by Anna
Message sent by Dmitry

Current action: DisplayHistory

--- Chat History ---
[2026-06-27 10:15:30] User 1 (Anna) says: Hello everyone!
[2026-06-27 10:15:30] User 2 (Dmitry) says: Hi Anna, this is the Rust chat app.
[2026-06-27 10:15:30] User 3 (Maria) says: Rust uses ownership and borrowing.
[2026-06-27 10:15:30] User 1 (Anna) says: We can search messages by keyword.
[2026-06-27 10:15:30] User 2 (Dmitry) says: We can also filter by user ID.

Total messages stored: 5

--- Search Results for keyword: Rust ---
[2026-06-27 10:15:30] User 2 (Dmitry) says: Hi Anna, this is the Rust chat app.
[2026-06-27 10:15:30] User 3 (Maria) says: Rust uses ownership and borrowing.

--- Filter Results for User ID: 2 ---
[2026-06-27 10:15:30] User 2 (Dmitry) says: Hi Anna, this is the Rust chat app.
[2026-06-27 10:15:30] User 2 (Dmitry) says: We can also filter by user ID.

--- Empty Message Test ---
Validation worked: Message cannot be empty.
```

The timestamps may be different each time because they are generated when the program runs.

---

## Testing

Basic testing was completed by running the application and checking the following scenarios:

| Test Case                    | Expected Result                            | Status |
| ---------------------------- | ------------------------------------------ | ------ |
| Multiple users send messages | Messages are stored in chat history        | Passed |
| Chat history is displayed    | All messages appear with timestamps        | Passed |
| Search by keyword            | Matching messages are displayed            | Passed |
| Filter by user ID            | Only messages from that user are displayed | Passed |
| Empty message test           | Application shows validation error         | Passed |

---

## Known Limitations

This is a simple local chat simulation. It does not include:

* Real-time network chat
* User login
* Database storage
* GUI interface
* Permanent message history after the program ends

These features were not included because the goal of the project is to compare programming language features, not to build a full production chat system.

---

## Conclusion

The Rust version of the Simple Chat Application demonstrates how Rust can be used to build a safe and organized terminal-based application. Rust’s ownership model, strong type system, structs, enums, vectors, modules, and error handling help make the application reliable and structured.

This implementation will be compared with the Go version to analyze differences in readability, writability, concurrency, memory management, error handling, and overall programming style.
