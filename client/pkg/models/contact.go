package models

type Contact struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	PublicKey string `json:"public_key"`
}
