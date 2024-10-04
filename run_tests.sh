#!/bin/bash

# Define log files and target directory
TARGET_DIR="."  # Update this to your test directory
REPORT_DIR="./report"
RESULT_LOG="$REPORT_DIR/result.log"
COVERAGE_FILE="$REPORT_DIR/coverage.out"
COVERAGE_HTML="$REPORT_DIR/coverage.html"
BENCHMARK_LOG="$REPORT_DIR/benchmark.log"

# Create report directory if it doesn't exist
mkdir -p $REPORT_DIR

# Clean previous log files
rm -f $RESULT_LOG $COVERAGE_FILE $COVERAGE_HTML $BENCHMARK_LOG

# Navigate to the target directory
cd $TARGET_DIR || { echo "Directory not found: $TARGET_DIR"; exit 1; }

# Run tests with coverage and verbose output, and log results
echo "Running integration tests in $TARGET_DIR..."
go test -v -coverprofile=$COVERAGE_FILE ./... > $RESULT_LOG 2>&1

# Check if the tests ran successfully
if grep -q "FAIL" $RESULT_LOG; then
    echo "Some tests failed. Check the log for details."
else
    echo "Tests completed successfully."
fi

# Generate HTML coverage report
echo "Generating coverage report..."
go tool cover -html=$COVERAGE_FILE -o $COVERAGE_HTML >> $RESULT_LOG 2>&1

# Optional: Run benchmarks and log results
echo "Running benchmarks..."
go test -bench=. ./... > $BENCHMARK_LOG 2>&1

echo "All tasks completed. Starting the main application..."