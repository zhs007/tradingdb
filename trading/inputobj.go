package trading

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
)

// candleInputType - Candle
//		you can see trading.graphql
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

// orderInputType - Order
//		you can see trading.graphql
var orderInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "OrderInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"orderID": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"orderType": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(enumORDERTYPEType),
			},
			"orderSide": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(enumORDERSIDEType),
			},
			"price": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"volume": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"startTime": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Timestamp),
			},
			"avgPrice": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"doneVolume": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"doneTime": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Timestamp),
			},
		},
	},
)

// tradeInputType - Trade
//		you can see trading.graphql
var tradeInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "TradeInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"tradeID": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"orderID": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"curTime": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Timestamp),
			},
			"price": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"volume": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
		},
	},
)
