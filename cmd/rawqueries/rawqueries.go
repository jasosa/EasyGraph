package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jasosa/EasyGraph"
)

func main() {

	fmt.Println("######## Initializing easygraph client... ########")
	c := easygraph.NewClient("http://localhost:8080/graphql")
	qb := c.QueryBuilder()

	fmt.Println("######## Testing queries... ########")
	for i := range rawQueries {
		q := qb.Query(rawQueries[i])
		fmt.Println("==> Query: ", q.GetString())
		res, err := c.Run(q)
		parseResponse(res, err)
	}

	fmt.Println("######## Testing queries with variables... ########")
	q := qb.Query(heroNameAndFriendsQuery)
	q.AddVariable("episode", "JEDI")
	fmt.Println("==> Query: ", q.GetString())
	res, err := c.Run(q)
	parseResponse(res, err)

	q = qb.Query(heroWithDirectiveQuery)
	q.AddVariable("episode", "JEDI")
	q.AddVariable("withFriends", false)
	fmt.Println("==> Query: ", q.GetString())
	res, err = c.Run(q)
	parseResponse(res, err)

	q = qb.Query(createEpisodeReviewMutation)
	q.AddVariable("ep", "JEDI")
	type Review struct {
		Stars      int    `json:"stars"`
		Commentary string `json:"commentary"`
	}
	q.AddVariable("review", Review{Commentary: "Fantastic movie!", Stars: 5})
	fmt.Println("==> Query: ", q.GetString())
	res, err = c.Run(q)
	parseResponse(res, err)

	q = qb.Query(heroByEpisodeQuery)
	q.AddVariable("ep", "JEDI")
	fmt.Println("==> Query: ", q.GetString())
	res, err = c.Run(q)
	parseResponse(res, err)

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
	`{
			search(text: "an") {
				__typename
				... on Human {
					name
				}
				... on Droid {
					name
				}
				... on Starship {
					name
				}
			}
		}`,
}

var heroNameAndFriendsQuery = `query HeroNameAndFriends($episode: Episode) {
		hero(episode: $episode) {
			name
			friends {
				name
			}
		}
	}`

var heroWithDirectiveQuery = `query Hero($episode: Episode, $withFriends: Boolean!) {
		hero(episode: $episode) {
			name
			friends @include(if: $withFriends) {
				name
			}
		}
	}`

var createEpisodeReviewMutation = `mutation CreateReviewForEpisode($ep: Episode!, $review: ReviewInput!) {
  createReview(episode: $ep, review: $review) {
    stars
    commentary
  }
}`

var heroByEpisodeQuery = `query HeroForEpisode($ep: Episode!) {
  hero(episode: $ep) {
    name
    ... on Droid {
      primaryFunction
    }
    ... on Human {
      height
    }
  }
}`
