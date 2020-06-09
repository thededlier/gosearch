package impl 

import (
	"bytes"
	"context"
	"log"
	"fmt"
	"encoding/json"
	"sync"
	"strings"

	"gosearch/internal/gRPC/domain"
	"gosearch/internal/gRPC/service"
	"gosearch/internal/searchclient"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type SearchService struct {
	esClient *elasticsearch.Client
}

type EsInfo struct {
	ClusterName string		`json:"cluster_name"`
	ClusterId	string		`json:"cluster_uuid"`
	Version		EsVersion	`json:"version"`
}

type EsVersion struct {
	Number		string		`json:"number"`	
}

func NewSearchService() *SearchService {
	esClient := searchclient.CreateNewElasticsearchClient()
	return &SearchService{
		esClient: esClient,
	}
}

// Get information of the elasticsearch cluster
func (searchService *SearchService) GetInfo(ctx context.Context, emptyRequest *domain.EmptyRequest) (*service.GetInfoResponse, error) {
	log.Println("Request::GetInfo")

	res, err := searchService.esClient.Info()

	log.Println(res)

	if err != nil {
		log.Fatalf("Error getting status: %v", err)

		requestError := service.Error{
			Code: "500",
			Message: "An exception occurred",
		}
		
		return &service.GetInfoResponse{
			Info: nil,
			Error: &requestError,
		}, nil
	} else {
		defer res.Body.Close()
		esInfo := &EsInfo{}
		
		// Decode info 
		if err := json.NewDecoder(res.Body).Decode(esInfo); err != nil {
			requestError := service.Error{
				Code: "500",
				Message: fmt.Sprintf("Failed to deserialize json: %v", err),
			}
			
			return &service.GetInfoResponse{
				Info: nil,
				Error: &requestError,
			}, nil
		}

		esStatus := domain.ElasticsearchInfo {
			Status: "OK",
			ClusterName: esInfo.ClusterName,
			ClusterId: esInfo.ClusterId,
			ElasticsearchVersion: esInfo.Version.Number,
		}
		
		return &service.GetInfoResponse{
			Info: &esStatus,
			Error: nil,
		}, nil
	}
}

// Index a list of documents
func (searchService *SearchService) IndexDocuments(ctx context.Context, indexRequest *service.IndexDocumentsRequest) (*service.IndexDocumentsResponse, error) {
	log.Println("Request::IndexDocuments")

	var (
		failedList []string
		wg sync.WaitGroup
	)

	// Loop through the documents and index each of them concurrently
	for i, doc := range indexRequest.DocumentsList {
		wg.Add(1)

		go func(i int, doc *domain.Document) {
			defer wg.Done()

			req := esapi.IndexRequest{
				Index:      doc.Index,
				DocumentID: doc.DocumentID,
				Body:       strings.NewReader(doc.Body),
			}
		
			// Perform the request with the client.
			res, err := req.Do(context.Background(), searchService.esClient)

			if err != nil {
				log.Fatalf("Failed to fetch response from Elasticsearch: %v", err)
				failedList = append(failedList, doc.DocumentID)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Println("[%s] Failed to index Document ID=%s", res.Status(), doc.DocumentID)
				failedList = append(failedList, doc.DocumentID)
			} else {
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				  log.Printf("Error parsing the response body: %s", err)
				} else {
				  log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		} (i, doc)
	}

	wg.Wait()

	return &service.IndexDocumentsResponse{
		DocumentsList: indexRequest.DocumentsList,
		Failed: failedList,
		Error: nil,
	}, nil
}

// Process a custom search query
func (searchService *SearchService) CustomSearch (ctx context.Context, searchRequest *service.CustomSearchRequest) (*service.SearchResponse, error) {
	log.Println("Request::CustomSearch")

	var result map[string]interface{}
	
	// Es search request
	res, err := searchService.esClient.Search(
		searchService.esClient.Search.WithContext(context.Background()),
		searchService.esClient.Search.WithIndex(searchRequest.Indexes),
		searchService.esClient.Search.WithBody(bytes.NewBuffer(searchRequest.Query)),
		searchService.esClient.Search.WithTrackTotalHits(true),
		searchService.esClient.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Failed to fetch response from elasticsearch : %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf(
				"Elasticsearch::[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(result["took"].(float64)),
	)

	var hitDocList []*domain.Document
	// Print the ID and document source for each hit.
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		docBody, _ := json.Marshal(hit.(map[string]interface{})["_source"])

		doc := domain.Document{
			Index: hit.(map[string]interface{})["_index"].(string),
			DocumentID: hit.(map[string]interface{})["_id"].(string),
			Body: string(docBody),
		}
		hitDocList = append(hitDocList, &doc)

		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	return &service.SearchResponse{
		TotalHits: int32(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		Hits: hitDocList,
	}, nil
}