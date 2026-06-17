package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/LearnWithSrini/go_pdf_report_generator/models"
)

type BackendClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewBackendClient(baseURL string, timeout int) *BackendClient {
	return &BackendClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

// GetStudent fetches student data from the Node.js backend API
func (c *BackendClient) GetStudent(id string) (*models.Student, error) {
	url := fmt.Sprintf("%s/api/v1/students/%s", c.baseURL, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch student data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("backend API returned status %d: %s", resp.StatusCode, string(body))
	}

	var apiResponse models.APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if !apiResponse.Success {
		return nil, fmt.Errorf("API returned error: %s", apiResponse.Message)
	}

	return &apiResponse.Data, nil
}
