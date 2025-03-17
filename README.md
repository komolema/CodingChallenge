# Flight Itinerary Reconstruction

A web application that reconstructs a complete flight itinerary from a sequence of flight tickets.

## Overview

This application provides an HTTP endpoint that accepts a JSON payload containing flight tickets and returns the reconstructed itinerary. Each flight ticket is represented as a pair ["Source", "Destination"], where "Source" is the departure airport code and "Destination" is the arrival airport code.

## Requirements

- Go 1.24 or later
- Echo framework v4.13.3 or later

## Setup and Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/CodingChallenge.git
   cd CodingChallenge
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Run the application:
   ```
   go run main.go
   ```

The server will start on port 8080.

## API Usage

### Reconstruct Itinerary

**Endpoint:** `POST /itinerary`

**Request Payload:**
```json
{
  "tickets": [
    ["LAX", "DXB"],
    ["JFK", "LAX"],
    ["SFO", "SJC"],
    ["DXB", "SFO"]
  ]
}
```

**Response:**
```json
{
  "itinerary": ["JFK", "LAX", "DXB", "SFO", "SJC"]
}
```

## Testing

A test script is provided to verify the functionality:

```
./test.sh
```

This script starts the server, sends a test request, and verifies the response.

## Development Choices

### Framework Selection

The Echo framework was chosen for its simplicity, performance, and ease of use. It provides a clean API for handling HTTP requests and responses, making it ideal for this task.

### Algorithm

The itinerary reconstruction problem is essentially finding an Eulerian path in a directed graph, where each flight ticket represents an edge. The algorithm works as follows:

1. Build a graph where each airport is a node and each flight is a directed edge.
2. Sort the destinations for each source in lexicographical order to ensure we always pick the lexicographically smallest destination first.
3. Find a starting point:
   - Find a node with more outgoing than incoming edges (which is a property of the starting node in an Eulerian path).
   - If no such node exists, pick any node.
4. Use depth-first search (DFS) to find the Eulerian path, adding airports to the itinerary in reverse order.

This approach ensures that we find a valid itinerary if one exists, and it handles edge cases such as multiple flights from the same source airport.

### Error Handling

The application includes basic error handling for invalid request payloads. More comprehensive error handling could be added for production use.

## Future Improvements

- Add more comprehensive error handling
- Implement input validation
- Add logging and monitoring
- Containerize the application for easier deployment
- Add more test cases to cover edge cases
