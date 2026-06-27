use crate::message::Message;

pub fn contains_keyword(message: &Message, keyword: &str) -> bool {
    message
        .text
        .to_lowercase()
        .contains(&keyword.to_lowercase())
}

pub fn belongs_to_user(message: &Message, user_id: u32) -> bool {
    message.user_id == user_id
}