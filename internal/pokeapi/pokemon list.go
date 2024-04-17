package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(name string) (RespExplore, error) {
	url := baseURL + "/location-area/" + name + "/"

	if val, ok := c.cache.Get(url); ok {
		exploreResp := RespExplore{}
		err := json.Unmarshal(val, &exploreResp)
		if err != nil {
			return RespExplore{}, err
		}
		return exploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return RespExplore{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return RespExplore{}, err
	}

	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExplore{}, err
	}

	exploreResp := RespExplore{}
	err = json.Unmarshal(dat, &exploreResp)
	if err != nil {
		return RespExplore{}, err
	}

	c.cache.Add(url, dat)
	return exploreResp, nil
}
