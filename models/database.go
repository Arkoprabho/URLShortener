package models

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ListTables(cfg aws.Config) {
	svc := dynamodb.NewFromConfig(cfg)

	resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})

	if err != nil {
		log.Fatalf("Failed to list tables: %v", err)

		log.Println("Tables: ")
		for _, tableName := range resp.TableNames {
			log.Println(tableName)
		}
	}
}

func GetItem(cfg aws.Config, tinyUrl URL, tableName string) (bool, error) {
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

	receivedItem := URL{}
	err = attributevalue.UnmarshalMap(resp.Item, &receivedItem)

	log.Printf("Short URL: %s\n", receivedItem.ShortenedUrl)
	log.Printf("Destination URL: %s\n", receivedItem.DestinationUrl)

	return true, err
}
