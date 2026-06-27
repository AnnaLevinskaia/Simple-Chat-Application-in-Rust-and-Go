mod chat;
mod message;
mod search;
mod user;

use chat::{Chat, ChatAction};
use std::sync::mpsc;
use std::thread;
use std::time::Duration;
use user::User;

fn main() {
    println!("Simple Chat Application - Rust Version");

    let user1 = User::new(1, "Anna");
    let user2 = User::new(2, "Dmitry");
    let user3 = User::new(3, "Maria");

    let mut chat = Chat::new();

    // Rust-specific concurrency example using threads and channels.
    let (sender, receiver) = mpsc::channel();

    let users_and_messages = vec![
        (user1.clone(), "Hello everyone!"),
        (user2.clone(), "Hi Anna, this is the Rust chat app."),
        (user3.clone(), "Rust uses ownership and borrowing."),
        (user1.clone(), "We can search messages by keyword."),
        (user2.clone(), "We can also filter by user ID."),
    ];

    for (user, text) in users_and_messages {
        let sender_clone = sender.clone();

        thread::spawn(move || {
            thread::sleep(Duration::from_millis(200));
            sender_clone.send((user, text.to_string())).unwrap();
        });
    }

    drop(sender);

    println!("\nSending messages...\n");

    for (user, text) in receiver {
        match chat.send_message(&user, &text) {
            Ok(_) => println!("Message sent by {}", user.name),
            Err(error) => println!("Error: {}", error),
        }
    }

    let current_action = ChatAction::DisplayHistory;
    println!("\nCurrent action: {:?}\n", current_action);

    chat.display_history();

    println!("\nTotal messages stored: {}", chat.message_count());

    println!("\n--- Search Results for keyword: Rust ---");
    let search_results = chat.search_by_keyword("Rust");

    if search_results.is_empty() {
        println!("No messages found.");
    } else {
        for message in search_results {
            message.display();
        }
    }

    println!("\n--- Filter Results for User ID: 2 ---");
    let filtered_results = chat.filter_by_user(2);

    if filtered_results.is_empty() {
        println!("No messages found for this user.");
    } else {
        for message in filtered_results {
            message.display();
        }
    }

    println!("\n--- Empty Message Test ---");
    match chat.send_message(&user1, "") {
        Ok(_) => println!("Empty message was sent."),
        Err(error) => println!("Validation worked: {}", error),
    }
}