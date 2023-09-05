package pokeclient

import (
	"fmt"
	"io"
	"net/http"
)

const (
	scheme   = "https"
	host     = "pokeapi.co"
	rootPath = "/api/v2"
	retries  = 3
)

type Client struct {
	requestUrl string
}

// Init to create client
//
// Example:
//
//	client := pokeclint.Init("pokemon")
func Init(endpoint string) *Client {
	POKEURL := fmt.Sprintf("%v://%v/%v", scheme, host, rootPath)
	reqUrl := fmt.Sprintf("%v/%v", POKEURL, endpoint)
	client := &Client{
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
