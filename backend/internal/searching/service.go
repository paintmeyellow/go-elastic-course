package searching

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
)

type Service interface {
	Search(ctx context.Context, req Request) ([]Car, error)
}

type service struct {
	esClient *elasticsearch.Client
	logger   log.Logger
}

func NewService(es *elasticsearch.Client, l log.Logger) Service {
	return &service{esClient: es, logger: l}
}

func (s service) Search(ctx context.Context, req Request) ([]Car, error) {
	var buf bytes.Buffer

	body := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	if req.Query != "" {
		body = map[string]interface{}{
			"query": map[string]interface{}{
				"multi_match": map[string]interface{}{
					"query":  req.Query,
					"fields": []string{"make", "model"},
				},
			},
		}
	}

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, err
	}

	esResp, err := s.esClient.Search(
		s.esClient.Search.WithIndex(req.Index),
		s.esClient.Search.WithBody(&buf),
		s.esClient.Search.WithContext(ctx),
	)
	if err != nil {
		return nil, err
	}

	defer esResp.Body.Close()

	if esResp.IsError() {
		var errRes struct {
			Error struct {
				Type   string `json:"type"`
				Reason string `json:"reason"`
			} `json:"error"`
		}

		if err = json.NewDecoder(esResp.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		return nil, errors.Errorf(" - [%s] %s: %s",
			esResp.Status(),
			errRes.Error.Type,
			errRes.Error.Reason,
		)
	}

	var resp Response

	err = json.NewDecoder(esResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	var cars []Car
	for _, hit := range resp.Hits.Hits {
		cars = append(cars, Car{
			ID:    hit.Source.ID,
			Make:  hit.Source.Make,
			Model: hit.Source.Model,
			Params: Params{
				Color: hit.Source.Params.Color,
				Year:  hit.Source.Params.Year,
				Price: hit.Source.Params.Price,
			},
		})
	}

	return cars, nil
}
