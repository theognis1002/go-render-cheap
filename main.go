package main

import (
	"fmt"
	"os"
	"strings"

	"go-render-services/render"
)

// main is the entry point of the script.
func main() {
	// Check if enough arguments are provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <action>")
		fmt.Println("action: suspend or unsuspend")
		fmt.Println("Environment variables:")
		fmt.Println("  RENDER_API_KEY: Your Render API key")
		fmt.Println("  RENDER_SERVICE_IDS: Comma-separated list of service IDs")
		os.Exit(1)
	}

	// Parse command-line arguments
	action := os.Args[1]

	// Validate the action
	if action != "suspend" && action != "unsuspend" {
		fmt.Println("Invalid action:", action)
		fmt.Println("Usage: go run script.go <action>")
		fmt.Println("action: suspend or unsuspend")
		os.Exit(1)
	}

	// Get the API key from the environment
	apiKey := os.Getenv("RENDER_API_KEY")
	if apiKey == "" {
		fmt.Println("RENDER_API_KEY environment variable is not set")
		os.Exit(1)
	}

	// Get service IDs from environment variable
	serviceIDsStr := os.Getenv("RENDER_SERVICE_IDS")
	if serviceIDsStr == "" {
		fmt.Println("RENDER_SERVICE_IDS environment variable is not set")
		os.Exit(1)
	}

	// Split service IDs by comma and trim whitespace
	serviceIDs := strings.Split(serviceIDsStr, ",")
	for i, id := range serviceIDs {
		serviceIDs[i] = strings.TrimSpace(id)
	}

	client := render.NewClient(apiKey)

	// Process each service ID
	for _, serviceID := range serviceIDs {
		var err error
		if action == "suspend" {
			err = client.SuspendService(serviceID)
		} else {
			err = client.ResumeService(serviceID)
		}

		if err != nil {
			fmt.Printf("Error with service %s: %v\n", serviceID, err)
			continue
		}

		if action == "suspend" {
			fmt.Printf("Successfully suspended service %s\n", serviceID)
		} else {
			fmt.Printf("Successfully initiated restart for service %s\n", serviceID)
		}
	}
}
