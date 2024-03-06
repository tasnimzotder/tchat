package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigDir(t *testing.T) {
	homeDir, _ := os.UserHomeDir()
	expectedDir := homeDir + "/.config/tchat"

	// Call the getConfigDir function
	actualDir := getConfigDir()

	// Assert that the returned directory matches the expected directory
	assert.Equal(t, expectedDir, actualDir)
}

func TestGetConfigDirError(t *testing.T) {
	expectedDir := ""

	// Call the getConfigDir function
	actualDir := getConfigDir()

	// Assert that the returned directory matches the expected directory
	assert.NotEqual(t, expectedDir, actualDir)
}

func TestGetFileContents(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	// Write test data to the file
	data := []byte("test data")
	_, err = tmpfile.Write(data)
	if err != nil {
		t.Fatalf("Failed to write test data to file: %v", err)
	}

	// Call the GetFileContents function
	contents, err := GetFileContents(tmpfile.Name())
	if err != nil {
		t.Fatalf("GetFileContents failed: %v", err)
	}

	// Assert that the returned contents match the test data
	assert.Equal(t, data, contents)
}

func TestSaveFile(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	// Test data
	data := []byte("test data")

	// Call the SaveFile function
	err = SaveFile(tmpfile.Name(), data)
	if err != nil {
		t.Fatalf("SaveFile failed: %v", err)
	}

	// Read the saved file
	savedData, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to read saved file: %v", err)
	}

	// Assert that the saved data matches the test data
	assert.Equal(t, data, savedData)
}
