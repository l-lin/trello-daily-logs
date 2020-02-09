package trello

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: time.Second * 10}

const trelloURL = "https://api.trello.com/1/lists/%s/cards"

// GetCards from trello
func GetCards(listID, key, token string) ([]Card, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(trelloURL, listID), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", key)
	q.Add("token", token)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Could not fetch the cards. Error status was %d", resp.StatusCode)
	}
	var cards []Card
	if err = json.Unmarshal(body, &cards); err != nil {
		return nil, err
	}
	return cards, nil
}
