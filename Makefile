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
	curl -L "https://app.swaggerhub.com/apiproxy/registry/BaSyx/basyx_asset_administration_shell_repository_http_rest_api/v1?resolved=true&flatten=true&pretty=true" > api/api.yaml

api/models.yaml:
	curl -L https://github.com/admin-shell-io/aas-specs/raw/master/schemas/yaml/aas-openapi.yaml > api/models.yaml
	cd api && git apply models.patch

api/merged.yaml: api/api.yaml api/models.yaml
	go run ./tools/merge.go ./api/models.yaml ./api/api.yaml > api/merged.yaml

.PHONY: api
api: api/merged.yaml
	./tools/generate.sh
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
	go mod tidy -compat=1.19
	go fmt ./api/...

