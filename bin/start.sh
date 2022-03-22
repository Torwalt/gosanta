#!/bin/sh

set -o allexport; source .env; set +o allexport
./bin/gobreach_server

