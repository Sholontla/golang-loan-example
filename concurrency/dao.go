package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func ChannelServiceDao(Ch <-chan Clients) {
	start := time.Now()

	// read values from channel
	for v := range Ch {
		file, _ := json.MarshalIndent(v, "", " ")
		os.WriteFile("clientChannTest.json", file, 0777) //0644
		log.Print(v)
	}
	elapsed := time.Since(start)
	log.Printf("Channel took %s", elapsed)
	wg.Done()

}

func NormalServiceDao(client Clients) {
	start := time.Now()

	// read values from channel

	file, _ := json.MarshalIndent(client, "", " ")
	os.WriteFile("clientNormalTest.json", file, 0777) //0644
	log.Print(client)

	elapsed := time.Since(start)
	log.Printf("Normal took %s", elapsed)
}
