package domain

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewFileInfo(t *testing.T) {
	t.Run("should create a file info successfuly with valid input", func(t *testing.T) {
		tmpDir := t.TempDir()
		filePath := filepath.Join(tmpDir, "test.txt")
		os.WriteFile(filePath, []byte("hello world test"), 0644)

		fileInfo, err := NewFileInfo(filePath)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		if fileInfo.Name != "test.txt" {
			t.Errorf("incorrect file name: expected 'test.txt', got '%s'", fileInfo.Name)
		}
		if fileInfo.FullName != filePath {
			t.Errorf("incoorect file path: expected '/home/test/Downloads/test.txt', got '%s'", fileInfo.FullName)
		}
		if fileInfo.Size != 16 {
			t.Errorf("incorrect file size: expected 16, got %d", fileInfo.Size)
		}
		if fileInfo.Extension != ".txt" {
			t.Errorf("incorrect file extension: expected '.txt', got '%s'", fileInfo.Extension)
		}
	})

	t.Run("shoudl return an error if file path is invalid", func(t *testing.T) {
		_, err := NewFileInfo("")
		if err == nil {
			t.Errorf("expected an error for invalid file path, got nil")
		}
	})
}

func TestFileInfo_Move(t *testing.T) {
	tmpDir := t.TempDir()
	srcDir := filepath.Join(tmpDir, "source")
	destDir := filepath.Join(tmpDir, "destination")

	os.Mkdir(srcDir, 0755)
	os.Mkdir(destDir, 0755)

	filePath := filepath.Join(srcDir, "test.txt")
	os.WriteFile(filePath, []byte("move test"), 0644)

	fileInfo, err := NewFileInfo(filePath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = fileInfo.Move(destDir)
	if err != nil {
		t.Fatalf("expected no error moving file, got: %v", err)
	}

	expectedNewPath := filepath.Join(destDir, "test.txt")
	if fileInfo.FullName != expectedNewPath {
		t.Errorf("expected updated FullName to be '%s', got '%s'", expectedNewPath, fileInfo.FullName)
	}

	if _, err := os.Stat(expectedNewPath); os.IsNotExist(err) {
		t.Errorf("expected file to exist at destination path")
	}
}
