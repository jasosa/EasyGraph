package graphql

import "sync"

// QueryBuilder represents a graphql query
type QueryBuilder struct {
	mux           sync.Mutex
	objects       map[string]*Object
	currentObject *Object
}

type Query struct {
	objects map[string]*Object
}

// AddObject adds  an object (or select an existing one) to be queried to the query with the given name
// and marks the object as the current one
func (q *QueryBuilder) AddObject(name string) *QueryBuilder {
	_, found := q.objects[name]
	if !found {
		q.objects[name] = &Object{name: name}
	}

	q.currentObject = q.objects[name]
	return q
}

// AddStringField adds a field of type string to the current object in the query
func (q *QueryBuilder) AddStringField(name string) *QueryBuilder {

	if q.currentObject == nil {
		return q
	}

	q.currentObject.fields = append(q.currentObject.fields, &stringField{name: name})
	return q
}

// ObjectField adds a field of type object to the current object in the query
func (q *QueryBuilder) ObjectField(objectname string) *QueryBuilder {
	if q.currentObject == nil {
		return q
	}

	objectField := &objectField{object: &Object{name: objectname}}
	q.currentObject.fields = append(q.currentObject.fields, objectField)
	q.currentObject = objectField.object
	return q
}

// Query returns a Query
func (q *QueryBuilder) Query() *Query {
	query := &Query{
		objects: make(map[string]*Object),
	}

	for k, v := range q.objects {
		query.objects[k] = v
	}

	return query
}

// print prints a representation of the query
func (q *QueryBuilder) print() string {
	return formatQuery(q.Query())
}

// getObjects get a map with the current objects in the query
func (q *Query) getObjects() map[string]*Object {
	return q.objects
}
