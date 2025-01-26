package internal

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	testFile := filepath.Join(os.TempDir(), "testfile.txt")
	defer os.Remove(testFile)

	data := []string{"line1", "line2"}
	err := writeToFile(testFile, data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("expected no error reading file, got %v", err)
	}

	expectedContent := "line1line2"
	if string(content) != expectedContent {
		t.Fatalf("expected content %s, got %s", expectedContent, string(content))
	}
}

func TestWriteCertificate(t *testing.T) {
	cfg := &Config{CertPath: filepath.Join(os.TempDir(), "cert.pem")}
	defer os.Remove(cfg.CertPath)

	certificate := "test-certificate"
	err := WriteCertificate(cfg, certificate)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	content, err := os.ReadFile(cfg.CertPath)
	if err != nil {
		t.Fatalf("expected no error reading file, got %v", err)
	}

	if string(content) != certificate {
		t.Fatalf("expected content %s, got %s", certificate, string(content))
	}
}

func TestWritePrivateKey(t *testing.T) {
	cfg := &Config{KeyPath: filepath.Join(os.TempDir(), "key.pem")}
	defer os.Remove(cfg.KeyPath)

	privateKey := "test-private-key"
	err := WritePrivateKey(cfg, privateKey)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	content, err := os.ReadFile(cfg.KeyPath)
	if err != nil {
		t.Fatalf("expected no error reading file, got %v", err)
	}

	if string(content) != privateKey {
		t.Fatalf("expected content %s, got %s", privateKey, string(content))
	}
}

func TestWriteCaChain(t *testing.T) {
	cfg := &Config{CaChainPath: filepath.Join(os.TempDir(), "ca-chain.pem")}
	defer os.Remove(cfg.CaChainPath)

	caChain := []string{"ca1", "ca2"}
	err := WriteCaChain(cfg, caChain)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	content, err := os.ReadFile(cfg.CaChainPath)
	if err != nil {
		t.Fatalf("expected no error reading file, got %v", err)
	}

	expectedContent := "ca1ca2"
	if string(content) != expectedContent {
		t.Fatalf("expected content %s, got %s", expectedContent, string(content))
	}
}
