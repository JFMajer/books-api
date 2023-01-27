package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
	Year        int    `json:"year"`
}

func getBook(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	title := request.QueryStringParameters["title"]
	author := request.QueryStringParameters["author"]

	// Create the DynamoDB client
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("failed to load config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	req, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"title": &types.AttributeValueMemberS{
				Value: title,
			},
			"author": &types.AttributeValueMemberS{
				Value: author,
			},
		},
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	})
	if err != nil {
		log.Fatalf("could not get item, %v", err)
	}

	// Create the response
	book := Book{
		Title:       req.Item["title"].(*types.AttributeValueMemberS).Value,
		Author:      req.Item["author"].(*types.AttributeValueMemberS).Value,
		ISBN:        req.Item["isbn"].(*types.AttributeValueMemberS).Value,
		Description: req.Item["description"].(*types.AttributeValueMemberS).Value,
		Year:        req.Item["year"].(*types.AttributeValueMemberN).Value,
	}

	// marshal the response into json
	bookToReturn, err := json.Marshal(book)
	if err != nil {
		log.Fatalf("could not marshal response, %v", err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(bookToReturn),
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil

}

func main() {
	lambda.Start(getBook)
}
