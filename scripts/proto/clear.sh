#!/usr/bin/env zsh

set -euo pipefail

clear_proto() {
  rm -rf ./pkg/proto
}

clear_proto