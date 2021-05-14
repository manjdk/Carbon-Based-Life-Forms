deps:
		go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.24
		go get -u github.com/go-swagger/go-swagger/cmd/swagger
		go mod download

lint:
		golangci-lint run --config=.golangci.yml ./...
		swagger validate --quiet swagger/openapi.yml

test: lint
		go test -cover -failfast ./...

test-all: lint
		go test -tags integration -cover -failfast -timeout=60s ./...

create-db:
		./scripts/create-table.sh

remove-db:
		./scripts/delete-table.sh

manage-queue:
		./scripts/manage-queue.sh