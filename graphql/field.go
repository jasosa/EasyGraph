package graphql

// Field represents every type of field in a query
type Field interface {
	GetString() string
}

type stringField struct {
	name string
}

func (s *stringField) GetString() string {
	return s.name
}

type objectField struct {
	object *Object
}

func (o *objectField) GetString() string {
	return print(o.object)
}
