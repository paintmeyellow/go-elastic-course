package rest

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"github.com/paintmeyellow/go-elastic-course/internal/searching"
	"net/http"
)

type searchRequest struct {
	index         string
	Query         string `json:"query"`
	ActiveFilters struct {
		Checkbox map[string][]string `json:"checkbox"`
		Range    map[string]struct {
			Min float64 `json:"min"`
			Max float64 `json:"max"`
		} `json:"range"`
	} `json:"active_filters"`
}

type searchResponse struct {
	Cars    []searching.Car   `json:"cars"`
	Filters searching.Filters `json:"filters"`
}

func DecodeSearchHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req searchRequest

	vars := mux.Vars(r)
	req.index = vars["index"]

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func MakeSearchEndpoint(svc searching.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(searchRequest)

		resp, err := svc.Search(context.Background(), searching.Request{
			Index: req.index,
			Query: req.Query,
			ActiveFilters: searching.ActiveFilters{
				Checkbox: req.ActiveFilters.Checkbox,
				Range:    req.ActiveFilters.Range,
			},
		})
		if err != nil {
			return searchResponse{}, err
		}

		return searchResponse{
			Cars:    resp.Cars,
			Filters: resp.Filters,
		}, nil
	}
}
