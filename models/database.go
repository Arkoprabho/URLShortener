package models

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
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
