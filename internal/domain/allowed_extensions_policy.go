package domain

import (
	"errors"
	"path/filepath"
	"slices"
)

var ErrEmptyExtensions = errors.New("The file extensions are empty.")

type AllowedExtensionsPolicy struct {
	Dir        Directory
	Extensions []string
}

func NewAllowedExtensionsPolicy(directoryPath string, exts []string) (*AllowedExtensionsPolicy, error) {
	dirName := filepath.Base(directoryPath)

	dir, err := NewDirectory(dirName, directoryPath, 0755)
	if err != nil {
		return nil, err
	}

	if len(exts) == 0 {
		return nil, ErrEmptyExtensions
	}

	return &AllowedExtensionsPolicy{
		Dir:        *dir,
		Extensions: exts,
	}, nil
}

func (p AllowedExtensionsPolicy) IsAllowed(ext string) bool {
	return slices.Contains(p.Extensions, ext)
}

func (p *AllowedExtensionsPolicy) AddExtension(ext string) {
	p.Extensions = append(p.Extensions, ext)
}
