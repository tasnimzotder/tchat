package utils

import "github.com/google/uuid"

func GenerateUUID() (string, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return newUUID.String(), nil
}
