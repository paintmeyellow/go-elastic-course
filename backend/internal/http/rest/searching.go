package rest

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"github.com/paintmeyellow/go-elastic-course/internal/searching"
	"net/http"
)

type searchRequest struct {
	Index string `json:"index"`
	Query string `json:"query"`
}
type searchResponse struct {
	Cars []searching.Car `json:"cars"`
}

func DecodeSearchHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req searchRequest

	vars := mux.Vars(r)
	req.Index = vars["index"]

	req.Query = r.URL.Query().Get("query")

	return req, nil
}

func MakeSearchEndpoint(svc searching.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(searchRequest)

		cars, err := svc.Search(context.Background(), searching.Request{
			Index: req.Index,
			Query: req.Query,
		})
		if err != nil {
			return searchResponse{}, err
		}

		return searchResponse{Cars: cars}, nil
	}
}
