use crate::message::Message;
use crate::search::{belongs_to_user, contains_keyword};
use crate::user::User;

// Enum shows Rust-specific design for chat actions.
#[derive(Debug)]
pub enum ChatAction {
    SendMessage,
    DisplayHistory,
    SearchByKeyword,
    FilterByUser,
}

pub struct Chat {
    messages: Vec<Message>,
}

impl Chat {
    pub fn new() -> Self {
        Self {
            messages: Vec::new(),
        }
    }

    pub fn send_message(&mut self, user: &User, text: &str) -> Result<(), String> {
        if text.trim().is_empty() {
            return Err("Message cannot be empty.".to_string());
        }

        let message = Message::new(user.id, user.name.clone(), text.to_string());
        self.messages.push(message);

        Ok(())
    }

    pub fn display_history(&self) {
        if self.messages.is_empty() {
            println!("No messages in chat history.");
            return;
        }

        println!("\n--- Chat History ---");
        for message in &self.messages {
            message.display();
        }
    }

    pub fn search_by_keyword(&self, keyword: &str) -> Vec<&Message> {
        self.messages
            .iter()
            .filter(|message| contains_keyword(message, keyword))
            .collect()
    }

    pub fn filter_by_user(&self, user_id: u32) -> Vec<&Message> {
        self.messages
            .iter()
            .filter(|message| belongs_to_user(message, user_id))
            .collect()
    }

    pub fn message_count(&self) -> usize {
        self.messages.len()
    }
}