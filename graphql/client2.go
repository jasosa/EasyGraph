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
	Object(name string) Query
	StringField(name string) (Query, error)
	ObjectField(objectname string) (Query, error)
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
		objects: make(map[string]*Object),
	}
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

func (q *query) StringField(name string) (Query, error) {

	if q.currentObject == nil {
		return q, ErrFieldsAddedWithoutObject
	}

	q.currentObject.fields = append(q.currentObject.fields, &stringField{name: name})
	return q, nil
}

func (q *query) ObjectField(objectname string) (Query, error) {
	if q.currentObject == nil {
		return q, ErrFieldsAddedWithoutObject
	}

	objectField := &objectField{object: &Object{name: objectname}}
	q.currentObject.fields = append(q.currentObject.fields, objectField)
	q.currentObject = objectField.object
	return q, nil
}

func (q *query) print() string {
	objectsAndFields := []string{}

	objectsAndFields = append(objectsAndFields, `{`)
	for _, object := range q.objects {
		objectsAndFields = append(objectsAndFields, print(object))
	}

	objectsAndFields = append(objectsAndFields, `}`)
	return strings.Join(objectsAndFields, "/")

}
