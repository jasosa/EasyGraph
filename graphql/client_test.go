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

func TestAddObjectAndFieldsToQuerySuccesfully(t *testing.T) {
	/* handler := &testHandler{}
	ts := httptest.NewServer(handler)
	*/
	c := NewClient2("myurl")
	q := c.Query()
	q, _ = q.Object("hero").StringField("name")
	printedQuery := q.print()
	if expectedPrintedHeroRaw != printedQuery {
		t.Errorf("\n `%s` \n was expected but got `%s`", expectedPrintedHeroRaw, printedQuery)
	}
}

func TestAddFieldsWithNoObjectReturnsError(t *testing.T) {
	c := NewClient2("myurl")
	q := c.Query()
	_, err := q.StringField("name")
	if err != ErrFieldsAddedWithoutObject {
		t.Errorf("Expected error %v but got %v", ErrFieldsAddedWithoutObject, err)
	}
}

func TestAddObjectAsAFieldSuccesfully(t *testing.T) {
	c := NewClient2("myurl")
	q := c.Query()
	q, _ = q.Object("hero").StringField("name")
	q, _ = q.ObjectField("friends")
	q, _ = q.StringField("name")
	printedQuery := q.print()
	if expectedPrintedHeroWithFriendsRaw != printedQuery {
		t.Errorf("\n `%s` \n was expected but got `%s`", expectedPrintedHeroWithFriendsRaw, printedQuery)
	}
}

var expectedPrintedHeroRaw = `{/hero/{/name/}/}`
var expectedPrintedHeroWithFriendsRaw = `{/hero/{/name/friends/{/name/}/}/}`

/* type testHandler struct {
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
} */
