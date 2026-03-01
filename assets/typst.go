package assets

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

//go:embed bin/*
var typstFS embed.FS

var (
	binaryPath string
	once       sync.Once
	initErr    error
)

func Init() (string, error) {
	once.Do(func() {
		var embedPath string
		switch runtime.GOOS {
		case "linux":
			embedPath = "bin/typst-linux"
		case "darwin":
			embedPath = "bin/typst-apple"
		case "windows":
			embedPath = "bin/typst-window.exe"
		default:
			initErr = fmt.Errorf("unsupported OS: %s\n", runtime.GOOS)
			return
		}

		data, err := typstFS.ReadFile(embedPath)
		if err != nil {
			initErr = fmt.Errorf("failed to read embedded binary: %w\n", err)
		}
		tmpDirPath := os.TempDir()
		tmpDir, err := os.MkdirTemp(tmpDirPath, "invoice-gen-*")
		if err != nil {
			initErr = err
			return
		}

		binaryName := filepath.Base(embedPath)
		binaryPath := filepath.Join(tmpDir, binaryName)

		if err := os.WriteFile(binaryPath, data, 0o755); err != nil {
			initErr = err
			return
		}
	})

	return binaryPath, initErr
}

func Cleanup() {
	if binaryPath != "" {
		os.RemoveAll(filepath.Dir(binaryPath))
	}
}
