package trading

import (
	"github.com/graphql-go/graphql"
	pb "github.com/zhs007/tradingdb/proto"
)

var enumORDERTYPEType = graphql.NewEnum(
	graphql.EnumConfig{
		Name:        "ORDERTYPE",
		Description: "Order Type",
		Values: graphql.EnumValueConfigMap{
			"INVALIDTYPE": &graphql.EnumValueConfig{
				Value:       pb.ORDERTYPE_INVALIDTYPE,
				Description: "invalid order type",
			},
			"LIMIT": &graphql.EnumValueConfig{
				Value:       pb.ORDERTYPE_LIMIT,
				Description: "limit order",
			},
		},
	},
)

var enumORDERSIDEType = graphql.NewEnum(
	graphql.EnumConfig{
		Name:        "ORDERSIDE",
		Description: "Order Side",
		Values: graphql.EnumValueConfigMap{
			"INVALIDSIDE": &graphql.EnumValueConfig{
				Value:       pb.ORDERSIDE_INVALID_ORDERSIDE,
				Description: "invalid side",
			},
			"BUY": &graphql.EnumValueConfig{
				Value:       pb.ORDERSIDE_ORDER_BUY,
				Description: "buy",
			},
			"SELL": &graphql.EnumValueConfig{
				Value:       pb.ORDERSIDE_ORDER_SELL,
				Description: "sell",
			},
		},
	},
)
