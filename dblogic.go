package main

import (
	"context"

	"github.com/graphql-go/graphql"

	"github.com/zhs007/ankadb"
)

// var curTypes = []graphql.Type{candleType}

type tradingDB struct {
	schema graphql.Schema
}

func newTradingDB() ankadb.DBLogic {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
			// Types:    curTypes,
		},
	)

	return &tradingDB{
		schema: schema,
	}
}

func (logic *tradingDB) OnQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error) {
	result := graphql.Do(graphql.Params{
		Schema:         logic.schema,
		RequestString:  request,
		VariableValues: values,
		Context:        ctx,
	})
	// if len(result.Errors) > 0 {
	// 	fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	// }

	return result, nil
}

// var schema, _ = graphql.NewSchema(
// 	graphql.SchemaConfig{
// 		Query:    queryType,
// 		Mutation: mutationType,
// 		// Types:    curTypes,
// 	},
// )

// func executeQuery(query string, mapvar map[string]interface{}, schema graphql.Schema) *graphql.Result {
// 	result := graphql.Do(graphql.Params{
// 		Schema:         schema,
// 		RequestString:  query,
// 		VariableValues: mapvar,
// 	})
// 	if len(result.Errors) > 0 {
// 		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
// 	}
// 	return result
// }
