# Message exchange for retreiving twitter and Mastodon direct messages

This is the message exchange which makes a short poll to both twitter api and mastodon api and retrieves unread messages from twitter and mastodon in real time.

There are multiple approaches to retreive direct messages from twitter and mastodon. 

One approach is to make a long poll request to the account activity API v2 of the twitter API. This is a webhook based API. We can do something similar to mastodon as well.
https://developer.twitter.com/en/docs/tutorials/getting-started-with-the-account-activity-api

Another approach is to use a social media integration platform APIs like pipedream etc there we dont make the heavy lifting of integration with different social media, we make a single integration to such platform and get messages from the social media of our choice.

For demonstration purpose i am intercepting the api calls using a mock server and returning a mock response which is similar to the twitter direct conversation api response. 

The exchange service polls the api every 2 minutes and that is equivalent to 30 requests per hour * 24 (hours) * 30 (days in a month) = ~21600 api calls per month.

## Prerequisites for running this service
Make sure the apache pulsar is running at localhost:6650

## Command to run
### Build the message_exchange docker container
`docker build -t message_exchange .`

## Run message_exchange
`docker run message_exchange`

## Unit test
go test

