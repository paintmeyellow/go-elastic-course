package main

import (
	"flag"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/paintmeyellow/go-elastic-course/internal/http/rest"
	"github.com/paintmeyellow/go-elastic-course/internal/searching"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	fs := flag.NewFlagSet("Indexer", flag.ExitOnError)
	var (
		hosts             = fs.String("hosts", "http://0.0.0.0:9200", "Elasticsearch hosts")
		serverBindAddress = fs.String("serverBindAddress", ":8080", "Server address")
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

	searchService := searching.NewService(es, logger)

	searchHandler := kithttp.NewServer(
		rest.MakeSearchEndpoint(searchService),
		rest.DecodeSearchHTTPRequest,
		kithttp.EncodeJSONResponse,
	)

	router := mux.NewRouter()
	router.Use(accessControl)
	router.Handle("/search/{index}", searchHandler).Methods(http.MethodGet)

	chanErrors := make(chan error, 2)

	go func(ch chan error, l log.Logger, ss searching.Service) {
		logger.Log("message", "server is running")
		ch <- http.ListenAndServe(*serverBindAddress, router)
	}(chanErrors, logger, searchService)

	go func(ch chan error) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		ch <- fmt.Errorf("%s", <-c)
	}(chanErrors)

	level.Error(logger).Log("terminated", <-chanErrors)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
