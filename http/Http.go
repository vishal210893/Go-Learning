package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	// Make an HTTP GET request
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println(os.Stderr, "Error fetching URL: %v\n", err)
		os.Exit(1)
	}
	// Ensure the response body is closed after we're done
	defer resp.Body.Close()

	// Check for non-200 HTTP status codes
	if resp.StatusCode != http.StatusOK {
		fmt.Println(os.Stderr, "Non-OK HTTP status: %s\n", resp.Status)
		os.Exit(1)
	}

	// Read the response body efficiently
	//body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(os.Stderr, "Error reading response body: %v\n", err)
		os.Exit(1)
	}

	// Output the response body as a string
	//fmt.Println(string(body))

	fmt.Println("-------------------")
	written, err := io.Copy(logWriter{}, resp.Body)
	if err == nil {
		fmt.Println("Wrote down", written, "bytes")
	}
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
