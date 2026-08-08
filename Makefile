.PHONY: setup lint test build generate

setup:
	git config core.hooksPath githooks
	chmod +x githooks/*
	npm install

lint:
	golangci-lint config verify -c lint/golangci.yml
	npx --yes @redocly/cli lint openapi/public-v1.yaml
	npx --yes ajv-cli compile -s "jobs/*.json" --spec=draft2019 --validate-formats=false
	npx --yes ajv-cli compile -s "events/*.json" --spec=draft2019 --validate-formats=false

test: build
	go test ./...

build: generate
	go build ./...

generate:
	./codegen/generate.sh
