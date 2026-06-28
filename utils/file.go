package utils

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

// CreateTempFile saves an uploaded multipart file to a temporary location
// and returns the closed file handle. The caller should os.Remove the
// path when done.
func CreateTempFile(file multipart.File, filename string) (*os.File, error) {
	tmpFile, err := os.CreateTemp("", "upload-*"+filepath.Ext(filename))
	if err != nil {
		return nil, fmt.Errorf("create temp file: %w", err)
	}

	if _, err := io.Copy(tmpFile, file); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return nil, fmt.Errorf("write temp file: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		os.Remove(tmpFile.Name())
		return nil, fmt.Errorf("close temp file: %w", err)
	}

	return tmpFile, nil
}

// RemoveFile deletes a file at path, logging but not returning errors.
func RemoveFile(path string) {
	if err := os.Remove(path); err != nil {
		log.Printf("failed to remove file %s: %v", path, err)
	}
}
