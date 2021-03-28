package main

import (
	"context"
	"flag"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/paintmeyellow/go-elastic-course/internal/indexing"
	"os"
	"strings"
)

func main() {
	fs := flag.NewFlagSet("Indexer", flag.ExitOnError)
	var (
		hosts     = fs.String("hosts", "http://0.0.0.0:9200", "Elasticsearch hosts")
		indexName = fs.String("indexName", "cars", "Elasticsearch index  123 name")
	)
	_ = fs.Parse(os.Args[1:])

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	conf := elasticsearch.Config{
		Addresses: strings.Split(*hosts, ","),
	}
	es, err := elasticsearch.NewClient(conf)
	if err != nil {
		level.Error(logger).Log("error", err)
		os.Exit(1)
	}

	indexService := indexing.NewService(es, logger)
	if err = indexService.ReIndex(context.Background(), *indexName); err != nil {
		level.Error(logger).Log("error", err)
		os.Exit(1)
	}
}
