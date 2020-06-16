package searchclient

import (
	"os"
	"log"
	"github.com/elastic/go-elasticsearch/v8"
)

func CreateNewElasticsearchClient() (esClient *elasticsearch.Client) {
	cfg := elasticsearch.Config{
		Addresses: []string{
		  	os.Getenv("ELASTICSEARCH_URL"),
		},
	}

	esClient, err := elasticsearch.NewClient(cfg)
	
	if err != nil {
		log.Fatalf("Failed to create Elasticsearch Client")
		panic(err)
	}
	
	return esClient
}