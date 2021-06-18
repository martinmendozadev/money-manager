package utils

import (
	"context"
	"net/http"
	"strings"
)

// SendWithContext send an HTTP request and accepting context
func SendWithContext(url string) (*http.Request, error) {
	ctx := context.Background()
	body := strings.NewReader("Some text")

	// Change NewRequest to NewRequestWithContext and pass context it
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, body)
	if err != nil {
		return req, nil
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return req, nil
	}
	defer req.Body.Close()

	return req, nil
}
