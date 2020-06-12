package searchclient

import (
	"log"
	
	"github.com/elastic/go-elasticsearch/v8"
)

func CreateNewElasticsearchClient() (esClient *elasticsearch.Client) {
	cfg := elasticsearch.Config{
		Addresses: []string{
		  	"http://127.0.0.1:9200",
		},
	}

	esClient, err := elasticsearch.NewClient(cfg)
	
	if err != nil {
		log.Fatalf("Failed to create Elasticsearch Client")
		panic(err)
	}
	
	return esClient
}