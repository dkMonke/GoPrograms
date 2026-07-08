// Package main demonstrates using context.Context to control cancellation
// and timeouts for an outgoing HTTP request.
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// fetch issues an HTTP GET request to the given url, bound to ctx so that
// the request is aborted if ctx is canceled or its deadline expires before
// the response is fully read. It returns at most the first 100 bytes of the
// response body as a string. It returns an error if the request cannot be
// constructed, the request fails (including due to context cancellation),
// or the response body cannot be read. The response body is always closed
// before fetch returns.
func fetch(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body[:min(100, len(body))]), nil
}

// main creates a context with a 1-millisecond timeout and uses it to fetch
// "https://example.com" via fetch. Because the timeout is so short, the
// request is expected to be canceled before it completes, so main typically
// prints the resulting error; on success it prints the fetched body instead.
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond) //2*time.Second/
	defer cancel()

	body, err := fetch(ctx, "https://example.com")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(body)

}
