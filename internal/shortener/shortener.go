package shortener

import (
	"path/filepath"
	"sync"

	"github.com/tiagosantini/shortgo/internal/models"
	"github.com/tiagosantini/shortgo/internal/random"
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

	// Generate URL hexadecimal string
	hexString, err := random.GenerateHexString()
	if err != nil {
		return nil, err
	}

	// Create and store shortened URL
	shortenedUrl := models.ShortenedUrl{
		Key: hexString,
		Url: original.Url,
	}

	storedData = append(storedData, shortenedUrl)

	err = storage.JsonEncode(appDataFilePath, storedData)
	if err != nil {
		return nil, err
	}

	return &shortenedUrl, nil;
}