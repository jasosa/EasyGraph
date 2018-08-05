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

		io.WriteString(w, expectedSimpleResponse)
	}))
	defer s.Close()

	c := NewClient(s.URL)
	q := c.QueryBuilder().Query(simpleQuery)
	res, err := c.Run(q)
	if err != nil {
		t.Fatalf("error was not expected but got %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("error  %q unmarshaling response", err.Error())
	}

	if string(body) != expectedSimpleResponse {
		t.Errorf("%q was expected but got %q", expectedSimpleResponse, string(body))
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

		io.WriteString(w, expectedQueryWithVariablesResponse)
	}))
	defer s.Close()

	c := NewClient(s.URL)
	q := c.QueryBuilder().Query(queryWithVariables)
	q.AddVariable("episode", "JEDI")

	res, err := c.Run(q)
	if err != nil {
		t.Fatalf("error was not expected but got %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("error  %q unmarshaling response", err.Error())
	}

	if string(body) != expectedQueryWithVariablesResponse {
		t.Errorf("%q was expected but got %q", expectedQueryWithVariablesResponse, string(body))
	}
}

func TestQueryWithCredentials(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("authorization")
		if authToken != expectedAuthentication {
			t.Errorf("authentication %q was expected but got %q", expectedAuthentication, authToken)
			return
		}
	}))
	defer s.Close()

	c := NewClient(s.URL)
	c.SetToken("12345678")
	_, err := c.Run(c.QueryBuilder().Query(simpleQuery))
	if err != nil {
		t.Fatalf("error  %q running query", err.Error())
	}
}

var (
	simpleQuery                 = `query { hero { name } }`
	expectedSimpleQueryInServer = `{"query": "query { hero { name } }"}`
	expectedSimpleResponse      = "{ hero { name : \"R2D2\" } }"

	queryWithVariables                 = `query { hero(episode: $episode) { name } }`
	expectedQueryWithVariablesInServer = `{"query": "query { hero(episode: $episode) { name } }","variables": {"episode":"JEDI"}}`
	expectedQueryWithVariablesResponse = "{ hero { name : \"R2D2\" } }"

	expectedAuthentication = `Bearer 12345678`
)
