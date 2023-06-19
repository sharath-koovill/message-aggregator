package mock

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

// getTimeNow returns the current date time in format 2006-01-02T15:04:05-000Z
func getTimeNow() string {
	now := time.Now()
	return now.Format("2006-01-02T15:04:05-000Z")
}

// Create a mock server to intercept API calls
func MockServer() *httptest.Server {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mocking the API response for twitter and mastodon api
		if r.URL.Path == "/twitter/conversations" {
			// Writing the mock response for twitter
			rawResponse := `{"id": "messageid123",
						"senderId": "userId123",
						"text": "Hi, i am sendng this message from twitter ..",
						"eventType": "MessageCreate",
						"messageSource": "twitter",
						"createdAt": %s,
						}`
			response := fmt.Sprintf(string(rawResponse), getTimeNow())
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		} else if r.URL.Path == "/mastodon/conversations" {
			// Writing the mock response for mastodon
			rawResponse := `{"id": "messageid345",
						"senderId": "userId345",
						"text": "Hi, i am sendng this message from mastodon..",
						"eventType": "MessageCreate",
						"messageSource": "mastodon",
						"createdAt": %s,
						}`
			response := fmt.Sprintf(string(rawResponse), getTimeNow())
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		} else {
			log.Printf("Api path Not Found: %s", r.URL.Path)
			http.Error(w, "Api path Not Found", http.StatusNotFound)
		}

	}))
	return mockServer
}
