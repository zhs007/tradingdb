package trading

import (
	"time"

	"github.com/goinggo/mapstructure"
	"github.com/golang/protobuf/proto"
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb"
	"github.com/zhs007/ankadb/graphqlext"
	pb "github.com/zhs007/tradingdb/proto"
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
				"timeZone": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				anka := ankadb.GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
				if anka == nil {
					return nil, ankadb.ErrCtxAnkaDB
				}

				curdb := anka.GetDBMgr().GetDB("candles")
				if curdb == nil {
					return nil, ankadb.ErrCtxCurDB
				}

				tz := getStringFromMapEx(params.Args, "timeZone", "")
				loc, err := time.LoadLocation(tz)
				if err != nil {
					return nil, err
				}

				code := params.Args["code"].(string)
				name := params.Args["name"].(string)
				st := params.Args["startTime"].(int64)
				keyid := makeCandleChunkKeyID(code, name, st, loc)

				cc := &pb.CandleChunk{
					Code:      code,
					Name:      name,
					StartTime: st,
					EndTime:   params.Args["endTime"].(int64),
					KeyID:     keyid,
				}

				data, err := proto.Marshal(cc)
				if err != nil {
					return nil, err
				}

				err = curdb.Put([]byte(keyid), data)
				if err != nil {
					return nil, err
				}

				return cc, nil
			},
		},
		"insertCandle": &graphql.Field{
			Type:        candleChunkType,
			Description: "insert candle",
			Args: graphql.FieldConfigArgument{
				"keyID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"candle": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(candleInputType),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				anka := ankadb.GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
				if anka == nil {
					return nil, ankadb.ErrCtxAnkaDB
				}

				curdb := anka.GetDBMgr().GetDB("candles")
				if curdb == nil {
					return nil, ankadb.ErrCtxCurDB
				}

				keyID := params.Args["keyID"].(string)

				buf, err := curdb.Get([]byte(keyID))
				cc := &pb.CandleChunk{}

				err = proto.Unmarshal(buf, cc)
				if err != nil {
					return nil, err
				}

				// name := params.Args["name"].(string)
				ci := params.Args["candle"].(map[string]interface{})

				var c pb.Candle
				if err := mapstructure.Decode(ci, &c); err != nil {
					return nil, err
				}

				// c := &pb.Candle{
				// 	CurTime:      ci["curTime"].(int64),
				// 	Open:         ci["open"].(int64),
				// 	Close:        ci["close"].(int64),
				// 	High:         ci["high"].(int64),
				// 	Low:          ci["low"].(int64),
				// 	Volume:       ci["volume"].(int64),
				// 	OpenInterest: ci["openInterest"].(int64),
				// }

				cc.Candles = append(cc.Candles, &c)
				// keyid := makeKeyID(code, name, st)

				// cc := &pb.CandleChunk{}

				data, err := proto.Marshal(cc)
				if err != nil {
					return nil, err
				}

				err = curdb.Put([]byte(keyID), data)
				if err != nil {
					return nil, err
				}

				return cc, nil
			},
		},
		"insertCandles": &graphql.Field{
			Type:        candleChunkType,
			Description: "insert candles",
			Args: graphql.FieldConfigArgument{
				"keyID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"candles": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.NewList(candleInputType)),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				anka := ankadb.GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
				if anka == nil {
					return nil, ankadb.ErrCtxAnkaDB
				}

				curdb := anka.GetDBMgr().GetDB("candles")
				if curdb == nil {
					return nil, ankadb.ErrCtxCurDB
				}

				keyID := params.Args["keyID"].(string)

				buf, err := curdb.Get([]byte(keyID))
				cc := &pb.CandleChunk{}

				err = proto.Unmarshal(buf, cc)
				if err != nil {
					return nil, err
				}

				// name := params.Args["name"].(string)
				lstci := params.Args["candles"].([]interface{})

				for _, cv := range lstci {
					ci := cv.(map[string]interface{})

					var c pb.Candle
					if err := mapstructure.Decode(ci, &c); err != nil {
						return nil, err
					}
					// c := &pb.Candle{
					// 	CurTime:      ci["curTime"].(int64),
					// 	Open:         ci["open"].(int64),
					// 	Close:        ci["close"].(int64),
					// 	High:         ci["high"].(int64),
					// 	Low:          ci["low"].(int64),
					// 	Volume:       ci["volume"].(int64),
					// 	OpenInterest: ci["openInterest"].(int64),
					// }

					cc.Candles = append(cc.Candles, &c)
				}

				// keyid := makeKeyID(code, name, st)

				// cc := &pb.CandleChunk{}

				data, err := proto.Marshal(cc)
				if err != nil {
					return nil, err
				}

				err = curdb.Put([]byte(keyID), data)
				if err != nil {
					return nil, err
				}

				return cc, nil
			},
		},
		"clearCandleChunk": &graphql.Field{
			Type:        candleChunkListType,
			Description: "clear candlechunks",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				anka := ankadb.GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
				if anka == nil {
					return nil, ankadb.ErrCtxAnkaDB
				}

				curdb := anka.GetDBMgr().GetDB("candles")
				if curdb == nil {
					return nil, ankadb.ErrCtxCurDB
				}

				name := params.Args["name"].(string)

				lst := pb.CandleChunkList{}

				curit := curdb.NewIteratorWithPrefix([]byte(name + ":"))
				for curit.Next() {
					key := curit.Key()

					lst.KeyIDs = append(lst.KeyIDs, string(key))
				}
				curit.Release()
				err := curit.Error()
				if err != nil {
					return lst, err
				}

				return lst, nil
			},
		},
		"setTradingData": &graphql.Field{
			Type:        tradingDataType,
			Description: "set TradingData",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"orders": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.NewList(orderInputType)),
				},
				"trades": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.NewList(tradeInputType)),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				anka := ankadb.GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
				if anka == nil {
					return nil, ankadb.ErrCtxAnkaDB
				}

				curdb := anka.GetDBMgr().GetDB("trades")
				if curdb == nil {
					return nil, ankadb.ErrCtxCurDB
				}

				name := params.Args["name"].(string)
				keyid := makeTradingDataKeyID(name)

				cc := &pb.TradingData{
					KeyID: keyid,
				}

				lstorder := params.Args["orders"].([]interface{})

				for _, cv := range lstorder {
					ci := cv.(map[string]interface{})

					var c pb.Order
					if err := mapstructure.Decode(ci, &c); err != nil {
						return nil, err
					}

					cc.Orders = append(cc.Orders, &c)
				}

				lsttrade := params.Args["trades"].([]interface{})

				for _, cv := range lsttrade {
					ci := cv.(map[string]interface{})

					var c pb.Trade
					if err := mapstructure.Decode(ci, &c); err != nil {
						return nil, err
					}

					cc.Trades = append(cc.Trades, &c)
				}

				data, err := proto.Marshal(cc)
				if err != nil {
					return nil, err
				}

				err = curdb.Put([]byte(keyid), data)
				if err != nil {
					return nil, err
				}

				return cc, nil
			},
		},
	},
})
