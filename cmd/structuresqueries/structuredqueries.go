package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jasosa/EasyGraph"
)

type FOOT struct {
}

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
	q = qb.AddObject("human").AddSingleField("name").AddSingleFieldWithArguments("height", easygraph.Argument{Name: "unit", Value: "\"FOOT\""}).Query()
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

}

func parseResponse(res *http.Response, err error) {
	defer fmt.Println("")
	if err != nil {
		fmt.Println("==> Error: ", err)
	}
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("==> Answer: ", string(body))
}
