package mastodon

import (
	"fmt"

	"github.com/mattn/go-mastodon"
)

func getDirectMessages(instanceURL string, accessToken string) ([]mastodon.Status, error) {
	/*
	   This function takes in an instance URL and an access token for a Mastodon account,
	   and returns a slice of direct messages (statuses) for that account.

	   Parameters:
	   instanceURL (string): The URL of the Mastodon instance
	   accessToken (string): The access token for the Mastodon account

	   Returns:
	   []mastodon.Status: A slice of direct messages (statuses) for the Mastodon account
	   error: Any errors encountered while fetching the direct messages
	*/
	client := mastodon.NewClient(&mastodon.Config{
		Server:      instanceURL,
		AccessToken: accessToken,
	})

	statuses, err := client.GetDirectMessages()
	if err != nil {
		// Log the error
		fmt.Printf("Error fetching direct messages: %v", err)
		return nil, err
	}

	return statuses, nil
}
