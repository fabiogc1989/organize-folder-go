package domain

import (
	"errors"
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
