.PHONY: fmt
## fmt: format the source code using gofmt
fmt:
	@go fmt ./...

.PHONY: test
## test: runs go test with default values
test:
	go test -failfast -v -timeout 5m ./...

.PHONY: test-cover-report
## test-cover-report: generates a HTML coverage report
test-cover-report:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

.PHONY: lint
## lint: lint using staticcheck
lint:
	staticcheck ./...

.PHONY: clean
## clean: cleans the working dir
clean:
	@echo "Cleaning..."
	@- rm -rf build
	@go clean -cache

.PHONY: help
## help: Prints this help message
help:
	@echo "                               _____..---========+*+==========---.._____\n  ______________________ __,-='=====____  =================== _____=====\`=\n (._____________________I__) - _-=_/    \`---------=+=--------'\n     /      /__...---===='---+---_'\n    '------'---.___ -  _ =   _.-'\n                    \`--------'"
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

api/api.yaml:
	mkdir -p api
	curl -L "https://app.swaggerhub.com/apiproxy/registry/BaSyx/basyx_asset_administration_shell_http_rest_api/v1?resolved=true&flatten=true&pretty=true" > api/api.yaml

.PHONY: api
api: api/api.yaml
	mkdir -p api
	./generate.sh
	rm -rf \
		api/.gitignore \
		api/.openapi-generator \
		api/.travis.yml \
		api/git_push.sh \
		api/go.mod api/go.sum \
		api/.openapi-generator-ignore \
		api/README.md \
		api/docs \
		api/api
	go mod tidy -compat=1.17
	go fmt ./api/...

