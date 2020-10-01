package shodan

import (
	"fmt"
	"net/http"
)

type HostLocation struct{
	City         string `json:"city"`
	RegionCode   string `json:"region_code"`
	AreaCode     string `json:"area_code"`
	Longitude    string `json:"longitude"`
	CountryCode3 string `json:"country_code3"`
	CountryName  string `json:"country_name"`
	PostalCode   string `json:"postal_code"`
	DMACode      string `json:"dma_code"`
	CountryCode  string  `json:"country_code"`
	Latitude     float32 `json:"latitude"`
}
// struct to store response of the location

type Host struct{
	OS        string       `json:"os"`
	Timestamp string       `json:"timestamp"`
	ISP       string       `json:"isp"`
	ASN       string       `json:"asn"`
	Hostnames []string     `json:"hostnames"`
	Location  HostLocation `json:"location"`
	IP        int64        `json:"ip"`
	Domains   []string     `json:"domains"`
	Org       string       `json:"org"`
	Data      string       `json:"data"`
	Port      int          `json:"port"`
	IPString  string       `json:"ip_str"`
}
// struct to store response of the host

type SingleCount struct{
	count int    `json:"count"`
	value string `json:"value"`
}
// stores a count of a single response

type HostSearch struct{
	Matches []Host `json:"matches"`
}
// stores responses from search

type CountSearch struct{
	counts []SingleCount `json:"org"` 
}
// stores response from counting the 


func (c *Client) HostSearch(query string) (*HostSearch, error) {

	response, err := http.Get(fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", baseUrl, c.apiKey, query))
	defer resposne.Body.Close()
	// gets response from the search

	if err != nil {
		return nil, err
		// just returns nothing if there's no resposne
	}

	var hostSearch HostSearch
	// init var to store decoded response

	err2 := json.NewDecoder(response.Body).Decode(&hostSearch)
	// decodes response and stores in hostSearch

	if err2 != nil {
		return nil, err2
	}
	return &hostSearch, nil
}

func (c *Client) CountSearch(query string) {

	response, err := http.Get(fmt.Sprintf("%s/shodan/host/count?key=%s&query=%s", baseUrl, c.apiKey, query))
	defer response.Body.Close()
	// gets the total number of results from a search

	if err != nil {
		return nil, err
		// if it can't be gotten return nothing
	}

	var countSearch CountSearch
	// init var to store responses from counts

	err2 := json.NewDecoder(response.Body).Decode(&countSearch)
	// decodes response

	if err2 != nil {
		return nil, err2
		// returns nothing if it can't be decoded
	}
	return &countSearch, nil
}