package graphql

import (
	"strconv"
	"strings"
)

var fmtSeparator = " "

func formatQuery(q Query) string {
	objectsAndFields := []string{}
	objectsAndFields = append(objectsAndFields, `query { `)
	for _, object := range q.getObjects() {
		objectsAndFields = append(objectsAndFields, formatObject(object))
	}

	objectsAndFields = append(objectsAndFields, `}`)
	query := strings.Join(objectsAndFields, "")
	asciiQuery := strconv.QuoteToASCII(query)
	return `{"query": ` + asciiQuery + `}`
}

func formatObject(object *Object) string {
	var stringsObject []string
	stringsObject = append(stringsObject, object.name)
	stringsObject = append(stringsObject, `{`)
	for _, field := range object.fields {
		stringsObject = append(stringsObject, field.GetString())
	}
	stringsObject = append(stringsObject, `}`)
	return strings.Join(stringsObject, fmtSeparator)
}
