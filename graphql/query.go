package graphql

// Query represents a graphql query
type Query interface {
	GetString() string
}

// StructuredQuery represents a graphql query
type StructuredQuery struct {
	objects map[string]*Object
}

// getObjects get a map with the current objects in the query
func (q *StructuredQuery) getObjects() map[string]*Object {
	return q.objects
}

// GetString gets a string representation of the query
func (q *StructuredQuery) GetString() string {
	return formatStructuredQuery(q)
}

type rawQuery struct {
	stringQuery string
	variables   []Variable
}

// GetString gets a string representation of the query
func (r *rawQuery) GetString() string {
	return formatRawQuery(r)
}
