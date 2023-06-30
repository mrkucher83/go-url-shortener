package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestInsertionAndRetieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortUrl := "Jsz4k57oAX"

	// Persist data mapping
	SaveUrlMapping(shortUrl, initialLink)

	// Retrieve initial URL
	retrieveUrl := RetrieveInitialUrl(shortUrl)

	assert.Equal(t, initialLink, retrieveUrl)
}
