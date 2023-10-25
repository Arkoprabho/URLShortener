package models

import (
	"testing"
)

func TestShortURL(t *testing.T) {
	tinyUrl := URL{
		DestinationUrl: "https://www.bing.com/search?q=method+documentation+in+golang",
	}
	expectedShortURL := "aHR0cHM6Ly93d3cuYmluZy5jb20vc2VhcmNoP3E9bWV0aG9kK2RvY3VtZW50YXRpb24raW4rZ29sYW5n"

	shortUrl := make(chan string)
	errorChannel := make(chan error)
	go tinyUrl.GenerateShortURL(shortUrl, errorChannel)
	generatedShortURL := <-shortUrl
	err := <-errorChannel

	if err != nil {
		t.Errorf("Error generating short URL with valid URL")
	}
	if generatedShortURL != expectedShortURL {
		t.Errorf("Generated URL: %v does not match expected URL: %v", generatedShortURL, expectedShortURL)
	}
}
