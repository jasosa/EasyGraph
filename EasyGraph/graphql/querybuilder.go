package graphql

import "sync"

//Variable respresents a query variable
type Variable struct {
	Name  string
	Value interface{}
}

// QueryBuilder is used to create Query Objects
type QueryBuilder struct {
	mux           sync.Mutex
	objects       map[string]*Object
	currentObject *Object
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

// AddSingleField adds a single field with no arguments to the current object in the query
func (q *QueryBuilder) AddSingleField(name string) *QueryBuilder {

	if q.currentObject == nil {
		return q
	}

	q.currentObject.fields = append(q.currentObject.fields, &singleField{name: name})
	return q
}

// AddSingleFieldWithArguments adds a single field with arguments to the current object in the query
func (q *QueryBuilder) AddSingleFieldWithArguments(name string, args ...Argument) *QueryBuilder {

	if q.currentObject == nil {
		return q
	}

	q.currentObject.fields = append(q.currentObject.fields, &singleField{name: name, args: args})
	return q
}

// AddObjectField adds a field of type object to the current object in the query
func (q *QueryBuilder) AddObjectField(objectname string) *QueryBuilder {
	if q.currentObject == nil {
		return q
	}

	objectField := &objectField{object: &Object{name: objectname}}
	q.currentObject.fields = append(q.currentObject.fields, objectField)
	q.currentObject = objectField.object
	return q
}

// Query returns a Query
func (q *QueryBuilder) Query() Query {
	query := &StructuredQuery{
		objects: make(map[string]*Object),
	}

	for k, v := range q.objects {
		query.objects[k] = v
	}

	return query
}

//CreateRawQuery creates and returns a raw query
func (q *QueryBuilder) CreateRawQuery(query string, variables ...Variable) Query {
	return &rawQuery{
		stringQuery: query,
		variables:   variables,
	}
}
