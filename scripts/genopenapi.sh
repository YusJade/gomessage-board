#!/usr/bin/env bash

set -euo pipefail

shopt -s globstar

source ./scripts/utils.sh

if ! [[ "$0" =~ scripts/genopenapi.sh ]]; then
  log_error "must be run from repository root"
  exit 255
fi

OPENAPI_ROOT="./api/openai"

GEN_SERVER=(
  # "chi-server"
  # "echo-server"
  "gin-server"
)


if [ "${#GEN_SERVER[@]}" -ne 1 ]; then
  log_error "GEN_SERVER enables more than 1 server, plz check"
  exit 255
fi


log_info "using ${GEN_SERVER[0]}"



function openapi_files {
  openapi_files=$(ls $OPENAPI_ROOT)
  echo "${openapi_files[@]}"
}


# output_dir, package_name, service_name
function gen() {
  local output_dir=$1
  local package=$2
  local service=$3

  run mkdir -p "$output_dir"
  run find "$output_dir" -type f -name "*.gen.go" -delete

  prepare_dir "common/client/$service"

  run oapi-codegen -generate types -o "$output_dir/openapi_types.gen.go" -package "$package" "api/$service.yaml"
  run oapi-codegen -generate "${GEN_SERVER[0]}" -o "$output_dir/openapi_api.gen.go" -package "$package" "api/$service.yaml"

  run oapi-codegen -generate client -o "common/client/$service/openapi_client.gen.go" -package "$service" "api/$service.yaml"
  run oapi-codegen -generate types -o "common/client/$service/openapi_types.gen.go" -package "$service" "api/$service.yaml"
}

gen message/ports ports message
