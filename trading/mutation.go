package trading

import (
	"time"

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

				loc, err := time.LoadLocation("Asia/Shanghai")
				if err != nil {
					return nil, err
				}

				code := params.Args["code"].(string)
				name := params.Args["name"].(string)
				st := params.Args["startTime"].(int64)
				keyid := makeKeyID(code, name, st, loc)

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
					Volume:       ci["volume"].(int64),
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
		"insertCandles": &graphql.Field{
			Type:        candleChunkType,
			Description: "insert candles",
			Args: graphql.FieldConfigArgument{
				"keyID": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"candles": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.NewList(candleInputType)),
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
				lstci := params.Args["candles"].([]interface{})

				for _, cv := range lstci {
					ci := cv.(map[string]interface{})

					c := &pb.Candle{
						CurTime:      ci["curTime"].(int64),
						Open:         ci["open"].(int64),
						Close:        ci["close"].(int64),
						High:         ci["high"].(int64),
						Low:          ci["low"].(int64),
						Volume:       ci["volume"].(int64),
						OpenInterest: ci["openInterest"].(int64),
					}

					cc.Candles = append(cc.Candles, c)
				}

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
		"clearCandleChunk": &graphql.Field{
			Type:        candleChunkListType,
			Description: "clear candlechunks",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				curdb := ankadb.GetContextValueDatabase(params.Context, interface{}("curdb"))
				if curdb == nil {
					return nil, ankadberr.NewError(ankadbpb.CODE_CTX_CURDB_ERR)
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
	},
})
