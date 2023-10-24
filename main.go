package main

import (
	"context"
	"log"
	"os"

	"github.com/Arkoprabho/URLShortener/models"
	"github.com/aws/aws-sdk-go-v2/config"
)

func init() {
	log.SetPrefix("[URL Shortener] ")
}

func main() {
	tableName := "URLShortener"
	log.Println("Starting URL Shortener")
	if len(os.Args) < 3 {
		log.Fatal("Not enough arguments")
		panic("Not enough arguments")
	}

	sourceUrl := os.Args[1]
	operation := os.Args[2]

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	if operation == "store" {
		log.Printf("Storing %v", sourceUrl)
		tinyUrl := models.URL{
			DestinationUrl: sourceUrl,
		}
		shortUrl, err := tinyUrl.GenerateShortURL()
		log.Printf("Short URL: %v", shortUrl)
		if err != nil {
			panic(err)
		}
		tinyUrl.ShortenedUrl = shortUrl
		_, err = tinyUrl.PutItem(cfg, tableName)
		if err != nil {
			panic(err)
		}
	}

	if operation == "fetch" {
		log.Printf("Fetching the destination URL for %v", sourceUrl)
		tinyUrl := models.URL{
			ShortenedUrl: sourceUrl,
		}
		tinyUrl.GetItem(cfg, tableName)
		log.Printf("%v", tinyUrl.DestinationUrl)
	}

}
