package render

import (
	"fmt"
	"io"
	"net/http"
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

// SuspendService suspends a single Render service or database
func (c *Client) SuspendService(serviceID string) error {
	url := fmt.Sprintf("https://api.render.com/v1/services/%s/suspend", serviceID)
	return c.sendRequest(url, serviceID, "suspend")
}

// ResumeService resumes a single Render service or database
func (c *Client) ResumeService(serviceID string) error {
	url := fmt.Sprintf("https://api.render.com/v1/services/%s/resume", serviceID)
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

	// Consider both 200 OK and 202 Accepted as success
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		body, _ := io.ReadAll(resp.Body)
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    resp.Status,
			Body:       string(body),
		}
	}

	return nil
}
