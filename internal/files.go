package internal

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func writeToFile(path string, dataArray []string) error {
	filename := filepath.Clean(path)
	_, err := os.Stat(filename)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("failed to stat file (%s): %w", filename, err)
		}
	}
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file (%s): %w", filename, err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("failed to close file (%s): %v\n", filename, err)
		}
	}(f)
	for _, data := range dataArray {
		_, err = f.WriteString(data)
		if err != nil {
			return fmt.Errorf("failed to write to file (%s): %w", filename, err)
		}
	}
	return nil
}

func WriteCertificate(cfg *Config, cerfificate string) error {
	err := writeToFile(cfg.CertPath, []string{cerfificate})
	if err != nil {
		return fmt.Errorf("failed to write certificate to file: %w", err)
	}
	return nil

}

func WritePrivateKey(cfg *Config, privateKey string) error {
	err := writeToFile(cfg.KeyPath, []string{privateKey})
	if err != nil {
		return fmt.Errorf("failed to write private key to file: %w", err)
	}
	return nil
}

func WriteCaChain(cfg *Config, caChain []string) error {
	err := writeToFile(cfg.CaChainPath, caChain)
	if err != nil {
		return fmt.Errorf("failed to write ca-chain to file: %w", err)
	}
	return nil
}
