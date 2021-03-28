FROM golang:1.16

ADD go.mod go.sum /m/
RUN cd /m; go mod download

WORKDIR /go/pkg/app

CMD [ "go", "run", "./cmd/search/main.go", "-hosts=http://elastic:9200", "-serverBindAddress=:80" ]
