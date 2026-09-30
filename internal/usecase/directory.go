package usecase

import (
	"fmt"
	"path/filepath"
	"slices"

	"github.com/fabiogc1989/organize-folder/internal/entity"
)

var (
	DefaultDirectoryNames = []string{"Documents", "Pictures", "Musics", "Videos"}
	Directories           []entity.Directory
)

func CreateDirectoryStructsByPath(rootPath string) ([]entity.Directory, error) {
	directories := make([]entity.Directory, 0, len(DefaultDirectoryNames))
	for _, dirName := range DefaultDirectoryNames {
		dir, err := entity.NewDirectory(dirName, filepath.Join(rootPath, dirName), 0755)
		if err != nil {
			return []entity.Directory{}, fmt.Errorf("\033[31mFailed to instantiate %s directory: %v\033[0m", dir.Name, err)
		}
		directories = append(directories, *dir)
	}
	return directories, nil
}

func CreateDirectories(dirs []entity.Directory) error {
	for _, dir := range dirs {
		if dirExists := dir.Exists(); dirExists == false {
			err := dir.Create()
			if err != nil {
				return fmt.Errorf("\033[31mFailed to create %s directory: %v\033[0m", dir.Name, err)
			}
		}
	}
	return nil
}

func FindDirectoryByName(name string) (*entity.Directory, bool) {
	dirIndex := slices.IndexFunc(Directories, func(d entity.Directory) bool {
		return d.Name == name
	})

	if dirIndex == -1 {
		return nil, false
	}

	return &Directories[dirIndex], true
}
