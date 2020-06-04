package main

import (
	"context"
	"fmt"

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

	status := domain.EmptyRequest{}

	if responseMessage, err := client.GetInfo(context.Background(), &status); err != nil {
		panic(fmt.Sprintf("Was not able to check status: %v", err))
	} else {
		fmt.Println("Status received")
		fmt.Println(responseMessage)
	}

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
}