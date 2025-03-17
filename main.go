package main

import (
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
)

// TicketsRequest represents the request payload
type TicketsRequest struct {
	Tickets [][]string `json:"tickets"`
}

// ItineraryResponse represents the response payload
type ItineraryResponse struct {
	Itinerary []string `json:"itinerary"`
}

// reconstructItinerary reconstructs the complete itinerary from the given tickets
func reconstructItinerary(tickets [][]string) []string {
	// Create a map of source to destinations (sorted lexicographically)
	graph := make(map[string][]string)

	// Track all airports to find potential starting points
	allAirports := make(map[string]bool)

	// Populate the graph
	for _, ticket := range tickets {
		source := ticket[0]
		dest := ticket[1]
		graph[source] = append(graph[source], dest)
		allAirports[source] = true
		allAirports[dest] = true
	}

	// Sort destinations for each source in lexicographical order
	// This ensures we always pick the lexicographically smallest destination first
	for source := range graph {
		destinations := graph[source]
		sort.Strings(destinations)
		graph[source] = destinations
	}

	// Find the starting point
	// Find a node with one more outgoing edge than incoming edges
	// If no such node exists, pick any node
	start := ""

	// Count incoming edges for each airport
	incomingEdges := make(map[string]int)
	for _, destinations := range graph {
		for _, dest := range destinations {
			incomingEdges[dest]++
		}
	}

	// Find a node with more outgoing than incoming edges
	for airport := range allAirports {
		outgoing := len(graph[airport])
		incoming := incomingEdges[airport]
		if outgoing > incoming {
			start = airport
			break
		}
	}

	// If no such node found, pick any node
	if start == "" {
		for airport := range allAirports {
			start = airport
			break
		}
	}

	// Initialize the itinerary with the starting point
	itinerary := []string{}

	// DFS function to find the Eulerian path
	var dfs func(airport string)
	dfs = func(airport string) {
		// While there are outgoing edges from this airport
		for len(graph[airport]) > 0 {
			// Get the next destination (lexicographically smallest)
			nextAirport := graph[airport][0]
			// Remove this edge from the graph
			graph[airport] = graph[airport][1:]
			// Recursively visit the next airport
			dfs(nextAirport)
		}
		// Add the current airport to the itinerary (in reverse order)
		itinerary = append([]string{airport}, itinerary...)
	}

	// Start DFS from the starting point
	dfs(start)

	return itinerary
}

func main() {
	e := echo.New()

	// Add the POST endpoint for reconstructing itinerary
	e.POST("/itinerary", func(c echo.Context) error {
		req := new(TicketsRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		itinerary := reconstructItinerary(req.Tickets)
		return c.JSON(http.StatusOK, ItineraryResponse{Itinerary: itinerary})
	})

	e.Logger.Fatal(e.Start(":9080"))
}
