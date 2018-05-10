package graphql

import (
	"strconv"
	"strings"
)

var fmtSeparator = " "

func formatRawQuery(q *rawQuery) string {
	quotedQuery := strconv.QuoteToASCII(q.stringQuery)
	return `{"query": ` + quotedQuery + `}`
}

func formatQuery(q *StructuredQuery) string {
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
