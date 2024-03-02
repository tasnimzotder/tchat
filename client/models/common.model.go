package models

type Message struct {
	SenderID    string `json:"sender_id"`
	RecipientID string `json:"recipient_id"`
	MessageType string `json:"message_type"`
	Payload     string `json:"payload"`
	Timestamp   string `json:"timestamp"`
}

type Config struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Contact struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}
