package files

import (
	"fmt"
	"os"
	"path/filepath"
)

// CheckDirWritable checks if a directory is writable
func CheckDirWritable(dir string) (bool, error) {
	testFile := filepath.Join(dir, ".writetest")
	testFile = filepath.Clean(testFile)
	file, err := os.Create(testFile)
	if err != nil {
		return false, nil
	}
	err = file.Close()
	if err != nil {
		return false, fmt.Errorf("failed to close file: %w", err)
	}
	err = os.Remove(testFile)
	if err != nil {
		return false, fmt.Errorf("failed to remove file: %w", err)
	}
	return true, nil
}

// CheckFileWritable checks if a file is writable
func CheckFileWritable(file string) (bool, error) {
	file = filepath.Clean(file)
	f, err := os.OpenFile(file, os.O_WRONLY, 0600)
	if err != nil {
		return false, nil
	}
	err = f.Close()
	if err != nil {
		return false, fmt.Errorf("failed to close file: %w", err)
	}
	return true, nil
}

// CheckFileExists checks if a file exists
func CheckFileExists(file string) (bool, error) {
	file = filepath.Clean(file)
	_, err := os.Stat(file)
	if err != nil {
		if err != os.ErrNotExist {
			return false, nil
		}
		return false, fmt.Errorf("failed to check file: %w", err)
	}
	return !os.IsNotExist(err), nil
}

// DirExists checks if a directory exists
func DirExists(dir string) bool {
	dir = filepath.Clean(dir)
	info, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// MkDir creates a directory if it does not exist
func MkDir(dir string) error {
	dir = filepath.Clean(dir)
	return os.MkdirAll(dir, 0750)
}

// GetDirFromFile returns the directory part of a file path
func GetDirFromFile(file string) string {
	file = filepath.Clean(file)
	return filepath.Dir(file)
}
