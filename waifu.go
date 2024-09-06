package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiResponse struct {
	Images []struct {
		URL   string `json:"url"`
		Source string `json:"source"`
	} `json:"images"`
	Message string `json:"message"` // Added for error handling
}

func main() {
	// Replace with your actual API request URL
	apiUrl := "https://api.waifu.im/search"

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers (optional)
	headers := map[string]string{"Accept-Version": "v5"} // Example header
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send the request and handle the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Check for successful response status code
	if resp.StatusCode == http.StatusOK {
		// Unmarshal the JSON response
		var response ApiResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("Error unmarshalling JSON response:", err)
			return
		}

		// Print image URLs and sources (if any)
		if len(response.Images) > 0 {
			fmt.Println("Image results:")
			for _, image := range response.Images {
				fmt.Printf("  URL: %s\n", image.URL)
				fmt.Printf("  Source: %s\n", image.Source)
			}
		} else {
			fmt.Println("No images found in the response.")
		}
	} else {
		// Handle non-200 status codes and potential errors in the response
		fmt.Printf("Error: API request failed with status code %d\n", resp.StatusCode)
		if len(body) > 0 {
			// Check for error message in the response body
			var errResponse ApiResponse
			err = json.Unmarshal(body, &errResponse)
			if err == nil && errResponse.Message != "" {
				fmt.Println("Error message from API:", errResponse.Message)
			}
		}
	}
}