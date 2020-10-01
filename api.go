package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/Ieuan-Steel11/simple-shodan-client"
)

type APIInfo struct {
	QueryCredits int    `json:"query_credits"`
	ScanCredits  int    `json:"scan_credits"`
	Telnet       bool   `json:"telnet"`
	Plan         string `json:"plan"`
	Https        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
}
// struct for the response about api info

func (c *Client) getAPIInfo() (*APIInfo, error) {

	response, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", baseUrl, c.apiKey))
	defer response.Body.Close()
	// gets api info

	if err != nil {
		return nil, err
		// returns nothing if you can't get a response
	}

	var apiInfo APIInfo
	// var to store the decoded response

	err2 := json.NewDecoder(response.Body).Decode(&apiInfo)
	// gets decoded version of response from api info

	if err2 != nil {
		return nil, err
		// returns nothing if you can't get a response
	}
	return &apiInfo, nil
}