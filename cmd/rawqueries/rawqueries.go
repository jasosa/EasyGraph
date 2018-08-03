package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jasosa/EasyGraph"
)

func main() {

	fmt.Println("Initializing easygraph client...")
	c := easygraph.NewClient("http://localhost:8080/graphql")
	qb := c.QueryBuilder()

	fmt.Println("Testing raw queries...")
	for i := range rawQueries {
		q := qb.Query(rawQueries[i])
		fmt.Println("==> Query: ", q.GetString())
		res, err := c.Run(q)
		parseResponse(res, err)
	}
}

func parseResponse(res *http.Response, err error) {
	defer fmt.Println("")
	if err != nil {
		fmt.Println("==> Error: ", err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("==> Answer: ", string(body))
}

var rawQueries = []string{
	` {
		hero {
		  name
		}
	  }`,

	`query {
		hero {
		  name
		  # Queries can have comments!
		  friends {
			name
		  }
		}
	  }`,
	`query {
		human(id: "1000") {
		  name
		  height
		}
	  }`,
	`query {
		human(id: "1000") {
		  name
		  height(unit: FOOT)
		}
	  }`,
	`query {
		empireHero: hero(episode: EMPIRE) {
		  name
		}
		jediHero: hero(episode: JEDI) {
		  name
		}
	  }`,
}
