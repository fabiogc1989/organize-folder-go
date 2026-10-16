package domain

import (
	"errors"
	"slices"
)

type AllowedExtensionsPolicy struct {
	Dir        Directory
	Extensions []string
}

func NewAllowedExtensionsPolicy(directoryPath string, exts []string) (*AllowedExtensionsPolicy, error) {
	dir, err := NewDirectory("sadad", directoryPath, 0755)
	if err != nil {
		return nil, err
	}

	if len(exts) == 0 {
		return nil, errors.New("The file extensions are empty.")
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
