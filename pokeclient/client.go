package pokeclient

import (
	"fmt"
	"io"
	"net/http"
)

const POKEURL = "https://pokeapi.co/api/v2"

type Client struct {
	// client     *http.Client
	requestUrl string
}

// Init to create client
//
// Specify an endpoint eg. pokemon
func Init(endpoint string) *Client {
	reqUrl := fmt.Sprintf("%v/%v", POKEURL, endpoint)
	client := &Client{
		// client:     &http.Client{},
		requestUrl: reqUrl,
	}

	return client
}

func (client *Client) Get(id string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%v/%v", client.requestUrl, id))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
