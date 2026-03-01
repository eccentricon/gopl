#!/bin/zsh

# Move to the script's directory so output files stay here
cd "$(dirname "$0")"

echo "Running Echo1 (Loop) benchmarks..."
go test -bench=. ../echo1 -count=10 > old.txt

echo "Running Echo3 (Join) benchmarks..."
go test -bench=. ../echo3 -count=10 > new.txt

echo "--- Results ---"
benchstat old.txt new.txt
