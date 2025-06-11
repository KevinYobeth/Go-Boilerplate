#!/bin/bash
set -e

mkdir -p pkg/genproto/$service

directory="api/protobuf"

function codegen {
  service="$1"
  output_dir=pkg/genproto

  [ -d "$output_dir/$service" ] || mkdir -p "$output_dir/$service"

  protoc \
    --proto_path=api/protobuf "$directory/$service.proto" \
    "--go_out=$output_dir/$service" --go_opt=paths=source_relative \
    --go-grpc_opt=require_unimplemented_servers=false \
    "--go-grpc_out=$output_dir/$service" --go-grpc_opt=paths=source_relative

  echo "$service.proto generated successfully under $output_dir."
}

for file in "$directory"/*
do
    if [ -f "$file" ]; then
        filename=$(basename "$file")
        filename_no_ext="${filename%.*}"

        codegen $filename_no_ext
    fi
done