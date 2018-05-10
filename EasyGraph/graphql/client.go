package graphql

import (
	"bytes"
	"net/http"
	"strconv"
)

type client struct {
	url   string
	token string
}

//Client is a graphql client
type Client interface {
	SetToken(token string)
	DoQuery(query string) (*http.Response, error)
	DoQueryWithVariables(query string, variables ...Variable) (*http.Response, error)
}

// NewClient initialize a new graphql client
func NewClient(url string) Client {
	return &client{url: url}
}

// SetToken sets the token to the client
func (c *client) SetToken(token string) {
	c.token = token
}

// DoQuery performs a graphql query
func (c *client) DoQuery(query string) (*http.Response, error) {
	formattedQuery := formatQueryToRequest(query)
	return doRequest(c.url, c.token, formattedQuery)
}

// DoQueryWithVariables performs a graphql query that includes one or more variables
func (c *client) DoQueryWithVariables(query string, variables ...Variable) (*http.Response, error) {
	variablesQuery := formatVariables(variables)
	formattedQuery := formatQueryWithvariables(
		strconv.QuoteToASCII(query),
		variablesQuery)

	return doRequest(c.url, c.token, formattedQuery)
}

func formatQueryToRequest(queryString string) string {
	quotedQuery := strconv.QuoteToASCII(queryString)
	return `{"query": ` + quotedQuery + `}`
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
