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
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "invalid request body",
		}, nil
	}

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "failed to marshal book",
		}, nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"))
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
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
		log.Fatalf("failed to put item, %v", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Book added",
	}, nil

}

func main() {
	lambda.Start(insertBook)
}
