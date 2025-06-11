setup:
    pre-commit install
    pre-commit run --all-files

build:
    go build -o ./bin/awarder cmd/awarder/main.go

up:
    docker compose up -d
    just build
    ./bin/awarder

reset-db:
    docker compose down -v
    docker compose up -d

