package domain

import (
	"errors"
	"slices"
	"testing"
)

func TestNewAllowedExtensionsPolicy(t *testing.T) {
	t.Run("Should create a new directory policy with valid inputs", func(t *testing.T) {
		dirPath := "C:\\temp\\Documents"
		extensions := []string{".docx", ".txt"}

		policy, err := NewAllowedExtensionsPolicy(dirPath, extensions)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		if policy.Dir.Path != dirPath {
			t.Errorf("expected folder path '%s', got '%s'", dirPath, policy.Dir.Path)
		}

		if len(policy.Extensions) != len(extensions) {
			t.Errorf("expected extensions length %d, got %d", len(extensions), len(policy.Extensions))
		}
	})

	t.Run("Should return an error when the extensions list is empty", func(t *testing.T) {
		dirPath := "C:\\temp\\Documents"
		extensions := []string{}

		_, err := NewAllowedExtensionsPolicy(dirPath, extensions)
		if !errors.Is(err, ErrEmptyExtensions) {
			t.Errorf("expected error ErrEmptyExtensions, got %v", err)
		}
	})
}

func TestIsAllowed(t *testing.T) {
	dirPath := "C:\\temp\\Documents"
	extensions := []string{".docx", ".txt"}

	policy, err := NewAllowedExtensionsPolicy(dirPath, extensions)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result := policy.IsAllowed(extensions[0]); result == false {
		t.Errorf("expected true, got %v", result)
	}
}

func TestAddExtension(t *testing.T) {
	dirPath := "C:\\temp\\Documents"
	extensions := []string{".docx", ".txt"}
	newExtension := ".ppt"

	policy, err := NewAllowedExtensionsPolicy(dirPath, extensions)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	policy.AddExtension(newExtension)
	if !slices.Contains(policy.Extensions, newExtension) {
		t.Errorf("expected extension '%s' in the list, got missing extension", newExtension)
	}
}
