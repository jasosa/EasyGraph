package easygraph

import "sync"

//Variable respresents a query variable
type Variable struct {
	Name  string
	Value interface{}
}

// QueryBuilder is used to create Query Objects
type QueryBuilder struct {
	mux sync.Mutex
}

//Query creates and returns a raw query
func (q *QueryBuilder) Query(query string, variables ...Variable) Query {
	return &rawQuery{
		stringQuery: query,
		variables:   variables,
	}
}
