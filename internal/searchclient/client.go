package searchclient

import (
	"log"
	
	"github.com/elastic/go-elasticsearch/v8"
)

func CreateNewElasticsearchClient() (esClient *elasticsearch.Client) {
	esClient, err := elasticsearch.NewDefaultClient()
	
	if err != nil {
		log.Fatalf("Failed to create Elasticsearch Client")
		panic(err)
	}
	
	return esClient
}