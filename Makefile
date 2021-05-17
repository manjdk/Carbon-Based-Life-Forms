deps:
		go get github.com/golangci/golangci-lint/cmd/golangci-lint
		go get -u github.com/go-swagger/go-swagger/cmd/swagger
		go mod download

test:
		go test -cover -failfast ./...

create-db:
		./scripts/create-table.sh

remove-db:
		./scripts/delete-table.sh

manage-queue:
		./scripts/manage-queue.sh