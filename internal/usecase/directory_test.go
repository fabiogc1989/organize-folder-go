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
