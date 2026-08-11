package shortener

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"path/filepath"
	"sync"

	"github.com/tiagosantini/shortgo/internal/models"
	"github.com/tiagosantini/shortgo/internal/storage"
)

// File lock mutex
var mu sync.Mutex

func ShortenUrl(original *models.OriginalUrl) (*models.ShortenedUrl, error) {
	
	// Get app data paths
	appDataDirPath, err := storage.GetAppDataDirPath()
	if err != nil {
		return nil, err
	}

	// Try to fetch stored app data
	appDataFilePath := filepath.Join(appDataDirPath, "data.json")

	mu.Lock()
	defer mu.Unlock()

	storedData, err := storage.JsonDecode(appDataFilePath)
	if err != nil {
		return nil, err
	}

	// Generate short Base64 URL key
	urlKey, err := GenerateBase64String()
	if err != nil {
		return nil, err
	}

	// Create and store shortened URL
	shortenedUrl := models.ShortenedUrl{
		Key: urlKey,
		Url: original.Url,
	}

	storedData = append(storedData, shortenedUrl)

	err = storage.JsonEncode(appDataFilePath, storedData)
	if err != nil {
		return nil, err
	}

	return &shortenedUrl, nil;
}

func GenerateBase64String() (string, error) {
	// Create random 8 bytes token
	token := make([]byte, 6)

	_, err := rand.Read(token)
	if err != nil {
		return "", errors.New("Failed to generate random bytes.")
	}

	// Encode token into base64 string
	return base64.RawURLEncoding.EncodeToString(token), nil
}