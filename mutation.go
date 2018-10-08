package main

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"newCandleChunk": &graphql.Field{
			Type:        candleChunkType,
			Description: "new candle chunk",
			Args: graphql.FieldConfigArgument{
				"code": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"startTime": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphqlext.Timestamp),
				},
				"endTime": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphqlext.Timestamp),
				},
				// "candle": &graphql.ArgumentConfig{
				// 	Type: candleInputType,
				// },
				// "candles": &graphql.ArgumentConfig{
				// 	Type: graphql.NewList(candleType),
				// },
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// ankadb.DBMgr
				// rand.Seed(time.Now().UnixNano())
				// product := Product{
				// 	ID:    int64(rand.Intn(100000)), // generate random ID
				// 	Name:  params.Args["name"].(string),
				// 	Info:  params.Args["info"].(string),
				// 	Price: params.Args["price"].(float64),
				// }
				// products = append(products, product)
				return nil, nil
			},
		},
	},
})
