reindex:
	docker-compose run --rm golang go run cmd/indexer/main.go -hosts="http://elastic:9200"