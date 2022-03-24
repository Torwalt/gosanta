#!/bin/sh

set -o allexport; source .env; set +o allexport
./bin/gosanta_server

