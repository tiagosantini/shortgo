package storage

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/tiagosantini/shortgo/internal/models"
)

func JsonEncode(destination string, data []models.ShortenedUrl) error {
	// Get destination directory
	directory := filepath.Dir(destination)

	// Generate random string for tempFile name
	tempfile, err := os.CreateTemp(directory, "*.tmp")
	if err != nil {
		return err
	}

	tempFilename := tempfile.Name()

	log.Printf("[%s] tempfile created", tempFilename)

	// Encode data into JSON
	encoder := json.NewEncoder(tempfile)

	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	// Sync changes and close tempfile
	if err = tempfile.Sync(); err != nil {
    	return err
	}
	if err = tempfile.Close(); err != nil {
		return err
	}

	if err = os.Rename(tempFilename, destination); err != nil {
		return err
	}

	os.Remove(tempFilename)

	log.Printf("[%s] tempfile synced successfully and was removed", tempFilename)

	return nil;
}

func JsonDecode(filePath string) ([]models.ShortenedUrl, error) {
	// Try to open data file handle
	fileHandle, err := os.Open(filePath)
    if errors.Is(err, os.ErrNotExist) {
        return []models.ShortenedUrl{}, nil
    }
    if err != nil {
        return nil, err
    }
	defer fileHandle.Close()

	var storedData []models.ShortenedUrl
	if err = json.NewDecoder(fileHandle).Decode(&storedData); err != nil {
		return nil, err
	}

	return storedData, nil
}

func GetAppDataDirPath() (string, error) {
	// Find/Create user Local App Data directory
	cacheDir, err := os.UserCacheDir();
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(cacheDir, "shortgo")
	if err = os.MkdirAll(appDir, 0o755); err != nil {
		return "nil", err
	}

	return appDir, nil;
}