package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jasosa/EasyGraph/EasyGraph/graphql"
)

func main() {

	c := graphql.NewClient("https://api.github.com/graphql")
	c.SetToken("edb547e97d20e321a30bbf2eda0463859bf4c683")

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

	c2 := graphql.NewClient2("https://api.github.com/graphql")
	c2.SetToken("edb547e97d20e321a30bbf2eda0463859bf4c683")

	q := c2.Query().Object("viewer").StringField("login")
	res, err := c2.Execute(q)
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
