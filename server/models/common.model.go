package models

type Message struct {
	SenderID    string `json:"sender_id"`
	RecipientID string `json:"recipient_id"`
	MessageType string `json:"message_type"`
	Payload     string `json:"payload"`
	Timestamp   string `json:"timestamp"`
	FileSize    int64  `json:"file_size"`
	FileName    string `json:"file_name"`
	FileMode    string `json:"file_mode"`
	FileExt     string `json:"file_ext"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Connection struct {
	Name        string `json:"name"`
	UserID      string `json:"user_id"`
	PassKey     string `json:"pass_key"`
	PublicKey   string `json:"public_key"`
	Expiration  string `json:"expiration"`
	Persistence string `json:"persistence"`
}
