package graphql

import (
	"bytes"
	"net/http"
	"sync"
)

// Client2 is a graphql client
type Client2 interface {
	SetToken(token string)
	Query() Query
	Execute(q Query) (*http.Response, error)
}

// Query represents a query operation
type Query interface {
	Object(name string) Query
	StringField(name string) Query
	ObjectField(objectname string) Query
	print() string
	getObjects() map[string]*Object
}

type client2 struct {
	url   string
	token string
}

// NewClient2 ...
func NewClient2(url string) Client2 {
	return &client2{
		url: url,
	}
}

func (c *client2) SetToken(token string) {
	c.token = token
}

func (c *client2) Query() Query {
	return &query{
		objects: make(map[string]*Object),
	}
}

func (c *client2) Execute(q Query) (*http.Response, error) {
	query := formatQuery(q)
	return doReq(c.url, c.token, query)
}

type query struct {
	mux           sync.Mutex
	objects       map[string]*Object
	currentObject *Object
}

func (q *query) Object(name string) Query {
	_, found := q.objects[name]
	if !found {
		q.objects[name] = &Object{name: name}
	}

	q.currentObject = q.objects[name]
	return q
}

func (q *query) StringField(name string) Query {

	if q.currentObject == nil {
		return q
	}

	q.currentObject.fields = append(q.currentObject.fields, &stringField{name: name})
	return q
}

func (q *query) ObjectField(objectname string) Query {
	if q.currentObject == nil {
		return q
	}

	objectField := &objectField{object: &Object{name: objectname}}
	q.currentObject.fields = append(q.currentObject.fields, objectField)
	q.currentObject = objectField.object
	return q
}

func (q *query) Execute() (string, error) {
	query := formatQuery(q)
	return query, nil
}

func (q *query) print() string {
	return formatQuery(q)
}

func (q *query) getObjects() map[string]*Object {
	return q.objects
}

func doRequest(url, token, formattedQuery string) (*http.Response, error) {
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
