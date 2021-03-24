package indexing

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var ErrDeleteIndex = errors.New("couldn't delete index")
var ErrCreateIndex = errors.New("couldn't create index")
var ErrIndexCars = errors.New("couldn't index")
var ErrReadCarsFile = errors.New("couldn't read cars file")
var ErrEncodeCars = errors.New("couldn't encode cars")

type Service interface {
	ReIndex(ctx context.Context, indexName string) error
}

type service struct {
	esClient *elasticsearch.Client
	logger   log.Logger
}

func NewService(es *elasticsearch.Client, l log.Logger) *service {
	return &service{esClient: es, logger: l}
}

func (s service) ReIndex(ctx context.Context, indexName string) error {
	//delete old index
	_, err := s.esClient.Indices.Delete([]string{indexName})
	if err != nil {
		return ErrDeleteIndex
	}

	//create new index
	body := `{
		"mappings": {
			"properties": {
				"make": { "type": "text" },
				"model": { "type": "text" },
				"params": {
					"properties": {
						"color": { "type": "keyword" },
						"year": { "type": "keyword" },
						"price": { "type": "scaled_float", "scaling_factor": 100 }
					}
				}
			}
		}
	}`

	res, err := s.esClient.Indices.Create(indexName,
		s.esClient.Indices.Create.WithBody(strings.NewReader(body)),
	)
	if err != nil || res.IsError() {
		return ErrCreateIndex
	}

	//reading file
	jsonFile, err := os.Open("resources/cars.json")
	if err != nil {
		return ErrReadCarsFile
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return ErrReadCarsFile
	}

	//encoding cars
	var cars []Car
	if err = json.Unmarshal(byteValue, &cars); err != nil {
		return ErrEncodeCars
	}

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:      indexName,
		Client:     s.esClient,
		FlushBytes: 4e+3,
	})
	if err != nil {
		return ErrIndexCars
	}

	//indexing cars
	for _, car := range cars {
		carBytes, err := json.Marshal(car)
		if err != nil {
			level.Error(s.logger).Log("error", err)
			continue
		}

		err = bi.Add(ctx, esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: strconv.Itoa(car.ID),
			Body:       bytes.NewReader(carBytes),
			OnSuccess: func(
				ctx context.Context,
				item esutil.BulkIndexerItem,
				res esutil.BulkIndexerResponseItem,
			) {
				fmt.Printf("\r[%d] %s test/%s\n", res.Status, res.Result, item.DocumentID)
			},
		})

		if err != nil {
			level.Error(s.logger).Log("error", err)
			continue
		}
	}

	if err := bi.Close(context.Background()); err != nil {
		return err
	}

	//check stats
	stats := bi.Stats()
	if stats.NumFailed > 0 {
		msg := fmt.Sprintf("Indexed [%d] documents with [%d] errors", stats.NumFlushed, stats.NumFailed)
		level.Error(s.logger).Log("msg", msg)
		return nil
	}

	msg := fmt.Sprintf("Successfully indexed [%d] documents. Done %d requests.", stats.NumFlushed, stats.NumRequests)
	level.Info(s.logger).Log("msg", msg)

	return nil
}
