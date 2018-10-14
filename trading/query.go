package trading

import (
	"github.com/golang/protobuf/proto"
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb"
	"github.com/zhs007/ankadb/err"
	"github.com/zhs007/ankadb/graphqlext"
	"github.com/zhs007/ankadb/proto"
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
						Type: graphql.String,
					},
					"startTime": &graphql.ArgumentConfig{
						Type: graphqlext.Timestamp,
					},
					"endTime": &graphql.ArgumentConfig{
						Type: graphqlext.Timestamp,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					curdb := ankadb.GetContextValueDatabase(params.Context, interface{}("curdb"))
					if curdb == nil {
						return nil, ankadberr.NewError(ankadbpb.CODE_CTX_CURDB_ERR)
					}

					code := params.Args["code"].(string)
					name := params.Args["name"].(string)
					st := params.Args["startTime"].(int64)
					et := params.Args["endTime"].(int64)

					lst := countKeyID(code, name, st, et)

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
								return nil, ankadberr.NewError(ankadbpb.CODE_PROTOBUF_ENCODE_ERR)
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
		},
	},
)
