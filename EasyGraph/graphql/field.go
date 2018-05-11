package graphql

import (
	"fmt"
	"strconv"
)

// Field represents every type of field in a query
type Field interface {
	GetString() string
}

//Argument respresents a field argument
type Argument struct {
	Name  string
	Value interface{}
}

type singleField struct {
	name string
	args []Argument
}

func (s *singleField) GetString() string {
	if s.args == nil {
		return s.name
	}

	sargs := s.name + " ("
	for _, arg := range s.args {
		s, ok := arg.Value.(string)
		if ok {
			sargs += fmt.Sprintf("%s: %v ", arg.Name, strconv.QuoteToASCII(s))
		} else {
			sargs += fmt.Sprintf("%s: %v ", arg.Name, arg.Value)
		}
	}

	sargs += ")"

	return sargs
}

type objectField struct {
	object *Object
}

func (o *objectField) GetString() string {
	return formatObject(o.object)
}
