package documents

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func ReadingFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return os.ReadFile(path)
}

func CheckFolderStucture(path string) ([]os.DirEntry, error) {
	info, err := os.ReadDir(path)
	if err != nil {
		if os.IsNotExist(err){
			fmt.Println("folder does not exist")
			return nil, err
		}
	}

	return info, nil
}



func FullWalkDown(path string) error {
	return filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			fmt.Println(path + " -> directory")
		} else {
			fmt.Println("L-> " + path + " -> file")
		}
		return nil
	})
}