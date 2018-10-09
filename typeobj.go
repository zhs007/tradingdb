package main

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
)

// candleType - Candle
var candleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Candle",
		Fields: graphql.Fields{
			"curTime": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Timestamp),
			},
			"open": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"close": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"high": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"low": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"volume": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"openInterest": &graphql.Field{
				Type: graphqlext.Int64,
			},
		},
	},
)

// candleChunkType - CandleChunk
var candleChunkType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CandleChunk",
		Fields: graphql.Fields{
			"code": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"keyID": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"startTime": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Timestamp),
			},
			"endTime": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Timestamp),
			},
			"candles": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(candleType)),
			},
		},
	},
)
