.PHONY: build clean deploy lint install run

install:
	go mod download
	npm i

build:
	sh ./build.sh

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
