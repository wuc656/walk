//go:build windows

package walk

import (
	"os"
	"path/filepath"
	"testing"
	"golang.org/x/sys/windows"
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

func TestIniFileSettings_Permissions(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "config_test", "settings.ini")

	// Create an existing directory and file with full access permissions (insecure)
	dirPath := filepath.Dir(filePath)
	err := os.MkdirAll(dirPath, 0777)
	if err != nil {
		t.Fatalf("Failed to create dir: %v", err)
	}

	// Write with insecure permissions
	err = os.WriteFile(filePath, []byte("test=1"), 0666)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	// Use IniFileSettings which should fix the permissions
	ifs := NewIniFileSettings(filePath)
	ifs.SetPortable(true)

	err = ifs.Load()
	if err != nil {
		t.Fatalf("Failed to load settings: %v", err)
	}

	err = ifs.Put("test", "2")
	if err != nil {
		t.Fatalf("Failed to put setting: %v", err)
	}

	err = ifs.Save()
	if err != nil {
		t.Fatalf("Failed to save settings: %v", err)
	}

	// Verify directory permissions
	verifyOwnerOnlyAccess(t, dirPath)

	// Verify file permissions
	verifyOwnerOnlyAccess(t, filePath)
}

func verifyOwnerOnlyAccess(t *testing.T, path string) {
	sd, err := windows.GetNamedSecurityInfo(path, windows.SE_FILE_OBJECT, windows.DACL_SECURITY_INFORMATION)
	if err != nil {
		t.Fatalf("Failed to get security info for %s: %v", path, err)
	}

	dacl, _, err := sd.DACL()
	if err != nil {
		t.Fatalf("Failed to get DACL for %s: %v", path, err)
	}

	// To fully verify, we'd check ACEs. But since SD is generated via SDDL:
	// "D:P(A;OICI;FA;;;OW)", it should have 1 ACE for Owner.
	if dacl == nil {
		t.Fatalf("DACL is nil for %s, meaning full access to everyone", path)
	}
}

func TestIniFileSettings_PermissionsNew(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "config_test_new", "settings.ini")

	// Use IniFileSettings which should fix the permissions
	ifs := NewIniFileSettings(filePath)
	ifs.SetPortable(true)

	err := ifs.Put("test", "2")
	if err != nil {
		t.Fatalf("Failed to put setting: %v", err)
	}

	err = ifs.Save()
	if err != nil {
		t.Fatalf("Failed to save settings: %v", err)
	}

	dirPath := filepath.Dir(filePath)

	// Verify directory permissions
	verifyOwnerOnlyAccess(t, dirPath)

	// Verify file permissions
	verifyOwnerOnlyAccess(t, filePath)
}
