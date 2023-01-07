package kDrive_test

import (
	"context"
	"kDrive"
	"net/http"
	"os"
	"testing"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// newTestClient returns *http.Client with Transport replaced to avoid making real calls
func newTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

// newMockedClient returns *http.Client which responds with content from given file
func newMockedClient(t *testing.T, requestMockFile string, statusCode int) *http.Client {
	return newTestClient(func(*http.Request) *http.Response {
		b, err := os.Open(requestMockFile)
		if err != nil {
			t.Fatal(err)
		}

		resp := &http.Response{
			StatusCode: statusCode,
			Body:       b,
			Header:     make(http.Header),
		}
		return resp
	})
}

func TestRateLimit(t *testing.T) {
	t.Run("List files and directories", func(t *testing.T) {
		c := newTestClient(func(*http.Request) *http.Response {
			return &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{"Retry-After": []string{"0"}},
			}
		}) // .WithHTTPClient(c)
		client := kDrive.NewClient("some_kDrive_Id", "some_token", kDrive.WithHTTPClient(c), kDrive.WithRetry(2))
		_, err := client.File.GetDirectoryAndFiles(context.Background(), "some_file_id")
		if err == nil {
			t.Errorf("Get() error = %v", err)
		}
		wantErr := "Retry request with 429 response failed after 2 retries"
		if err.Error() != wantErr {
			t.Errorf("Get() error = %v, wantErr %s", err, wantErr)
		}
	})

	t.Run("should make maxRetries attempts", func(t *testing.T) {
		attempts := 0
		maxRetries := 2
		c := newTestClient(func(*http.Request) *http.Response {
			attempts++
			return &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{"Retry-After": []string{"0"}},
			}
		})
		client := kDrive.NewClient("some_kDrive_Id", "some_token", kDrive.WithHTTPClient(c), kDrive.WithRetry(maxRetries))
		_, err := client.File.GetDirectoryAndFiles(context.Background(), "some_file_id")
		if err == nil {
			t.Errorf("Get() error = %v", err)
		}
		if attempts != maxRetries {
			t.Errorf("Get() attempts = %v, want %v", attempts, maxRetries)
		}
	})
}
