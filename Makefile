BINARY_NAME=awardservice

default: test

install-tools: download
	@echo Installing tools from tools.go
	@go list -f '{{range .Imports}}{{.}} {{end}}' tools.go | xargs go install

build:
	go build -o ./bin/${BINARY_NAME} ./cmd/awarder/main.go

run-docker-compose:
	sudo docker-compose up --build

run-docker-composed:
	sudo docker-compose up --build -d

down-docker-compose:
	sudo docker-compose down

run:
	make build
	./bin/${BINARY_NAME}

clean:
	go clean
	rm ./bin/${BINARY_NAME}

test:
	gotestsum --format testname

test-cov:
	go test ./... -v -short -coverprofile cover.out && \
	go tool cover -html=cover.out

init-db:
	cat ./internal/ports/persistence/sql/schema.sql | psql -h ${POSTGRES_HOST} -p ${POSTGRES_PORT} -U ${POSTGRES_USER} -d ${POSTGRES_NAME} -1

