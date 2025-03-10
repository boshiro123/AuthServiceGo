#!/bin/bash

if ! [ -x "$(command -v protoc)" ]; then
  echo 'Error: protoc is not installed.' >&2
  exit 1
fi

SCHEMA_DIR="./pkg/proto/schemas"
BASE_OUTPUT_DIR="./pkg/proto/generated"

mkdir -p $BASE_OUTPUT_DIR

for PROTO_FILE in $(find $SCHEMA_DIR -name "*.proto"); do
  FILE_NAME=$(basename $PROTO_FILE)
  FILE_NAME_NO_EXT="${FILE_NAME%.*}"
  OUTPUT_DIR="$BASE_OUTPUT_DIR/$FILE_NAME_NO_EXT"

  mkdir -p $OUTPUT_DIR

  protoc \
   --go_out=$OUTPUT_DIR \
   --go_opt=paths=source_relative \
   --proto_path=$SCHEMA_DIR \
   --go-grpc_out=$OUTPUT_DIR \
   --go-grpc_opt=paths=source_relative \
   $PROTO_FILE
done

if find "$BASE_OUTPUT_DIR" -name "*.pb.go" | read; then
  echo "Found .pb.go files in $BASE_OUTPUT_DIR"
else
  echo "No .pb.go files found in $BASE_OUTPUT_DIR"
  exit 1
fi