setup:
    pre-commit install
    pre-commit run --all-files

up:
    docker compose up -d
