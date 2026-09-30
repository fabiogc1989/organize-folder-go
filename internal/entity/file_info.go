package entity

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Name      string
	FullName  string
	Size      int64
	Extension string
}

func NewFileInfo(path string) (*FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return &FileInfo{
		Name:      info.Name(),
		FullName:  path,
		Size:      info.Size(),
		Extension: filepath.Ext(path),
	}, nil
}

func (f *FileInfo) Move(folderPath string) error {
	newPath := filepath.Join(folderPath, f.Name)
	if err := os.Rename(f.FullName, newPath); err != nil {
		return fmt.Errorf("\033[31mError moving file %s: %v\033[0m", f.Name, err)
	}

	f.FullName = newPath

	return nil
}
