#!/bin/bash

if [ "$COME_RUN_MIGRATION" == "1" ]; then
  echo "Running Migration";
  go run cmd/migration/main.go
fi

go run cmd/api/main.go