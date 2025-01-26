package files_test

import (
	"os"
	"path/filepath"
	"testing"
	"vault-get-cert/internal/files"
)

func TestCheckDirWritable(t *testing.T) {
	dir := os.TempDir()
	result, err := files.CheckDirWritable(dir)
	if err != nil {
		t.Fatalf("failed to check directory: %v", err)
	}
	if !result {
		t.Fatalf("expected directory %s to be writable", dir)
	}
}

func TestCheckFileWritable(t *testing.T) {
	file := filepath.Join(os.TempDir(), "testfile.txt")
	defer os.Remove(file)

	if err := os.WriteFile(file, []byte("test"), 0666); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	result, err := files.CheckFileWritable(file)
	if err != nil {
		t.Fatalf("failed to check file: %v", err)
	}
	if !result {
		t.Fatalf("expected file %s to be writable", file)
	}
}

func TestCheckFileExists(t *testing.T) {
	file := filepath.Join(os.TempDir(), "testfile.txt")
	defer os.Remove(file)

	if err := os.WriteFile(file, []byte("test"), 0666); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	result, err := files.CheckFileExists(file)
	if err != nil {
		t.Fatalf("failed to check file: %v", err)
	}
	if !result {
		t.Fatalf("expected file %s to exist", file)
	}

	result, err = files.CheckFileExists(filepath.Join(os.TempDir(), "nonexistent.txt"))
	if err != nil {
		t.Fatalf("failed to check file: %v", err)
	}
	if result {
		t.Fatalf("expected file nonexistent.txt to not exist")
	}
}

func TestDirExists(t *testing.T) {
	dir := os.TempDir()
	if !files.DirExists(dir) {
		t.Fatalf("expected directory %s to exist", dir)
	}

	if files.DirExists(filepath.Join(os.TempDir(), "nonexistent")) {
		t.Fatalf("expected directory nonexistent to not exist")
	}
}

func TestMkDir(t *testing.T) {
	dir := filepath.Join(os.TempDir(), "testdir")
	defer os.Remove(dir)

	if err := files.MkDir(dir); err != nil {
		t.Fatalf("failed to create directory %s: %v", dir, err)
	}

	if !files.DirExists(dir) {
		t.Fatalf("expected directory %s to exist", dir)
	}
}

func TestGetDirFromFile(t *testing.T) {
	file := "/path/to/file.txt"
	expectedDir := "/path/to"
	if dir := files.GetDirFromFile(file); dir != expectedDir {
		t.Fatalf("expected directory %s, got %s", expectedDir, dir)
	}
}
