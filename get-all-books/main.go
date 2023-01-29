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
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
	Year        int    `json:"year"`
}

func getAllBooks(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	numOfItems, err := strconv.Atoi(request.Headers["numOfItems"])
	if err != nil {
		log.Printf("Failed to grab numbers of items to return from request headers, %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Failed to grab numbers of items to return from request headers",
		}, err
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"))
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to load configuration",
		}, err
	}

	client := dynamodb.NewFromConfig(cfg)
	req, err := client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Limit:     aws.Int32(int32(numOfItems)),
	})
	if err != nil {
		log.Fatalf("failed to scan table, %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to scan table",
		}, err
	}

	books := make([]Book, 0)
	err = attributevalue.UnmarshalListOfMaps(req.Items, &books)
	if err != nil {
		log.Fatalf("failed to unmarshal books, %v", err)
	}

	body, err := json.Marshal(books)
	if err != nil {
		log.Fatalf("failed to marshal books, %v", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}, nil

}

func main() {
	lambda.Start(getAllBooks)
}
