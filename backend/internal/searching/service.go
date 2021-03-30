package searching

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-kit/kit/log"
	"github.com/paintmeyellow/go-elastic-course/internal/elastic"
)

type Service interface {
	Search(ctx context.Context, req Request) (*Response, error)
}

type service struct {
	esClient *elasticsearch.Client
	logger   log.Logger
}

func NewService(es *elasticsearch.Client, l log.Logger) Service {
	return &service{esClient: es, logger: l}
}

func (s service) Search(ctx context.Context, req Request) (*Response, error) {
	esReq := &elastic.SearchRequest{
		Index: req.Index,
		Query: req.Query,
		ActiveFilters: elastic.ActiveFilters{
			Checkbox: req.ActiveFilters.Checkbox,
			Range:    req.ActiveFilters.Range,
		},
	}
	esResp, err := esReq.Do(ctx, s.esClient)
	if err != nil {
		return nil, err
	}

	var cars []Car
	for _, hit := range esResp.Hits.Hits {
		cars = append(cars, Car{
			ID:    hit.Source.ID,
			Make:  hit.Source.Make,
			Model: hit.Source.Model,
			Price: hit.Source.Price,
			Params: Params{
				Color: hit.Source.Params.Color,
				Year:  hit.Source.Params.Year,
			},
		})
	}

	filters := esResp.Filters()

	return &Response{
		Cars: cars,
		Filters: Filters{
			Checkbox: filters.Checkbox,
			Range:    filters.Range,
		},
	}, nil
}
