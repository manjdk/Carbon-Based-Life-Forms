deps:
		go mod download

test:
		go test -cover -failfast ./...

create-db:
		./scripts/create-table.sh

remove-db:
		./scripts/delete-table.sh

manage-queue:
		./scripts/manage-queue.sh