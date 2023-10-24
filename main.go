package main

import (
	"context"
	"log"

	"github.com/Arkoprabho/URLShortener/models"
	"github.com/aws/aws-sdk-go-v2/config"
)

func init() {
	log.SetPrefix("[URL Shortener] ")
}

func main() {
	sourceURL := "Superman"
	log.Println("Starting URL Shortener")
	prefix := ""
	valid := models.IsValidUrl(sourceURL)
	if !valid {
		prefix = "in"
	}
	log.Printf("URL is %svalid", prefix)

	log.Println("Listing tables")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}
	tinyUrl := models.URL{
		ShortenedUrl: "http://localhost:3000/aHR0cHM6Ly9nb2J5ZXhhbXBsZS5jb20vY29tbWFuZC1saW5lLWFyZ3VtZW50cwo=",
	}
	tinyUrl.GetItem(cfg, "URLShortener")
}
