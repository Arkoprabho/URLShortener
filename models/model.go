package models

import (
	"net/url"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type URL struct {
	ShortenedUrl   string
	DestinationUrl string
}

func (tinyUrl *URL) GetKey() map[string]types.AttributeValue {
	shortenedUrl, err := attributevalue.Marshal(tinyUrl.ShortenedUrl)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"shortenedUrl": shortenedUrl}
}

func IsValidUrl(sourceUrl string) bool {
	value, err := url.Parse(sourceUrl)
	return err == nil && value.Scheme != "" && value.Host != ""
}
