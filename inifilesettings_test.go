//go:build windows

package walk

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIniFileSettings_fileExists(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create an instance of IniFileSettings
	ifs := NewIniFileSettings("test.ini")
	ifs.SetPortable(true)

	// Case 1: File does not exist
	ifs.fileName = filepath.Join(tempDir, "does_not_exist.ini")
	exists, err := ifs.fileExists()
	if exists != false || err != nil {
		t.Errorf("Expected false, nil, got %v, %v", exists, err)
	}

	// Case 2: File exists
	ifs.fileName = filepath.Join(tempDir, "exists.ini")
	f, createErr := os.Create(ifs.fileName)
	if createErr != nil {
		t.Fatalf("Failed to create test file: %v", createErr)
	}
	f.Close()

	exists, err = ifs.fileExists()
	if exists != true || err != nil {
		t.Errorf("Expected true, nil, got %v, %v", exists, err)
	}

	// Case 3: Error other than IsNotExist
	// On Windows, if a file has an invalid path character, os.Stat returns an error
	// other than IsNotExist (e.g., ERROR_INVALID_NAME).
	// We'll simulate this by using an invalid filename.
	// We only check if an error is returned.
	ifs.fileName = filepath.Join(tempDir, "invalid\x00file.ini")
	exists, err = ifs.fileExists()
	if exists != false || err == nil {
		t.Errorf("Expected false, error, got %v, %v", exists, err)
	}
}
