package main

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
)

var candleInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CandleInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"curTime": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Timestamp),
			},
			"open": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"close": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"high": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"low": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"volume": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"openInterest": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
		},
	},
)
