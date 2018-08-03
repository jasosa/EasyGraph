package easygraph

// Query represents a graphql query
type Query interface {
	GetString() string
}

type rawQuery struct {
	stringQuery string
	variables   []Variable
}

// GetString gets a string representation of the query
func (r *rawQuery) GetString() string {
	return formatRawQuery(r)
}
