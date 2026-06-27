use chrono::{DateTime, Local};

#[derive(Clone, Debug)]
pub struct Message {
    pub user_id: u32,
    pub user_name: String,
    pub text: String,
    pub timestamp: DateTime<Local>,
}

impl Message {
    pub fn new(user_id: u32, user_name: String, text: String) -> Self {
        Self {
            user_id,
            user_name,
            text,
            timestamp: Local::now(),
        }
    }

    pub fn display(&self) {
        println!(
            "[{}] User {} ({}) says: {}",
            self.timestamp.format("%Y-%m-%d %H:%M:%S"),
            self.user_id,
            self.user_name,
            self.text
        );
    }
}