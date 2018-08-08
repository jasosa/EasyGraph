# Easy Graph

## What it this?

EasyGraph is a very small library to interact with GraphQL APIs. Probably is a library that is solving
a non-existent problem for anyone apart of me! :D The motivations behind it were basically two:
    * Make my life a bit easier in a project I'm working currently in my job
    * Improve my knowledge in Golang

## What is the status?

EasyGraph is, as far as I know, fully functional. You can execute any kind of query against a Graphql API.  


 ## How it works?

// First you need to initialize the client and create a query builder:
	
    c := easygraph.NewClient(graphqlapiurl)
    qb := c.QueryBuilder()
    

// Then you can create a query doing the following:
   
    q := qb.Query(` {
	hero {
	        name
	    }
	}`)

// Finally you can use the clien to run the query:

	res, err := c.Run(q)

// If you need variables in your query you can use the following method:

    q.AddVariable("episode", "JEDI")
  
// If the API needs Bearer token authentication you can add the token to the client in the 
    following way (before executing a query):
    
    c.SetToken(token)