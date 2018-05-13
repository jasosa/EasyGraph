package easygraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueryReturnsAnEmptyQuery(t *testing.T) {
	c := NewClient("myurl")
	q := c.QueryBuilder()
	if q == nil {
		t.Errorf("Query object was expected but got nil")
	}
}

func TestAddSingleFieldWithoutArguments(t *testing.T) {

	c := NewClient("myurl")
	q := c.QueryBuilder().AddObject("hero").AddSingleField("name")
	printedQuery := q.Query().GetString()
	if expectedSingleFieldWithoutArguments != printedQuery {
		t.Errorf("\n `%s` \n was expected but got `%s`", expectedSingleFieldWithoutArguments, printedQuery)
	}
}

func TestAddNestedObjectField(t *testing.T) {
	c := NewClient("myurl")
	q := c.QueryBuilder().AddObject("hero").AddSingleField("name").AddObjectField("friends").AddSingleField("name")
	printedQuery := q.Query().GetString()
	if expectedNestedObjectField != printedQuery {
		t.Errorf("\n `%s` \n was expected but got `%s`", expectedNestedObjectField, printedQuery)
	}
}
func TestAddSimpleFieldWithArguments(t *testing.T) {
	c := NewClient("myurl")
	arg := Argument{
		Name:  "id",
		Value: "1000",
	}
	q := c.QueryBuilder().AddObject("heroes").AddSingleFieldWithArguments("human", arg)
	printedQuery := q.Query().GetString()
	if expectedSingleWithArguments != printedQuery {
		t.Errorf("\n `%s` \n was expected but got `%s`", expectedSingleWithArguments, printedQuery)
	}
}

func TestExecuteSimpleQuerySuccesfully(t *testing.T) {
	handler := &testHandler{}
	ts := httptest.NewServer(handler)

	c := NewClient(ts.URL)
	q := c.QueryBuilder()
	q = q.AddObject("viewer").AddSingleField("login")
	res, err := c.Execute(q.Query())
	if err != nil {
		t.Errorf("Error was not expected but got %s", err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	if string(body) != string(loginCallAnswer) {
		t.Errorf("%v was expected but got %v", string(loginCallAnswer), string(body))
	}
}

func TestExecuteQueryWithArgumentsSuccesfully(t *testing.T) {
	handler := &testHandler{}
	ts := httptest.NewServer(handler)

	c := NewClient(ts.URL)
	qb := c.QueryBuilder()
	arg := Argument{
		Name:  "size",
		Value: 512,
	}

	query := qb.AddObject("viewer").AddSingleFieldWithArguments("avatarUrl", arg).Query()
	res, err := c.Execute(query)
	if err != nil {
		t.Errorf("Error was not expected but got %s", err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	if string(body) != string(avatarCallAnswer) {
		t.Errorf("%v was expected but got %v", string(avatarCallAnswer), string(body))
	}
}

func TestExecuteRawQuery(t *testing.T) {
	handler := &testHandler{}
	ts := httptest.NewServer(handler)

	c := NewClient(ts.URL)
	qb := c.QueryBuilder()
	_ = qb
	//qb.CreateRawQuery(``)
}

var expectedSingleFieldWithoutArguments = `{"query": "query { hero { name }}"}`
var expectedNestedObjectField = `{"query": "query { hero { name friends { name } }}"}`
var expectedSingleWithArguments = `{"query": "query { heroes { human (id: \"1000\" ) }}"}`

var loginCallAnswer = []byte(`{
	"data": {
	  "viewer": {
		"login": "jasosa"
	  }
	}
  }`)

var avatarCallAnswer = []byte(`{
	"data": {
	  "viewer": {
		"avatarUrl": "myavataurl"
	  }
	}
  }`)

var (
	loginCall  = "query { viewer { login }}"
	avatarCall = "query { viewer { avatarUrl (size: 512 ) }}"
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
	case avatarCall:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(avatarCallAnswer)
	}

}
