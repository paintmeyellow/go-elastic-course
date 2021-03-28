package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
)

type SearchRequest struct {
	Index string
	Query string
}

func (req SearchRequest) Do(ctx context.Context, es *elasticsearch.Client) (*SearchResponse, error) {
	body := map[string]interface{}{
		"query": req.query(),
		"aggs":  req.aggs(),
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, err
	}
	resp, err := es.Search(
		es.Search.WithIndex(req.Index),
		es.Search.WithBody(&buf),
		es.Search.WithContext(ctx),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		var errResp struct {
			Error struct {
				Type   string `json:"type"`
				Reason string `json:"reason"`
			} `json:"error"`
		}
		if err = json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return nil, err
		}
		return nil, errors.Errorf(" - [%s] %s: %s",
			resp.Status(),
			errResp.Error.Type,
			errResp.Error.Reason,
		)
	}

	var decodeResp SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&decodeResp)
	if err != nil {
		return nil, err
	}

	return &decodeResp, nil
}

func (req SearchRequest) query() map[string]interface{} {
	var query map[string]interface{}

	switch {

	case req.Query == "":
		query = map[string]interface{}{
			"match_all": map[string]interface{}{},
		}

	case req.Query != "":
		query = map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  req.Query,
				"fields": []string{"make", "model"},
			},
		}
	}

	return query
}

func (req SearchRequest) aggs() map[string]interface{} {
	var activeFilters []map[string]interface{}

	return map[string]interface{}{
		"min_price": map[string]interface{}{
			"filter": map[string]interface{}{
				"bool": map[string]interface{}{
					"filter": activeFilters,
				},
			},
			"aggs": map[string]interface{}{
				"val": map[string]interface{}{
					"min": map[string]interface{}{
						"field": "params.price",
					},
				},
			},
		},
		"max_price": map[string]interface{}{
			"filter": map[string]interface{}{
				"bool": map[string]interface{}{
					"filter": activeFilters,
				},
			},
			"aggs": map[string]interface{}{
				"val": map[string]interface{}{
					"max": map[string]interface{}{
						"field": "params.price",
					},
				},
			},
		},
		"color": map[string]interface{}{
			"filter": map[string]interface{}{
				"bool": map[string]interface{}{
					"filter": activeFilters,
				},
			},
			"aggs": map[string]interface{}{
				"vars": map[string]interface{}{
					"terms": map[string]interface{}{
						"field": "params.color",
					},
				},
			},
		},
		"year": map[string]interface{}{
			"filter": map[string]interface{}{
				"bool": map[string]interface{}{
					"filter": activeFilters,
				},
			},
			"aggs": map[string]interface{}{
				"vars": map[string]interface{}{
					"terms": map[string]interface{}{
						"field": "params.year",
					},
				},
			},
		},
	}
}
