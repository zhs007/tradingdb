package main

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"candleChunks": &graphql.Field{
				Type: candleType,
				Args: graphql.FieldConfigArgument{
					"code": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"startTime": &graphql.ArgumentConfig{
						Type: graphqlext.Timestamp,
					},
					"endTime": &graphql.ArgumentConfig{
						Type: graphqlext.Timestamp,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// idQuery, isOK := p.Args["id"].(string)
					// if isOK {
					// return data[idQuery], nil
					// }
					return nil, nil
				},
			},
		},
	},
)
