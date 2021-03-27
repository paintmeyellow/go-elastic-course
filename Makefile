up: docker-down docker-up search-up
down: docker-down

docker-up:
	docker-compose up -d elastic kibana frontend
docker-down:
	docker-compose down --remove-orphans -v

search-up:
	docker-compose run  --rm -p 8081:80 golang go run cmd/searcher/main.go \
		 -hosts="http://elastic:9200" \
		 -serverBindAddress=":80"

reindex:
	docker-compose run --rm golang go run cmd/indexer/main.go -hosts="http://elastic:9200"
