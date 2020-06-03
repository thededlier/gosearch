package impl 

import (
	"context"
	"log"
	"fmt"
	"encoding/json"

	"gosearch/internal/gRPC/domain"
	"gosearch/internal/gRPC/service"
	"gosearch/internal/searchclient"

	"github.com/elastic/go-elasticsearch/v8"
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