package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// var curTypes = []graphql.Type{candleType}

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
		// Types:    curTypes,
	},
)

func executeQuery(query string, mapvar map[string]interface{}, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  query,
		VariableValues: mapvar,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
