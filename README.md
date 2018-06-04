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
  query but with a little layer of abstraction on it. StructuredQueries wants to simplify even more the creation of GraphQL queries. Both modes
  are in a very early stage of development. 

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

 ## How it works?

TBD

 ## How I run the tests?

TBD
