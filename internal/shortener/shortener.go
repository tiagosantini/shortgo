package shortener

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"path/filepath"
	"sync"

	"github.com/tiagosantini/shortgo/internal/models"
	"github.com/tiagosantini/shortgo/internal/storage"
)

// File lock mutex
var mu sync.Mutex

func GenerateRandomHexString() (string, error) {
	// Create random 8 bytes token
	token := make([]byte, 8)

	_, err := rand.Read(token)
	if err != nil {
		return "", errors.New("Failed to generate random bytes.")
	}

	// Encode token into 16 (8 * 2) character hex string
	return hex.EncodeToString(token), nil
}

func ShortenUrl(original *models.OriginalUrl) (*models.ShortenedUrl, error) {
	
	// Get app data paths
	appDataDirPath, err := storage.GetAppDataDirPath()
	if err != nil {
		log.Fatalf(err.Error(), err)

		return nil, err
	}

	// Try to fetch stored app data
	appDataFilePath := filepath.Join(appDataDirPath, "data.json")

	mu.Lock()
	defer mu.Unlock()

	storedData, err := storage.JsonDecode(appDataFilePath)
	if err != nil {
		log.Fatalf(err.Error(), err)

		return nil, err
	}

	// Generate URL hexadecimal string
	hexString, err := GenerateRandomHexString()
	if err != nil {
		log.Fatalf(err.Error(), err)

		return nil, err
	}

	// Create shortened URL key using the first 8 hex characters
	var shortenedUrl = models.ShortenedUrl{
		Key: hexString,
		Url: original.Url,
	}

	storedData = append(storedData, shortenedUrl)
	err = storage.JsonEncode(appDataFilePath, storedData)
	if err != nil {
		log.Fatalf(err.Error(), err)

		return nil, err
	}

	return &shortenedUrl, nil;
}