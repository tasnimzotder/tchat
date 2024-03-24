package models

type Message struct {
	ID          uint   `json:"id" gorm:"primaryKey,autoIncrement"`
	SenderID    string `json:"sender_id"`
	RecipientID string `json:"recipient_id"`
	MessageType string `json:"message_type"`
	Payload     string `json:"payload"`
	Timestamp   string `json:"timestamp"`
	FileSize    int64  `json:"file_size"`
	FileName    string `json:"file_name"`
	FileMode    string `json:"file_mode"`
	FileExt     string `json:"file_ext"`
	CreatedAt   int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt   int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
