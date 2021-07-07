package utils

import (
	"context"
	"io"
	"log"
	"net/http"
)

// SendWithContext send an HTTP request and accepting context
func SendWithContext(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		log.Panicf("Cannot create request: %s\n", err)
		return req, err
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicf("Cannot response request: %s\n", err)
		return req, err
	}
	defer rsp.Body.Close()

	return req, nil
}
