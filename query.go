package easygraph

// Query represents a graphql query
type Query interface {
	AddVariable(name string, value interface{})
	GetString() string
}

type rawQuery struct {
	stringQuery string
	variables   []variable
}

func (r *rawQuery) GetString() string {
	return formatRawQuery(r)
}

func (r *rawQuery) AddVariable(name string, value interface{}) {
	r.variables = append(r.variables, variable{Name: name, Value: value})
}
