package graphql

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueryReturnsAnEmptyQuery(t *testing.T) {
	c := NewClient2("myurl")
	q := c.Query()
	if q == nil {
		t.Errorf("Query object was expected but got nil")
	}
}

func TestAddObjectAndFieldsToQuerySuccesfully(t *testing.T) {

	c := NewClient2("myurl")
	q := c.Query().Object("hero").StringField("name")
	printedQuery := q.print()
	if expectedPrintedHeroRaw != printedQuery {
		t.Errorf("\n `%s` \n was expected but got `%s`", expectedPrintedHeroRaw, printedQuery)
	}
}

func TestAddObjectAsAFieldSuccesfully(t *testing.T) {
	c := NewClient2("myurl")
	q := c.Query().Object("hero").StringField("name").ObjectField("friends").StringField("name")
	printedQuery := q.print()
	if expectedPrintedHeroWithFriendsRaw != printedQuery {
		t.Errorf("\n `%s` \n was expected but got `%s`", expectedPrintedHeroWithFriendsRaw, printedQuery)
	}
}

func TestExecuteSimpleQuerySuccesfully(t *testing.T) {
	handler := &testHandler{}
	ts := httptest.NewServer(handler)

	c := NewClient2(ts.URL)
	q := c.Query()
	q = q.Object("viewer").StringField("login")
	res, err := c.Execute(q)
	if err != nil {
		t.Errorf("Error was not expected but got %s", err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	if string(body) != string(loginCallAnswer) {
		t.Errorf("%v was expected but got %v", string(loginCallAnswer), string(body))
	}

}

var expectedPrintedHeroRaw = `{"query": "query { hero { name }}"}`
var expectedPrintedHeroWithFriendsRaw = `{"query": "query { hero { name friends { name } }}"}`

var loginCallAnswer = []byte(`{
	"data": {
	  "viewer": {
		"login": "jasosa"
	  }
	}
  }`)

var (
	loginCall = "query { viewer { login }}"
)

type graphqlQuery struct {
	Query     string            `json:"query"`
	Variables map[string]string `json:"variables"`
}

type testHandler struct {
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	gq := &graphqlQuery{}
	err := json.Unmarshal(body, gq)
	if err != nil {
		fmt.Println(err)
	}

	switch gq.Query {
	case loginCall:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(loginCallAnswer)
	}
}
