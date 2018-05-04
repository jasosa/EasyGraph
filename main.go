package main

import (
	"fmt"
	"github/jasosa/EasyGraph/graphql"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	c := graphql.NewClient("https://api.github.com/graphql")
	c.SetToken("4a5e08f8e85f809a0c7ba278f0db572d27ae0821")

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
	}
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
	`query { viewer { login }}`,
	`query {repository(owner: "jasosa", name:"StringCalculator") {
		id
	}}`,
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
