package shodan

const baseUrl = "http://api.shodan.io"
// base url to be used easily to access the api

type Client struct{
	apiKey string
}
// type for all clients accessing all methods go throught that struct

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}