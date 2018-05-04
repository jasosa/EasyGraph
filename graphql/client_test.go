package graphql

import (
	"testing"
)

func TestQueryReturnsAnEmptyQuery(t *testing.T) {
	c := NewClient2("myurl")
	q := c.Query()
	if q == nil {
		t.Errorf("Query object was expected but got nil")
	}
}

func TestAddObjectToQuerySuccesfully(t *testing.T) {
	/* handler := &testHandler{}
	ts := httptest.NewServer(handler)
	*/
	c := NewClient2("myurl")
	q := c.Query()
	q, _ = q.Fields("hero", "name")
	printedQuery := q.print()
	if heroQueryRaw != printedQuery {
		t.Errorf("\n `%s` \n was expected but got `%s`", heroQueryRaw, printedQuery)
	}
}

var heroQueryRaw = `{hero{name}}`

/* type testHandler struct {
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
} */
