.PHONY: build clean deploy format build_sam_template install run

install:
	go mod download
	npm i

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/users/create src/users/create.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

build_sam_template:
	sls sam export --output ./template.yml

run: build build_sam_template
	sam local start-api

deploy: clean build
	sls deploy --verbose

format: 
	golangci-lint run
