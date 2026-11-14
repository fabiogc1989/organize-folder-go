package usecase

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/fabiogc1989/organize-folder/internal/domain"
)

var (
	extensionsByDirectory = map[string][]string{
		"Documents": {".docx"},
		"Pictures":  {".jpg", ".jpeg", ".png", ".webp", ".gif", ".avif", ".svg", ".ai", ".eps", ".raw", ".dng", ".cr2", ".cr3", ".nef", ".nrw", ".arw", ".psd", ".xcf", ".tiff", ".tif", ".heic", ".heif", ".bmp", ".ico"},
		"Musics":    {".flac", ".wav", ".alac", ".aiff", ".ape", ".wv", ".mp3", ".aac", ".ogg", ".opus", ".wma", ".m4a", ".mid", ".midi", ".m4b"},
		"Videos":    {".mp4", ".mkv", ".webm", ".mov", ".avi", ".m4v", ".wmv", ".flv", ".3gp", ".vob", ".ogv", ".mts", ".m2ts", ".mxf"},
	}
)

func CreateDefaultPolicyStructsByPath(rootPath string) ([]domain.AllowedExtensionsPolicy, error) {
	policies := []domain.AllowedExtensionsPolicy{}
	for dirName, fileExtensions := range extensionsByDirectory {
		policy, err := domain.NewAllowedExtensionsPolicy(filepath.Join(rootPath, dirName), fileExtensions)
		if err != nil {
			return nil, fmt.Errorf("Failed to instantiate %s directory: %v", policy.Dir.Name, err)
		}
		policies = append(policies, *policy)
	}
	return policies, nil
}

func FindDirectoryByName(arr []domain.AllowedExtensionsPolicy, name string) (*domain.AllowedExtensionsPolicy, bool) {
	dirIndex := slices.IndexFunc(arr, func(d domain.AllowedExtensionsPolicy) bool {
		return d.Dir.Name == name
	})

	if dirIndex == -1 {
		return nil, false
	}

	return &arr[dirIndex], true
}

func FindDirectoryByFileExtension(arr []domain.AllowedExtensionsPolicy, ext string) (*domain.AllowedExtensionsPolicy, bool) {
	dirIndex := slices.IndexFunc(arr, func(d domain.AllowedExtensionsPolicy) bool {
		return d.IsAllowed(ext) == true
	})

	if dirIndex == -1 {
		return nil, false
	}

	return &arr[dirIndex], true
}

func IsDirectoryPath(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}
