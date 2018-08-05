package easygraph

import (
	"bytes"
	"net/http"
)

// Client is a graphql client
type Client interface {
	SetToken(token string)
	QueryBuilder() *QueryBuilder
	Run(q Query) (*http.Response, error)
}

type client struct {
	url   string
	token string
}

// NewClient initializes a new client
func NewClient(url string) Client {
	return &client{
		url: url,
	}
}

func (c *client) SetToken(token string) {
	c.token = token
}

func (c *client) QueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (c *client) Run(q Query) (*http.Response, error) {
	query := q.GetString()
	return doReq(c.url, c.token, query)
}

func doReq(url, token, formattedQuery string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(formattedQuery))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if token != "" {
		var bearer = "Bearer " + token
		req.Header.Add("authorization", bearer)
	}

	return http.DefaultClient.Do(req)
}
