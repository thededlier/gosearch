package impl 

import (
	"gosearch/internal/gRPC/domain"
	"gosearch/internal/gRPC/service"
	"context"
	"log"
)

type SearchService struct {}

func NewSearchService() *SearchService {
	return &SearchService{}
}

func (searchService *SearchService) GetStatus(ctx context.Context, esStatus *domain.ElasticsearchStatus) (*service.GetStatusResponse, error) {
	log.Println("Request::GetStatus")

	return &service.GetStatusResponse{
		Status: esStatus,
		Error: nil,
	}, nil
}