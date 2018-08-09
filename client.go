package easygraph

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Client is a graphql client
type Client interface {
	SetToken(token string)
	QueryBuilder() *QueryBuilder
	Run(q Query, response interface{}) error
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

func (c *client) Run2(q Query) (*http.Response, error) {
	_ = q.GetString()
	//return doReq(c.url, c.token, query)
	return nil, nil
}

func (c *client) Run(q Query, response interface{}) error {
	query := q.GetString()
	return doReq(c.url, c.token, query, response)
}

func doReq(url, token, formattedQuery string, response interface{}) error {
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(formattedQuery))
	if err != nil {
		return errors.Wrap(err, "error creating request")
	}

	req.Header.Set("Content-Type", "application/json")

	if token != "" {
		var bearer = "Bearer " + token
		req.Header.Add("authorization", bearer)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "error executing request")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "error reading response")
	}

	if err := json.Unmarshal(body, response); err != nil {
		return errors.Wrap(err, "error decoding response")
	}

	return nil
}
