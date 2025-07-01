#!/bin/bash

for dir in day*/; do
  if [[ -f "$dir/main.go" ]]; then
    echo "Running $dir..."
    go run "$dir/main.go" > "$dir/output.txt"
    echo "Saved output to $dir/output.txt ✅"
  else
    echo "⚠️ Skipping $dir (no main.go found)"
  fi
done
