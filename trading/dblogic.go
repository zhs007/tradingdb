package trading

import (
	"context"

	"github.com/graphql-go/graphql"

	"github.com/zhs007/ankadb"
)

// tradingDB -
type tradingDB struct {
	schema graphql.Schema
}

// NewTradingDB -
func NewTradingDB() ankadb.DBLogic {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
			// Types:    curTypes,
		},
	)

	return &tradingDB{
		schema: schema,
	}
}

// OnQuery -
func (logic *tradingDB) OnQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error) {
	result := graphql.Do(graphql.Params{
		Schema:         logic.schema,
		RequestString:  request,
		VariableValues: values,
		Context:        ctx,
	})
	// if len(result.Errors) > 0 {
	// 	fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	// }

	return result, nil
}

// OnQueryStream -
func (logic *tradingDB) OnQueryStream(ctx context.Context, request string, values map[string]interface{}, funcOnQueryStream ankadb.FuncOnQueryStream) error {
	return nil
}
