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
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				curdb := ankadb.GetContextValueDatabase(params.Context, interface{}("curdb"))
				if curdb == nil {
					return nil, ankadberr.NewError(ankadbpb.CODE_CTX_CURDB_ERR)
				}

				code := params.Args["code"].(string)
				name := params.Args["name"].(string)
				st := params.Args["startTime"].(int64)
				keyid := makeKeyID(code, name, st)

				cc := &pb.CandleChunk{
					Code:      code,
					Name:      name,
					StartTime: st,
					EndTime:   params.Args["endTime"].(int64),
					KeyID:     keyid,
				}

				data, err := proto.Marshal(cc)
				if err != nil {
					return nil, ankadberr.NewError(ankadbpb.CODE_PROTOBUF_ENCODE_ERR)
				}

				err = curdb.Put([]byte(keyid), data)
				if err != nil {
					return nil, ankadberr.NewError(ankadbpb.CODE_DB_PUT_ERR)
				}

				return cc, nil
			},
		},
		"insertCandle": &graphql.Field{
			Type:        candleChunkType,
			Description: "insert candle",
			Args: graphql.FieldConfigArgument{
				"keyID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"candle": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(candleInputType),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				curdb := ankadb.GetContextValueDatabase(params.Context, interface{}("curdb"))
				if curdb == nil {
					return nil, ankadberr.NewError(ankadbpb.CODE_CTX_CURDB_ERR)
				}

				keyID := params.Args["keyID"].(string)

				buf, err := curdb.Get([]byte(keyID))
				cc := &pb.CandleChunk{}

				err = proto.Unmarshal(buf, cc)
				if err != nil {
					return nil, ankadberr.NewError(ankadbpb.CODE_PROTOBUF_ENCODE_ERR)
				}

				// name := params.Args["name"].(string)
				ci := params.Args["candle"].(map[string]interface{})

				c := &pb.Candle{
					CurTime:      ci["curTime"].(int64),
					Open:         ci["open"].(int64),
					Close:        ci["close"].(int64),
					High:         ci["high"].(int64),
					Low:          ci["low"].(int64),
					Volume:       ci["close"].(int64),
					OpenInterest: ci["openInterest"].(int64),
				}

				cc.Candles = append(cc.Candles, c)
				// keyid := makeKeyID(code, name, st)

				// cc := &pb.CandleChunk{}

				data, err := proto.Marshal(cc)
				if err != nil {
					return nil, ankadberr.NewError(ankadbpb.CODE_PROTOBUF_ENCODE_ERR)
				}

				err = curdb.Put([]byte(keyID), data)
				if err != nil {
					return nil, ankadberr.NewError(ankadbpb.CODE_DB_PUT_ERR)
				}

				return cc, nil
			},
		},
	},
})
