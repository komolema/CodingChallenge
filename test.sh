#!/bin/bash

# Start the server in the background
go run main.go &
SERVER_PID=$!

# Wait for the server to start
sleep 2

# Test the /itinerary endpoint with the example input
echo "Testing /itinerary endpoint with example input..."
curl -X POST -H "Content-Type: application/json" -d '{"tickets":[["LAX","DXB"], ["JFK","LAX"], ["SFO","SJC"], ["DXB","SFO"]]}' http://localhost:1323/itinerary

# Kill the server
kill $SERVER_PID

echo -e "\nTest completed."