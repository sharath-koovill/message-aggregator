package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestLookupEnvOrString(t *testing.T) {
	// Test when key exists
	os.Setenv("TEST_KEY", "test_value")
	val := LookupEnvOrString("TEST_KEY", "default_value")
	if val != "test_value" {
		t.Errorf("LookupEnvOrString failed, expected %s but got %s", "test_value", val)
	}

	// Test when key does not exist
	os.Unsetenv("TEST_KEY")
	val = LookupEnvOrString("TEST_KEY", "default_value")
	if val != "default_value" {
		t.Errorf("LookupEnvOrString failed, expected %s but got %s", "default_value", val)
	}
}

func TestLookupEnvOrInt(t *testing.T) {
	// Test when key exists
	os.Setenv("TEST_KEY", "10")
	val := LookupEnvOrInt("TEST_KEY", 5)
	if val != 10 {
		t.Errorf("LookupEnvOrInt failed, expected %d but got %d", 10, val)
	}

	// Test when key does not exist
	os.Unsetenv("TEST_KEY")
	val = LookupEnvOrInt("TEST_KEY", 5)
	if val != 5 {
		t.Errorf("LookupEnvOrInt failed, expected %d but got %d", 5, val)
	}

	// Test when key exists but value is not an integer
	os.Setenv("TEST_KEY", "not_an_integer")
	defer os.Unsetenv("TEST_KEY")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("LookupEnvOrInt did not panic")
		}
	}()
	LookupEnvOrInt("TEST_KEY", 5)
}

func TestPollAPI(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mocking the API response for twitter and mastodon api
		if r.URL.Path == "/twitter/converations" || r.URL.Path == "/mastodon/converations" {
			rawResponse := `{"id": "messageid123",
						"senderId": "userId123",
						"text": "Hi, i am sendng this message to inform ..",
						"eventType": "MessageCreate",
						"createdAt": %s,
						}`
			response := fmt.Sprintf(string(rawResponse), getTimeNow())
			// Writing the mock response for both twitter and mastodon
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		} else {
			t.Errorf("Api path Not Found: %s", r.URL.Path)
			http.Error(w, "Api path Not Found", http.StatusNotFound)
		}

	}))
	defer mockServer.Close()

	// Test when API call is successful
	ch := make(chan APIResponse)
	go pollAPI(mockServer.URL+"/twitter/converations", ch)
	response := <-ch
	if response.Err != nil {
		t.Errorf("pollAPI failed, expected no error but got %v", response.Err)
	}
	if response.URL != mockServer.URL+"/twitter/converations" {
		t.Errorf("pollAPI failed, expected URL %s but got %s", mockServer.URL+"/twitter/converations", response.URL)
	}
	if response.Payload == "" {
		t.Errorf("pollAPI failed, expected non-empty payload but got empty payload")
	}

	// Test when API call fails
	go pollAPI(mockServer.URL+"/invalid/path", ch)
	response = <-ch
	if response.Err == nil {
		t.Errorf("pollAPI failed, expected error but got no error")
	}
	if response.URL != mockServer.URL+"/invalid/path" {
		t.Errorf("pollAPI failed, expected URL %s but got %s", mockServer.URL+"/invalid/path", response.URL)
	}
	if response.Payload != "" {
		t.Errorf("pollAPI failed, expected empty payload but got non-empty payload")
	}
}

func TestGetTimeNow(t *testing.T) {
	// Test that the function returns the current date time in the correct format
	expectedFormat := "2006-01-02T15:04:05-000Z"
	now := time.Now()
	expectedTime := now.Format(expectedFormat)
	actualTime := getTimeNow()
	if actualTime != expectedTime {
		t.Errorf("getTimeNow failed, expected %s but got %s", expectedTime, actualTime)
	}
}

func TestMain(m *testing.M) {
	// Load environment variables from dev.env
	os.Setenv("TARGET_ENV", "DEVELOPMENT")
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal("Error loading dev.env file")
	}

	// Run the tests
	code := m.Run()

	// Load environment variables from prod.env
	os.Setenv("TARGET_ENV", "PRODUCTION")
	err = godotenv.Load("prod.env")
	if err != nil {
		log.Fatal("Error loading prod.env file")
	}

	os.Exit(code)
}
