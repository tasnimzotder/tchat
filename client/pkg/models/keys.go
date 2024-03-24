package models

type RSAKeys struct {
	// gorm.Model
	ID         uint64 `gorm:"primaryKey;autoIncrement:true" json:"id"`
	PrivateKey []byte `json:"private_key"`
	PublicKey  []byte `json:"public_key"`
}
