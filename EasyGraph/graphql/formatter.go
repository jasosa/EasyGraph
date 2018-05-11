package graphql

import (
	"fmt"
	"strconv"
	"strings"
)

var fmtSeparator = " "

func formatStructuredQuery(q *StructuredQuery) string {
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

func formatRawQuery(q *rawQuery) string {
	var formattedQuery string
	if len(q.variables) > 0 {
		variablesQuery := formatVariables(q.variables)
		formattedQuery = formatQueryWithvariables(
			strconv.QuoteToASCII(q.stringQuery),
			variablesQuery)
	} else {
		formattedQuery = formatQuery(q.stringQuery)
	}

	return formattedQuery
}

func formatVariables(variables []Variable) string {
	queryVariables := []string{}
	var keyvalue string
	for _, v := range variables {
		s, ok := v.Value.(string)
		if ok {
			keyvalue = fmt.Sprintf(`%s:%v`, strconv.QuoteToASCII(v.Name), strconv.QuoteToASCII(s))
		} else {
			keyvalue = fmt.Sprintf(`%s:%v`, strconv.QuoteToASCII(v.Name), v.Value)
		}
		queryVariables = append(queryVariables, keyvalue)
	}
	return strings.Join(queryVariables, ",")
}

func formatQueryWithvariables(queryString string, variablesString string) string {
	query := `{"query": ` + queryString + `,` +
		`"variables": {` + variablesString + `}}`
	return query
}

func formatQuery(queryString string) string {
	quotedQuery := strconv.QuoteToASCII(queryString)
	return `{"query": ` + quotedQuery + `}`
}
