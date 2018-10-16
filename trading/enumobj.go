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
			"ORDERTYPE_INVALIDTYPE": &graphql.EnumValueConfig{
				Value:       pb.ORDERTYPE_INVALIDTYPE,
				Description: "invalid order type",
			},
			"ORDERTYPE_LIMIT": &graphql.EnumValueConfig{
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
			"ORDERSIDE_INVALIDSIDE": &graphql.EnumValueConfig{
				Value:       pb.ORDERSIDE_INVALIDSIDE,
				Description: "invalid side",
			},
			"ORDERSIDE_BUY": &graphql.EnumValueConfig{
				Value:       pb.ORDERSIDE_BUY,
				Description: "buy",
			},
			"ORDERSIDE_SELL": &graphql.EnumValueConfig{
				Value:       pb.ORDERSIDE_SELL,
				Description: "sell",
			},
		},
	},
)
