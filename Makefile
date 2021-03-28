init: docker-down docker-build docker-up index-data
up: docker-down docker-up
down: docker-down

docker-up:
	docker-compose up -d
docker-down:
	docker-compose down --remove-orphans -v
docker-build:
	docker-compose build

search-restart:
	docker-compose restart search

index-data: wait-for-elastic
	docker-compose run --rm indexer go run cmd/indexer/main.go -hosts="http://elastic:9200" -indexName="cars"

wait-for-elastic:
	timeout 30 sh -c "until curl --silent --output /dev/null http://127.0.0.1:9200/_cat/health?h=st; do printf '.'; sleep 2; done; printf '\n'"
