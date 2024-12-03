#!/bin/bash

# 表名 all 为当前连接库中的所有表
tableName=$1

PROJECT_DIR=$(dirname "$0")
GENERATE_DIR="$PROJECT_DIR/cmd/gen"

cd "$GENERATE_DIR" || exit

echo "Start Generating"
go run . "$tableName"