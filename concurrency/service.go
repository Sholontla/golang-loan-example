package main

import (
	"sync"

	"github.com/google/uuid"
	"github.com/projects/loans/utils/date_utils"
)

var wg sync.WaitGroup

func ChannelRequestService(client Clients) (*Clients, error) {
	Ch := make(chan Clients)

	errVal := client.UserValidation()
	if errVal != nil {
		return nil, errVal
	}
	client.ClientId = uuid.New()
	client.ClientDateCreated = date_utils.GetNowString()

	go func() {
		Ch <- client
		close(Ch)
	}()

	wg.Add(1)
	go ChannelServiceDao(Ch)
	wg.Wait()

	return &client, nil
}

func NormalRequestService(client Clients) (*Clients, error) {

	errVal := client.UserValidation()
	if errVal != nil {
		return nil, errVal
	}
	client.ClientId = uuid.New()
	client.ClientDateCreated = date_utils.GetNowString()

	NormalServiceDao(client)

	return &client, nil
}
