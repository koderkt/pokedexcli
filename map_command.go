package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func cmdMap() error {
	fmt.Println("In Map.......")
	// var requestURL string
	// if config.next == "" {
	requestURL := "https://pokeapi.co/api/v2/location-area"
	// } else {
	// 	requestURL = config.next
	// }

	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	locRes := &locationApiResponse{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, locRes)
	if err != nil {
		return err
	}
	for _, location := range locRes.results {
		fmt.Printf("Location Name: %s, URL: %s\n", location.name, location.url)
	}

	// fmt.Println()
	return nil
}
