package graphql

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

//Variable respresents a query variable
type Variable struct {
	Name  string
	Value interface{}
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
	variablesQuery := variablesToString(variables)
	formattedQuery := formatQueryWithvariablesToRequest(
		strconv.QuoteToASCII(query),
		variablesQuery)

	return doRequest(c.url, c.token, formattedQuery)
}

func variablesToString(variables []Variable) string {
	queryVariables := []string{}
	var keyvalue string
	for _, v := range variables {
		s, ok := v.Value.(string)
		if ok {
			keyvalue = fmt.Sprintf(`%s:%v`, strconv.QuoteToASCII(v.Name), strconv.QuoteToASCII(s))
		} else {
			keyvalue = fmt.Sprintf(`%s:%v`, strconv.QuoteToASCII(v.Name), v.Value)
		}
		queryVariables = append(queryVariables, keyvalue)
	}
	return strings.Join(queryVariables, ",")
}

func formatQueryWithvariablesToRequest(queryString string, variablesString string) string {
	query := `{"query": ` + queryString + `,` +
		`"variables": {` + variablesString + `}}`
	return query
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
