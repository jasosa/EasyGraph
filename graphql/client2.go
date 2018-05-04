package graphql

import (
	"strings"
	"sync"
)

// Client2 is a graphql client
type Client2 interface {
	SetToken(token string)
	Query() Query
}

// Query represents a query operation
type Query interface {
	Fields(object string, fields ...string) (Query, error)
	print() string
	//Execute() (string, error)
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
		objects: make(map[string][]field),
	}
}

type query struct {
	mux     sync.Mutex
	objects map[string][]field
}

type field struct {
	name string
}

func (q *query) Fields(object string, fields ...string) (Query, error) {
	o, found := q.objects[object]
	if !found {
		q.objects[object] = []field{}
		o = q.objects[object]
	}

	for _, f := range fields {
		o = append(o, field{name: f})
	}

	q.objects[object] = o
	return q, nil
}

func (q *query) print() string {
	objectsAndFields := []string{}

	objectsAndFields = append(objectsAndFields, `{`)
	for oname, fields := range q.objects {
		objectsAndFields = append(objectsAndFields, oname)
		objectsAndFields = append(objectsAndFields, `{`)
		for _, field := range fields {
			objectsAndFields = append(objectsAndFields, field.name)
		}
		objectsAndFields = append(objectsAndFields, `}`)
	}

	objectsAndFields = append(objectsAndFields, `}`)
	return strings.Join(objectsAndFields, "")

}
