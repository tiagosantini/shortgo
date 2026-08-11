package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/tiagosantini/shortgo/internal/models"
)

func GetAppDataDirPath() (string, error) {
	// Find/Create user Local App Data directory
	cacheDir, err := os.UserCacheDir();
	if err != nil {
		return "", errors.New("Failed to read user cache directory.")
	}

	appDir := filepath.Join(cacheDir, "shortgo")
	if err = os.MkdirAll(appDir, 0o755); err != nil {
		return "nil", errors.New("Failed to create app data folder.")
	}

	return appDir, nil;
}

func JsonEncode(filePath string, data []models.ShortenedUrl) error {
	// Create or open data file handle
	fileHandle, err := os.Create(filePath)
	if err != nil {
		return errors.New("Failed to locate or create app data.")
	}
	defer fileHandle.Close()

	// Encode data into JSON
	encoder := json.NewEncoder(fileHandle)
	encoder.SetIndent("", "    ") 

	err = encoder.Encode(data)
	if err != nil {
		return errors.New("Failed to encode JSON.")
	}

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