package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description"`
	Year        int    `json:"year"`
}

// function that returns all books written by author specified in path parameter
func getBook(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get the author from the path parameter and convert %20 to space
	author := request.PathParameters["author"]
	toReplace := "%20"
	author = strings.ReplaceAll(author, toReplace, " ")
	log.Printf("author: %s", author)

	// Create the DynamoDB client
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"))
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)

	// Create the expression to fill the input struct with.
	// Get all books written by author
	keyExp := expression.Key("author").Equal(expression.Value(author))
	expr, err := expression.NewBuilder().WithKeyCondition(keyExp).Build()
	if err != nil {
		log.Fatalf("failed to create expression, %v", err)
	}

	resp, err := client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String(os.Getenv("TABLE_NAME")),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	if err != nil {
		log.Fatalf("failed to query table, %v", err)
	}

	// if response is empty, return 404
	if len(resp.Items) == 0 {
		return events.APIGatewayProxyResponse{
			Body:       "No books found",
			StatusCode: http.StatusNotFound,
		}, nil
	}

	// Unmarshal the items in the response
	var books []Book
	err = attributevalue.UnmarshalListOfMaps(resp.Items, &books)
	if err != nil {
		log.Fatalf("failed to unmarshal Dynamodb Scan Items, %v", err)
	}

	// Marshal the books into a JSON string
	booksJSON, err := json.Marshal(books)
	if err != nil {
		log.Fatalf("failed to marshal books, %v", err)
	}

	// Return the books as a JSON string
	return events.APIGatewayProxyResponse{
		Body:       string(booksJSON),
		StatusCode: http.StatusOK,
	}, nil

}

func main() {
	lambda.Start(getBook)
}
