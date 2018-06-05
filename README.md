# Easy Graph

## What it this?

EasyGraph is a little library to interact with GraphQL APIs in a slightly structured way. Probably is a library that is solving
a non-existent problem for anyone apart of me! :D The motivations behind it were basically two:
    * Make my life a bit easier in a project I'm working currently in my job
    * Improve my knowledge in Golang

## What is the status?

 Easy Graph it's still in a very early stage of development. If you want to contribute, you are welcome!

 EasyGraph is based on the following documentation: https://graphql.org/learn/  

 It supports two different modes of working: RawQueries and StructuredQueries. Using RawQueries is similar to write a GraphQL
  query but with a little layer of abstraction on it. StructuredQueries wants to simplify even more the creation of GraphQL queries. Both modes are in a very early stage of development. 

Using RawQueries currently allows to:
* Create and execute queries with simple fields 
* Create and execute queries with nested fields
* Create and execute queries with arguments
* Create and execute queries with aliases

Not implemented in RawQueries:
* Fragments
* Variables
* Directives
* Mutations
* Inline Fragments

Using StructuredQueries currently allows to:
* Create and execute queries with simple fields 
* Create and execute queries with nested fields

Not implemented in StructuredQueries:
* Arguments
* Aliases
* Fragments
* Variables
* Directives
* Mutations
* Inline Fragments

Also implemented:
* Support for Bearer token authentication




 ## How it works?

As mentioned before, there are two ways of execute queries, but the starting point is the same for both:


    // First you need to initialize the client and create a query builder:
	
        c := easygraph.NewClient(graphqlapiurl)
	    qb := c.QueryBuilder()
    

    // Then you can create a RawQuery doing the following:
   
        q := qb.CreateRawQuery(myquery)
   

    // being myquery, for example:
    
    `query {
		hero {
		  name
		}
	  }`

    // Finally you can use the clien to run the query:
        
        res, err := c.Execute(q)

    // Creating the same query using StructuredQuery will require to do the following, once you have initialized the client:
    
        qb := qb.AddObject("hero") // Add the object you want to query, this returns the QueryBuilder again
        qb := q.AddSingleField("name") // Add the field/s you want to request
        q := qb.Query() // Get the query
        res, err := c.Execute(q) // Execute it
    

    // If you need to use Bearer token authentication you can do it using the client in the 
        following way (before executing a query):
    
        c.SetToken(token)
     

 ## How I run the tests?

TBD
