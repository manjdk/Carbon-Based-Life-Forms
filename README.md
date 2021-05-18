# Carbon-Based-Life-Forms

AWS should be configured on local machine.

Steps to run a program:
- checkout repository
- use `docker-compose up -d` to run AWS services locally
- use `make manage-queue` to create SQS locally
- use `make create-db` to create database instance locally
- use `go run main.go` to run entry point of api (default port `:8181`)
- use `go run manager/main.go` to run manager (default port `:8282`)
- use ` go run factory/main.go` to run factory (default port `:8383`)
- in order to change default ports flag `-p` can be used. Example: `go run main.go -p 8000`
- when `go run main.go` will be running, documentation can be found in WEB browser. Address `http://localhost:8181/api/doc.html`
- `make test` runs all the unit tests