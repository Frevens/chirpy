#!/bin/bash

set -e

echo "🔨 Building Chirpy..."
go build -o out

echo "🚀 Starting Chirpy..."
./out &
SERVER_PID=$!

cleanup() {
    echo
    echo "🛑 Stopping Chirpy..."
    kill "$SERVER_PID" 2>/dev/null || true
}

trap cleanup EXIT

echo "⏳ Waiting for server..."

until curl -s http://localhost:8080/api/healthz > /dev/null; do
    sleep 0.2
done

echo "✅ Server is ready"
echo
echo "🧪 Running Boot.dev tests..."
echo

bootdev run "$@"
