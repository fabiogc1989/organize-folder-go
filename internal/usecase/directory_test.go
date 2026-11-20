package usecase

import (
	"maps"
	"slices"
	"testing"
)

func TestCreateDefaultPolicyStructsByPath(t *testing.T) {
	dirPath := "C:\\temp\\test"

	arr, err := CreateDefaultPolicyStructsByPath(dirPath)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	folderNames := slices.Collect(maps.Keys(extensionsByDirectory))

	if len(arr) != len(extensionsByDirectory) {
		t.Errorf("expected %d policies, got %d", len(extensionsByDirectory), len(arr))
	}

	for _, policy := range arr {
		if !slices.Contains(folderNames, policy.Dir.Name) {
			t.Errorf("unexpected directory name found: %s", policy.Dir.Name)
		}
	}
}

func TestFindDirectoryByName(t *testing.T) {
	t.Run("should return true if the folder is found", func(t *testing.T) {
		dirPath := "C:\\temp\\test"

		arr, err := CreateDefaultPolicyStructsByPath(dirPath)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		policy, found := FindDirectoryByName(arr, slices.Collect(maps.Keys(extensionsByDirectory))[0])
		if found != true {
			t.Errorf("expected 'true', got %v", found)
		}

		if policy.Dir.Name != slices.Collect(maps.Keys(extensionsByDirectory))[0] {
			t.Errorf("expected dir name '%s', got '%s'", slices.Collect(maps.Keys(extensionsByDirectory))[0], policy.Dir.Name)
		}
	})

	t.Run("should retuern false if the folder name is not found", func(t *testing.T) {
		dirPath := "C:\\temp\\test"

		arr, err := CreateDefaultPolicyStructsByPath(dirPath)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		_, found := FindDirectoryByName(arr, "ABC")
		if found != false {
			t.Errorf("expected 'false', got %v", found)
		}
	})
}

func TestFindDirectoryByFileExtension(t *testing.T) {
	t.Run("should find the policy if the extension is found", func(t *testing.T) {
		dirPath := "C:\\temp\\test"

		arr, err := CreateDefaultPolicyStructsByPath(dirPath)
		if err != nil {
			t.Fatalf("expected no error creating policies, got: %v", err)
		}
		if len(arr) == 0 {
			t.Fatal("expected at least one policy struct, got none")
		}

		targetExt := arr[0].Dir.Name
		policy, found := FindDirectoryByFileExtension(arr, targetExt)
		if found != true {
			t.Fatalf("expected no error finding extension %q, got: %v", targetExt, err)
		}
		if policy.Dir.Name != targetExt {
			t.Errorf("expected directory/extension %q, got %q", targetExt, policy.Dir.Name)
		}
	})

	t.Run("should return false if the extension is not found", func(t *testing.T) {
		dirPath := "C:\\temp\\test"

		arr, err := CreateDefaultPolicyStructsByPath(dirPath)
		if err != nil {
			t.Fatalf("expected no error creating policies, got: %v", err)
		}
		if len(arr) == 0 {
			t.Fatal("expected at least one policy struct, got none")
		}

		targetExt := "Fab"
		_, found := FindDirectoryByFileExtension(arr, targetExt)
		if found != false {
			t.Errorf("expected 'false', got %v", found)
		}
	})
}

func TestIsDirectoryPath(t *testing.T) {
	t.Run("should return true if the input is a folder path", func(t *testing.T) {
		isDir := IsDirectoryPath(t.TempDir())
		if !isDir {
			t.Errorf("expected true, got %v", isDir)
		}
	})

	t.Run("should return false if the input is not a folder path", func(t *testing.T) {
		isDir := IsDirectoryPath("ptewer")
		if isDir != false {
			t.Errorf("expected value 'false', got %v", isDir)
		}
	})
}
