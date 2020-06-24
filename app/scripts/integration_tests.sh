#!/bin/bash

source "./scripts/$ENV_FILE"

go test -tags integration ./...
