BINARY_NAME=awardservice


build:
	go build -o ./bin/ranker ./cmd/rankingservice/main.go
	go build -o ./bin/eventlogger ./cmd/eventreader/main.go
	go build -o ./bin/awarder ./cmd/awarding/main.go

build-docker:
	sudo docker build \
		--build-arg HTTP_PORT=${HTTP_PORT} \
		--build-arg BINARY_NAME=${BINARY_NAME} . -t ${BINARY_NAME}

run-docker:
	sudo docker run -p ${HTTP_PORT}:${HTTP_PORT} ${BINARY_NAME}

run-dockerd:
	sudo docker run -p ${HTTP_PORT}:${HTTP_PORT} -d ${BINARY_NAME}

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
	rm ./bin/ranker
	rm ./bin/eventlogger
	rm ./bin/awarder

test:
	go test ./... -v -short

test-cov:
	go test ./... -v -short -coverprofile cover.out && \
	go tool cover -html=cover.out

init-db:
	cat ./internal/ports/persistence/sql/schema.sql | psql -h ${POSTGRES_HOST} -p ${POSTGRES_PORT} -U ${POSTGRES_USER} -d ${POSTGRES_NAME} -1


