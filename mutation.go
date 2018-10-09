package main

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
	},
})
