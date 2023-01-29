package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

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

func insertBook(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var book Book
	// Unmarshal the request body into a Book struct
	err := json.Unmarshal([]byte(request.Body), &book)

	if err != nil {
		log.Printf("Failed to unmarshal book, here is why: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Failed to marshal book",
		}, nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"))
	if err != nil {
		log.Fatalf("Failed to load configuration, here is why: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to load configuration",
		}, nil
	}

	client := dynamodb.NewFromConfig(cfg)
	resp, err := client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Item: map[string]types.AttributeValue{
			"isbn":        &types.AttributeValueMemberS{Value: book.ISBN},
			"title":       &types.AttributeValueMemberS{Value: book.Title},
			"author":      &types.AttributeValueMemberS{Value: book.Author},
			"description": &types.AttributeValueMemberS{Value: book.Description},
			"year":        &types.AttributeValueMemberN{Value: strconv.Itoa(book.Year)},
		},
	})
	if err != nil {
		log.Fatalf("Failed to put item, here is why: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to put item",
		}, nil
	}

	// Marshal the response into a JSON string
	returnBody, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Failed to marshal response, here is why: %v", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(returnBody),
	}, nil

}

func main() {
	lambda.Start(insertBook)
}
