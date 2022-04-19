#!/usr/bin/env zsh

set -euo pipefail

generate_proto() {
  protoc --go_out=./ --go-grpc_out=./ ./**/*.proto
}

generate_proto