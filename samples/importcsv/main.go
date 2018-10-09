package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
	"os"

	"github.com/zhs007/ankadb/client"
)

func importCSV(ctx context.Context, dbname string, filename string, c ankadbclient.AnkaClient) error {
	strQuery := `mutation InsertCandles($code:String!, $name:String!, $candle:CandleInput){
		insertCandles(code:$code,name:$name,candle:$candle){
			code,
			name}
		}`
	fr, err := os.Open(filename)
	if err != nil {
		return err
	}

	csvr := csv.NewReader(fr)

	for {
		record, err := csvr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		mapval := make(map[string]interface{})
		mapval["code"] = record[0]

		buf, err1 := json.Marshal(mapval)
		if err1 != nil {
			return err1
		}

		c.Query(ctx, dbname, strQuery, string(buf))
		// fmt.Println(record)
	}

	return nil
}

func main() {
	c := ankadbclient.NewClient()
	c.Start("")
}
