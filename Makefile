.PHONY: build clean deploy lint install run

install:
	go mod download
	npm i

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/users/create src/users/create.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/utils/headers src/utils/headers.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock ./.serverless ./template.yml

build_sam_template:
	sls sam export --output ./template.yml

run: build build_sam_template
	sam local start-api

deploy: clean build
	sls deploy --verbose

lint: 
	golangci-lint run
