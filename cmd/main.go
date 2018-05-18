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

	fmt.Println("Testing structured queries...")

	qb := c.QueryBuilder()
	q := qb.AddObject("hero").AddSingleField("name").Query()
	fmt.Println("==> Query: ", q.GetString())
	res, err := c.Execute(q)
	parseResponse(res, err)

	qb = c.QueryBuilder()
	q = qb.AddObject("hero").AddSingleField("name").AddObjectField("friends").AddSingleField("name").Query()
	fmt.Println("==> Query: ", q.GetString())
	res, err = c.Execute(q)
	parseResponse(res, err)

	qb = c.QueryBuilder()
	fmt.Println("==> Add object fields with arguments")
	fmt.Println("==> Not implemented")
	fmt.Println("")

	qb = c.QueryBuilder()
	fmt.Println("==> Add single fields with arguments inside object fields with arguments")
	fmt.Println("==> Not implemented")
	fmt.Println("")

	qb = c.QueryBuilder()
	fmt.Println("==> Aliases")
	fmt.Println("==> Not implemented")
	fmt.Println("")

	fmt.Println("Testing raw queries...")

	for i := range rawQueries {
		q = qb.CreateRawQuery(rawQueries[i])
		fmt.Println("==> Query: ", q.GetString())
		res, err = c.Execute(q)
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
	`query {
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

/* func execNpmCommand(command string) {
	cmd := exec.Command("npm", "--prefix", "../../../../../repos/starwars-server", command)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
} */
