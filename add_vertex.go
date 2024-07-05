package addVertex

import (
	"fmt"
	"log"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
)

// GremlinClient is a struct that holds the Gremlin client
type GremlinClient struct {
	client *gremlingo.Client
}

// newGremlinClient initializes and returns a new GremlinClient
func newGremlinClient() (*GremlinClient, error) {
	gremlinClient, err := gremlingo.NewClient("ws://localhost:8182/gremlin")
	if err != nil {
		return nil, fmt.Errorf("failed to create Gremlin client: %v", err)
	}
	return &GremlinClient{client: gremlinClient}, nil
}

// AddVertex adds a vertex with the given properties to the graph
func AddVertex(label string, properties map[string]interface{}) ([]map[string]interface{}, error) {
	// Initialize the Gremlin client
	gc, err := newGremlinClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create Gremlin client: %v", err)
	}
	// Ensure the client is closed when the function returns
	defer gc.client.Close()

	// Construct the Gremlin traversal string
	traversalString := fmt.Sprintf(`g.addV('%s')`, label)
	for key, value := range properties {
		traversalString += fmt.Sprintf(`.property('%s', '%v')`, key, value)
	}

	// Execute the traversal
	res, err := gc.client.Submit(traversalString)
	if err != nil {
		return nil, fmt.Errorf("failed to execute traversal: %v", err)
	}
	var resultsMap []map[string]interface{}
	// results := res.GetStatusAttributes()
	log.Printf("ffetre %v", res)
	return resultsMap, nil
}
