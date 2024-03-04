package config

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvVariable(t *testing.T) {
	envVariables := map[string]string{
		"TEST":  "test",
		"TEST2": "test2",
		"TEST3": "",
	}

	for key, value := range envVariables {
		t.Run(key, func(t *testing.T) {
			err := os.Setenv("TC_"+key, value)
			if err != nil {
				log.Printf("Error setting environment variable: %v", err)
				return
			}

			result := GetEnvVariable(key)
			assert.Equal(t, value, result)
		})
	}
}
