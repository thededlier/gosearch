package main

import (
	"context"
	"fmt"
	"encoding/json"

	"gosearch/internal/gRPC/domain"
	"gosearch/internal/gRPC/service"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:8888"

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := service.NewSearchServiceClient(conn)


	// Get elasticsearch information

	status := domain.EmptyRequest{}

	if responseMessage, err := client.GetInfo(context.Background(), &status); err != nil {
		panic(fmt.Sprintf("Was not able to check status: %v", err))
	} else {
		fmt.Println("Status received")
		fmt.Println(responseMessage)
	}

	// Index a document

	document := domain.Document {
		Index: "test",
		DocumentID: "1",
		Body: `{ "title": "Test" }`,
	}

	docList := []*domain.Document { &document }

	indexRequest := service.IndexDocumentsRequest {
		DocumentsList: docList,
	}

	if responseMessage, err := client.IndexDocuments(context.Background(), &indexRequest); err != nil {
		panic(fmt.Sprintf("Failed to index: %v", err))
	} else {
		fmt.Println("Status received")
		fmt.Println(responseMessage)
	}

	// Search for the document
	
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "test",
			},
		},
	}

	queryString, _ := json.Marshal(query) 

	searchRequest := service.CustomSearchRequest{
		Indexes: "test",
		Query: queryString,
	}

	if responseMessage, err := client.CustomSearch(context.Background(), &searchRequest); err != nil {
		panic(fmt.Sprintf("Failed to search: %v", err))
	} else {
		fmt.Println("Search Completed")
		fmt.Println(responseMessage)
	}
}