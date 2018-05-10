package graphql

import (
	"bytes"
	"net/http"
)

// Client2 is a graphql client
type Client2 interface {
	SetToken(token string)
	QueryBuilder() *QueryBuilder
	Execute(q Query) (*http.Response, error)
}

type client2 struct {
	url   string
	token string
}

// NewClient2 ...
func NewClient2(url string) Client2 {
	return &client2{
		url: url,
	}
}

func (c *client2) SetToken(token string) {
	c.token = token
}

func (c *client2) QueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		objects: make(map[string]*Object),
	}
}

func (c *client2) Execute(q Query) (*http.Response, error) {
	query := q.GetString()
	return doReq(c.url, c.token, query)
}

func doRequest(url, token, formattedQuery string) (*http.Response, error) {
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
