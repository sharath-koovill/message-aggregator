package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"message_service/internal/mock"
	"message_service/internal/pulsar"

	"github.com/joho/godotenv"
)

// LookupEnvOrString looks up an environment variable and returns its value or a default value if not found
func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

// LookupEnvOrInt looks up an environment variable and returns its value as an int or a default value if not found
func LookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("LookupEnvOrInt[%s]: %v", key, err)
		}
		return v
	}
	return defaultVal
}

// init loads environment variables from dev.env or prod.env depending on the value of TARGET_ENV
func init() {
	targetEnv := LookupEnvOrString("TARGET_ENV", "DEVELOPMENT")

	if targetEnv == "DEVELOPMENT" {
		err := godotenv.Load("dev.env")
		if err != nil {
			log.Fatal("Error loading dev.env file")
		}

	} else {
		err := godotenv.Load("prod.env")
		if err != nil {
			log.Fatal("Error loading prod.env file")
		}
	}

}

type APIResponse struct {
	URL     string
	Payload string
	Err     error
}

// pollAPI call the api provided and sends the response to the channel
func pollAPI(url string, ch chan<- APIResponse) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- APIResponse{URL: url, Err: err}
		return
	}
	defer resp.Body.Close()

	// Reading the response body
	payload := make([]byte, 1024)
	n, _ := resp.Body.Read(payload)
	ch <- APIResponse{URL: url, Payload: string(payload[:n])}
}

func main() {
	var (
		pollInterval = flag.Int("poll_interval", LookupEnvOrInt("POLL_INTERVAL", 120), "Polling interval")
		pulsarUrl    = flag.String("pulsar_url", LookupEnvOrString("PULSAR_URL", ""), "Pulsar Url")
		pulsarTopic  = flag.String("pulsar_topic", LookupEnvOrString("PULSAR_TOPIC", ""), "Pulsar Topic")
	)
	flag.Parse()

	mockServer := mock.MockServer()
	pulsarClient := pulsar.GetPulsarClient(pulsarUrl)

	apiURLs := []string{
		mockServer.URL + "/twitter/conversations",
		mockServer.URL + "/mastodon/conversations",
	}

	ch := make(chan APIResponse)

	// Waiting for responses from all APIs
	for {
		for _, url := range apiURLs {
			go pollAPI(url, ch)
		}

		for range apiURLs {
			select {
			case response := <-ch:
				if response.Err != nil {
					// Add sentry alerts here
					log.Fatalf("Error fetching data from %s: %s\n", response.URL, response.Err)
				} else {
					// Sending the real time messages to apache pulsar
					pulsar.ProduceMessage(pulsarClient, pulsarTopic, response.Payload)
					log.Printf("Data from %s: %s\n", response.URL, response.Payload)
				}
			case <-time.After(5 * time.Second):
				log.Printf("Timeout: No response received within 5 seconds.")
			}
		}

		time.Sleep(time.Duration(*pollInterval) * time.Second)
	}
	defer pulsarClient.Close()
}
