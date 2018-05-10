package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jasosa/EasyGraph/EasyGraph/graphql"
)

func main() {

	/* c := graphql.NewClient("https://api.github.com/graphql")
	c.SetToken("752ef8aa4fe9ee69adfac01edc769fae7ac06bcc")

	for _, q := range queries {
		res, err := c.DoQuery(q)
		parse(res, err)
	}

	for _, q := range queriesWithVariables {
		v := graphql.Variable{
			Name:  "number_of_repos",
			Value: 3,
		}
		res, err := c.DoQueryWithVariables(q, v)
		parse(res, err)
	} */

	fmt.Println("################ client2 #############")

	c2 := graphql.NewClient2("https://api.github.com/graphql")
	c2.SetToken("752ef8aa4fe9ee69adfac01edc769fae7ac06bcc")

	q := c2.QueryBuilder().CreateRawQuery(`query {viewer{login}}`)
	res, err := c2.Execute(q)
	parse(res, err)

	q = c2.QueryBuilder().CreateRawQuery(`mutation { addStar (input: {
	   		starrableId: "MDEwOlJlcG9zaXRvcnkxODk3MTY5MQ==",
	   		clientMutationId: jasosa
	   	}){clientMutationId }
		   }`)
	res, err = c2.Execute(q)
	parse(res, err)

	qb := c2.QueryBuilder().AddObject("viewer").AddSingleFieldWithArguments("avatarUrl", graphql.Argument{Name: "size", Value: 512})
	res, err = c2.Execute(qb.Query())
	parse(res, err)

	qb = c2.QueryBuilder().AddObject("viewer").AddSingleField("email").AddSingleField("bio")
	res, err = c2.Execute(qb.Query())
	parse(res, err)
}

func parse(res *http.Response, err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}

var queries = []string{
	`query {viewer{login}}`,
	`query { meta { isPasswordAuthenticationVerifiable } }`,
	`query {repository(owner: "jasosa", name:"StringCalculator"){id}}`,
	`mutation { addStar (input: {
	   		starrableId: "MDEwOlJlcG9zaXRvcnkxODk3MTY5MQ==",
	   		clientMutationId: jasosa
	   	}){clientMutationId }
		   }`,
}

var queriesWithVariables = []string{
	`query ($number_of_repos:Int!) {
		viewer {
		  name
		   repositories(last: $number_of_repos) {
			 nodes {
			   name
			 }
		   }
		 }
	  }`,
}
