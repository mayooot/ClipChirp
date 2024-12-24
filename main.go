package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Validate command-line arguments
	if len(os.Args) < 2 {
		log.Fatal("File path argument is required")
	}

	// Get the input file path
	inputPath := os.Args[1]

	// Resolve the absolute path using realpath
	absPath, err := resolveAbsolutePathInGo(inputPath)
	if err != nil {
		log.Fatalf("Failed to resolve absolute path: %v", err)
	}

	// Copy the file to the clipboard using Finder as context
	if err := copyToClipboardUsingFinder(absPath); err != nil {
		log.Fatalf("Failed to copy file to clipboard: %v", err)
	}

	fmt.Printf("🦤 Successfully copied %s to clipboard\n", absPath)
}

// resolveAbsolutePathUsingRealpath resolves the absolute path of a file using realpath
func resolveAbsolutePathUsingRealpath(path string) (string, error) {
	cmd := exec.Command("realpath", path)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error executing realpath: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func resolveAbsolutePathInGo(path string) (string, error) {
	return filepath.Abs(path)
}

// copyToClipboardUsingFinder copies the given file to the clipboard using Finder as context
func copyToClipboardUsingFinder(path string) error {
	script := fmt.Sprintf(`tell application "Finder" to set the clipboard to (POSIX file "%s")`, path)
	cmd := exec.Command("osascript", "-e", script)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error executing osascript: %w", err)
	}
	return nil
}
