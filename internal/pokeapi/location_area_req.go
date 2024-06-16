package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const baseURL = "https://pokeapi.co/api/v2"

func (c *Client) LocationAreaRequest(URL string) (LocationApiResponse, error) {
	endPoint := "/location-area"
	fullURL := baseURL + endPoint
	if URL != "" {
		fullURL = URL
	}
	res, err := c.client.Get(fullURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	locApiResp := &LocationApiResponse{}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationApiResponse{}, err
	}
	err = json.Unmarshal(data, locApiResp)
	if err != nil {
		return LocationApiResponse{}, err
	}
	return *locApiResp, nil
}
