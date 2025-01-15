#!/bin/bash

# Function to stop all processes on exit
cleanup() {
    echo "Stopping all processes..."
    kill $(jobs -p) 2>/dev/null
    exit
}

# Set up cleanup on script exit
trap cleanup EXIT

# Create database if it doesn't exist
mariadb -h 127.0.0.1 -P 3307 -u root -proot -e "CREATE DATABASE IF NOT EXISTS acheisuacara;"

# Start the backend
echo "Starting backend..."
go run cmd/api/main.go &

# Start the frontend
echo "Starting frontend..."
cd frontend && pnpm dev &

# Wait for any process to exit
wait

# Cleanup will be called automatically on exit 