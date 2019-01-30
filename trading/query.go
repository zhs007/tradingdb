package trading

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb"
	"github.com/zhs007/ankadb/graphqlext"
	pb "github.com/zhs007/tradingdb/proto"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"candleChunks": &graphql.Field{
				Type: candleChunkType,
				Args: graphql.FieldConfigArgument{
					"code": &graphql.ArgumentConfig{
						Type: graphql.String,
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

					code := params.Args["code"].(string)
					name := params.Args["name"].(string)
					st := params.Args["startTime"].(int64)
					et := params.Args["endTime"].(int64)

					tz := getStringFromMapEx(params.Args, "timeZone", "")
					loc, err := time.LoadLocation(tz)
					if err != nil {
						return nil, err
					}

					lst := countCandleChunkKeyID(code, name, st, et, loc)

					retcc := &pb.CandleChunk{
						Code:      code,
						Name:      name,
						KeyID:     "tmp",
						StartTime: st,
						EndTime:   et,
					}

					for _, v := range lst {
						buf, err := curdb.Get([]byte(v))
						if buf != nil && err == nil {
							cc := &pb.CandleChunk{}

							err = proto.Unmarshal(buf, cc)
							if err != nil {
								return nil, err
							}

							for _, c := range cc.Candles {
								retcc.Candles = append(retcc.Candles, c)
							}
						}
					}

					// idQuery, isOK := p.Args["id"].(string)
					// if isOK {
					// return data[idQuery], nil
					// }
					return retcc, nil
				},
			},
			"tradingData": &graphql.Field{
				Type: tradingDataType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
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
					buf, err := curdb.Get([]byte(keyid))
					td := &pb.TradingData{KeyID: keyid}

					err = proto.Unmarshal(buf, td)
					if err != nil {
						return nil, err
					}

					return td, nil
				},
			},
		},
	},
)
