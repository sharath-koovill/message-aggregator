package main

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

const (
	ConsumerKey       = "Need to be read from the google secrets"
	ConsumerSecret    = "Need to be read from the google secrets"
	AccessToken       = "Need to be read from the google secrets"
	AccessTokenSecret = "Need to be read from the google secrets"
)

func main() {
	// Authenticate with Twitter API
	anaconda.SetConsumerKey(ConsumerKey)
	anaconda.SetConsumerSecret(ConsumerSecret)
	api := anaconda.NewTwitterApi(AccessToken, AccessTokenSecret)

	// Get the direct messages
	messages, err := api.GetDirectMessages(nil)
	if err != nil {
		fmt.Println("Error fetching direct messages:", err)
		return
	}

	// Print the messages
	for _, message := range messages {
		fmt.Println("Sender:", message.Sender.ScreenName)
		fmt.Println("Recipient:", message.Recipient.ScreenName)
		fmt.Println("Message:", message.Text)
		fmt.Println("-----")
	}
}
