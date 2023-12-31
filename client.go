package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/patrickmn/go-cache"
)

type RESTClient struct {
	client     *http.Client
	requestUrl *url.URL
	cache      *cache.Cache
}

func init() {
	newCache := cache.New(defaultCacheSettings.MinExpire, defaultCacheSettings.MaxExpire)
	papiurl, err := buildClientEndpoint(papiEnv)
	if err != nil {
		panic(err)
	}
	client = &RESTClient{
		requestUrl: papiurl,
		cache:      newCache,
	}
}

func (c *RESTClient) Get(pokeapiUrl string, pokeObj any) error {
	item, found := client.cache.Get(pokeapiUrl)
	if found && CacheSettings.UseCache {
		if err := json.Unmarshal(item.([]byte), &pokeObj); err != nil {
			return err
		}
		return nil
	}

	resp, err := http.Get(pokeapiUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &pokeObj); err != nil {
		return err
	}
	return nil

}

// Creates a URL for the given endpoint using the current configuration
func buildClientEndpoint(env string) (*url.URL, error) {
	papiurl, err := url.Parse(os.Getenv(env))
	if err != nil {
		return nil, fmt.Errorf("could not parse env: %v", err)
	}
	if papiEnvEmpty(papiurl.String()) {
		return &papiDefault, nil
	}
	if papiurl.Host == "" {
		return nil, fmt.Errorf("unable to parse url host from: %s\n", papiurl)
	}

	return papiurl, nil
}

func (c *RESTClient) urlBuilder(endpoint, id string) string {
	return fmt.Sprintf("%v/%v/%v", client.requestUrl.String(), endpoints[endpoint], id)
}

func papiEnvEmpty(env string) bool {
	if env == "" {
		return true
	}
	return false
}
