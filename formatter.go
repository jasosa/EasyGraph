package easygraph

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

var fmtSeparator = " "

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

func formatVariables(variables []variable) string {
	queryVariables := []string{}
	var keyvalue string
	for _, v := range variables {
		bytesValue, _ := json.Marshal(v.Value)
		keyvalue = fmt.Sprintf(`%s:%v`, strconv.QuoteToASCII(v.Name), string(bytesValue))
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
