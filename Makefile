IMAGE = github.com/7phs/coding-challenge-search
VERSION = latest

build: export GO111MODULE=on
build:
	go mod vendor
	go build -o search-service -ldflags "-X github.com/7phs/coding-challenge-search/cmd.BuildTime=`date +%Y-%m-%d:%H:%M:%S` -X github.com/7phs/coding-challenge-search/cmd.GitHash=`git rev-parse --short HEAD`"

testing: export GO111MODULE=on
testing:
	go mod vendor
	LOG_LEVEL=error ADDR=:8080 go test ./...

run: export GO111MODULE=on
run:
	go mod vendor
	LOG_LEVEL=info ADDR=:8080 STAGE=production go run main.go run

image:
	docker build -t $(IMAGE):$(VERSION) .

image-run:
	docker run --rm -it -p 8080:8080 $(IMAGE):$(VERSION)

all: build
