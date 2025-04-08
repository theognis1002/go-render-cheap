package render

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Client represents a Render API client
type Client struct {
	APIKey string
	client *http.Client
}

// APIError represents an error response from the Render API
type APIError struct {
	StatusCode int
	Message    string
	Body       string
}

func (e *APIError) Error() string {
	if e.Body != "" {
		return fmt.Sprintf("%d %s: %s", e.StatusCode, e.Message, e.Body)
	}
	return fmt.Sprintf("%d %s", e.StatusCode, e.Message)
}

// NewClient creates a new Render API client
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// isDatabase checks if the service ID is a database (starts with "dpg-")
func isDatabase(serviceID string) bool {
	return strings.HasPrefix(serviceID, "dpg-")
}

// SuspendService suspends a single Render service or database
func (c *Client) SuspendService(serviceID string) error {
	var url string
	if isDatabase(serviceID) {
		url = fmt.Sprintf("https://api.render.com/v1/databases/%s/suspend", serviceID)
	} else {
		url = fmt.Sprintf("https://api.render.com/v1/services/%s/suspend", serviceID)
	}
	return c.sendRequest(url, serviceID, "suspend")
}

// ResumeService resumes a single Render service or database
func (c *Client) ResumeService(serviceID string) error {
	var url string
	if isDatabase(serviceID) {
		url = fmt.Sprintf("https://api.render.com/v1/databases/%s/resume", serviceID)
	} else {
		url = fmt.Sprintf("https://api.render.com/v1/services/%s/restart", serviceID)
	}
	return c.sendRequest(url, serviceID, "resume")
}

// sendRequest sends the HTTP request to the Render API
func (c *Client) sendRequest(url, serviceID, action string) error {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request for %s: %v", serviceID, err)
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request for %s: %v", serviceID, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    resp.Status,
			Body:       string(body),
		}
	}

	return nil
}
