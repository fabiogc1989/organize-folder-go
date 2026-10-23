package domain

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestNewDirectory(t *testing.T) {
	t.Run("should create a directory successfully with valid input", func(t *testing.T) {
		dir, err := NewDirectory("MyFolder", "/test/path", 0755)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		if dir.Name != "MyFolder" {
			t.Errorf("expected name 'MyFolder', got '%s'", dir.Name)
		}

		if dir.Perm != 0755 {
			t.Errorf("expected permission 0755, got %v", dir.Perm)
		}
	})

	t.Run("should return an error if name is empty", func(t *testing.T) {
		_, err := NewDirectory("    ", "/test/path", 0755)
		if !errors.Is(err, ErrInvalidName) {
			t.Errorf("expected error ErrInvalidName, got: %v", err)
		}
	})

	t.Run("should assign default permission if perm is 0", func(t *testing.T) {
		dir, err := NewDirectory("Folder", "/test/path", 0)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		if dir.Perm != 0755 {
			t.Errorf("expected permission 0755, got %v", dir.Perm)
		}
	})
}

func TestDirectoryDiskOperations(t *testing.T) {
	// Create an isolated temporary directory for this test
	tmpDir := t.TempDir()

	subDirPath := filepath.Join(tmpDir, "test_folder")

	// 1. Instantiating Directory with a temporary path
	dir, err := NewDirectory("test", subDirPath, 0755)
	if err != nil {
		t.Fatalf("error creating Directory instance: %v", err)
	}

	// 2. Test Exists before creation (should return false)
	if dir.Exists() {
		t.Errorf("expected directory to not exist yet")
	}

	// 3. Test Create
	err = dir.Create()
	if err != nil {
		t.Errorf("error creating directory on disk: %v", err)
	}

	//4. Test Exists after creation (should return true)
	if !dir.Exists() {
		t.Errorf("expected directory to exist after calling Create()")
	}

	// 5. Test GetFiles on an empty folder
	files, err := dir.GetFiles()
	if err != nil {
		t.Fatalf("error getting files: %v", err)
	}
	if len(files) != 0 {
		t.Errorf("expected 0 files, got %d", len(files))
	}

	// 6. Create a dummy file inside the temporary directory to test GetFiles with contents
	filePath := filepath.Join(subDirPath, "test_file.txt")
	err = os.WriteFile(filePath, []byte("hello world test"), 0644)
	if err != nil {
		t.Fatalf("error creating test file: %v", err)
	}

	// 7. Test GetFiles again (should return 1 file)
	files, err = dir.GetFiles()
	if err != nil {
		t.Fatalf("error fetching files after file creation: %v", err)
	}

	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(files))
	}

	if files[0].Name != "test_file.txt" {
		t.Errorf("incorrect file name: got '%s'", files[0].Name)
	}
	if files[0].Extension != ".txt" {
		t.Errorf("incorrect file extension: expected '.txt', got '%s'", files[0].Extension)
	}
	if files[0].Size != 16 { // "hello world test" = 16 bytes
		t.Errorf("incorrect file size: expect 16, got %d", files[0].Size)
	}
}
