package main

import (
	"fmt"
	"log"
)

func main() {
	// Define vertex properties
	properties := map[string]interface{}{
		"name":   "Aliceeee",
		"age":    36,
		"status": "active",
	}

	// Add a vertex using the Gremlin client
	_, err := AddVertex("person", properties)
	if err != nil {
		log.Fatalf("Failed to add vertex: %v\n", err)
	}

	fmt.Println("Vertex added successfully")
}
