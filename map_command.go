package main

import (
	"errors"
	"fmt"
	"os"
)

func cmdMap(conf *Config) error {
	locRes, err := conf.client.LocationAreaRequest(conf.next)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	conf.next = locRes.Next
	conf.previous = locRes.Previous
	for _, location := range locRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func cmdMapb(conf *Config) error {
	if conf.previous == "" {
		return errors.New("you're on the first page")
	}

	locRes, err := conf.client.LocationAreaRequest(conf.previous)

	if err != nil {
		return err
	}
	conf.next = locRes.Next
	conf.previous = locRes.Previous
	for _, location := range locRes.Results {
		fmt.Println(location.Name)
	}
	return nil
}
