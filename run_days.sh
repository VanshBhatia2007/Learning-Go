#!/bin/bash

# Loop through each "dayXX" folder (e.g., day01, day02, day03)
for day_folder in day*/; do
  echo "📁 Searching in $day_folder..."

  # Find ALL main.go files recursively inside that day folder
  find "$day_folder" -type f -name "main.go" | while read -r main_file; do
    # Get directory of the main.go
    dir_of_main=$(dirname "$main_file")
    output_file="$dir_of_main/output.txt"

    echo "▶️ Running $main_file..."
    go run "$main_file" > "$output_file"
    echo "✅ Output saved to $output_file"
  done
done
