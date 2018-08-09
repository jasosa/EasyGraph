package easygraph

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClientInitializesAClient(t *testing.T) {
	c := NewClient("myurl")
	if c == nil {
		t.Errorf("An initialized client was expected but got nil")
	}
}

func TestSendSimpleQuery(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("error reading query"))
			return
		}

		sbody := string(body)
		if sbody != expectedSimpleQueryInServer {
			t.Errorf("wrong query sent: %q was expected to be sent but got %q", expectedSimpleQueryInServer, sbody)
			return
		}

		io.WriteString(w, serverSimpleQueryResponse)
	}))
	defer s.Close()

	c := NewClient(s.URL)
	q := c.QueryBuilder().Query(simpleQuery)
	var resp QueryResponse
	err := c.Run(q, &resp)
	if err != nil {
		t.Fatalf("error was not expected but got %v", err)
	}

	if resp.Data.Hero.Name != expectedQueryResponse {
		t.Errorf("%q was expected but got %q", expectedQueryResponse, resp.Data.Hero.Name)
	}
}

func TestSendWithVariables(t *testing.T) {

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("error reading query"))
			return
		}

		sbody := string(body)
		if sbody != expectedQueryWithVariablesInServer {
			t.Errorf("wrong query sent: %q was expected to be sent but got %q", expectedQueryWithVariablesInServer, sbody)
			return
		}

		io.WriteString(w, serverQueryWithVariablesResponse)
	}))
	defer s.Close()

	c := NewClient(s.URL)
	q := c.QueryBuilder().Query(queryWithVariables)
	q.AddVariable("episode", "JEDI")

	var resp QueryResponse
	err := c.Run(q, &resp)
	if err != nil {
		t.Fatalf("error was not expected but got %v", err)
	}

	if resp.Data.Hero.Name != expectedQueryResponse {
		t.Errorf("%q was expected but got %q", expectedQueryResponse, resp.Data.Hero.Name)
	}
}

func TestQueryWithCredentials(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("authorization")
		if authToken != expectedAuthentication {
			t.Errorf("authentication %q was expected but got %q", expectedAuthentication, authToken)
			return
		}

		io.WriteString(w, serverSimpleQueryResponse)
	}))
	defer s.Close()

	c := NewClient(s.URL)
	c.SetToken("12345678")
	m := make(map[string]interface{})
	err := c.Run(c.QueryBuilder().Query(simpleQuery), &m)
	if err != nil {
		t.Fatalf("error  %q running query", err.Error())
	}
}

var (
	simpleQuery                 = `query { hero { name } }`
	expectedSimpleQueryInServer = `{"query": "query { hero { name } }"}`
	serverSimpleQueryResponse   = `{"data":{"hero":{"name":"R2-D2"}}}`

	queryWithVariables                 = `query { hero(episode: $episode) { name } }`
	expectedQueryWithVariablesInServer = `{"query": "query { hero(episode: $episode) { name } }","variables": {"episode":"JEDI"}}`
	serverQueryWithVariablesResponse   = `{"data":{"hero":{"name":"R2-D2"}}}`

	expectedQueryResponse = "R2-D2"

	expectedAuthentication = `Bearer 12345678`
)

type QueryResponse struct {
	Data struct {
		Hero struct {
			Name string `json:"name"`
		} `json:"hero"`
	} `json:"data"`
}
