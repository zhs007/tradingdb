package trading

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
)

// candleType - Candle
//		you can see trading.graphql
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
//		you can see trading.graphql
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

// candleChunkListType - CandleChunkList
//		you can see trading.graphql
var candleChunkListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CandleChunkList",
		Fields: graphql.Fields{
			"keyIDs": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
			},
		},
	},
)

// orderType - Order
//		you can see trading.graphql
var orderType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"orderID": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"orderType": &graphql.Field{
				Type: graphql.NewNonNull(enumORDERTYPEType),
			},
			"orderSide": &graphql.Field{
				Type: graphql.NewNonNull(enumORDERSIDEType),
			},
			"price": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"volume": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"startTime": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Timestamp),
			},
			"avgPrice": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"doneVolume": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"doneTime": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Timestamp),
			},
		},
	},
)

// tradeType - Trade
//		you can see trading.graphql
var tradeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Trade",
		Fields: graphql.Fields{
			"tradeID": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"orderID": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"curTime": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Timestamp),
			},
			"price": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
			"volume": &graphql.Field{
				Type: graphql.NewNonNull(graphqlext.Int64),
			},
		},
	},
)

// tradingDataType - TradingData
//		you can see trading.graphql
var tradingDataType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TradingData",
		Fields: graphql.Fields{
			"keyID": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"orders": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(orderType)),
			},
			"trades": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(tradeType)),
			},
		},
	},
)
