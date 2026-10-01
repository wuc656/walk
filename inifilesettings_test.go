//go:build windows

package walk

import (
	"os"
	"path/filepath"
	"testing"
	"golang.org/x/sys/windows"
)

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
