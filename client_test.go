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
			w.Write([]byte("Error reading query"))
			return
		}

		sbody := string(body)
		if sbody != expectedSimpleQueryInServer {
			t.Errorf("Wrong query sent: %q was expected to be sent but got %q", expectedSimpleQueryInServer, sbody)
			return
		}

		io.WriteString(w, "{ hero { name : \"R2D2\" } }")
	}))
	defer s.Close()

	c := NewClient(s.URL)
	q := c.QueryBuilder().Query(simpleQuery)
	res, err := c.Run(q)
	if err != nil {
		t.Fatalf("Error was not expected but got %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error  %q unmarshaling response", err.Error())
	}

	if string(body) != "{ hero { name : \"R2D2\" } }" {
		t.Errorf("%q was expected but got %q", "{ hero { name : \"R2D2\" } }", string(body))
	}
}

var (
	simpleQuery                 = `query { hero { name } }`
	expectedSimpleQueryInServer = `{"query": "query { hero { name } }"}`
)
