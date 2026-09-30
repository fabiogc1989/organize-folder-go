package entity

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

var ErrInvalidName = errors.New("The directory name cannot be empty.")

type Directory struct {
	Name string
	Path string
	Perm os.FileMode
}

func NewDirectory(name string, path string, perm os.FileMode) (*Directory, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, ErrInvalidName
	}

	cleanPath := filepath.Clean(path)

	if perm == 0 {
		perm = 0755
	}

	return &Directory{
		Name: name,
		Path: cleanPath,
		Perm: perm,
	}, nil
}

func (d Directory) Exists() bool {
	info, err := os.Stat(d.Path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return err == nil && info.IsDir()
}

func (d Directory) Create() error {
	if err := os.Mkdir(d.Path, d.Perm); err != nil {
		return err
	}
	return nil
}

func (d Directory) GetFiles() ([]FileInfo, error) {
	var (
		files   []FileInfo
		lastErr error
	)
	files = []FileInfo{}

	entries, err := os.ReadDir(d.Path)
	if err != nil {
		return files, err
	}

	for _, fileEntry := range entries {
		if fileEntry.IsDir() {
			continue
		}

		fileInfo, err := fileEntry.Info()
		if err != nil {
			lastErr = err
		}

		fullName := filepath.Join(d.Path, fileEntry.Name())

		files = append(files, FileInfo{
			Name:      fileEntry.Name(),
			FullName:  fullName,
			Size:      fileInfo.Size(),
			Extension: filepath.Ext(fullName),
		})
	}

	return files, lastErr
}
