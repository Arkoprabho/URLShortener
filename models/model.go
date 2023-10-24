package models

import (
	"net/url"

	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type URL struct {
	ShortenedUrl   string `dynamodbav:"shortenedUrl"`
	DestinationUrl string `dynamodbav:"destinationUrl"`
}

func (tinyUrl *URL) GetKey() map[string]types.AttributeValue {
	shortenedUrl, err := attributevalue.Marshal(tinyUrl.ShortenedUrl)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"shortenedUrl": shortenedUrl}
}

func (tinyUrl *URL) GetItem(cfg aws.Config, tableName string) (bool, error) {
	svc := dynamodb.NewFromConfig(cfg)

	resp, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key:       tinyUrl.GetKey(),
		TableName: aws.String(tableName),
	})
	if err != nil {
		log.Fatalf("Coud not get item, %v", err)
	}
	if resp.Item == nil {
		log.Printf("Could not find %v\n", tinyUrl.ShortenedUrl)
		msg := "Could not find '" + tinyUrl.ShortenedUrl + "'"
		return false, errors.New(msg)
	}

	err = attributevalue.UnmarshalMap(resp.Item, &tinyUrl)

	log.Printf("Short URL: %s\n", tinyUrl.ShortenedUrl)
	log.Printf("Destination URL: %s\n", tinyUrl.DestinationUrl)

	return true, err
}

func (tinyUrl *URL) PutItem(cfg aws.Config, tableName string) (bool, error) {
	svc := dynamodb.NewFromConfig(cfg)
	item, err := attributevalue.MarshalMap(tinyUrl)

	if err != nil {
		log.Fatalf("Unable to marshall object %v", err)
		return false, err
	}
	_, err = svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})
	if err != nil {
		log.Fatalf("Unable to add item in DB: %v", err)
		return false, err
	}
	log.Printf("Put item: %v to table: %v", tinyUrl, tableName)
	return true, nil
}
func isValidUrl(sourceUrl string) bool {
	value, err := url.Parse(sourceUrl)
	return err == nil && value.Scheme != "" && value.Host != ""
}
