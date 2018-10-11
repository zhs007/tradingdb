package main

import (
	"context"
	"encoding/json"

	"github.com/zhs007/ankadb/client"
	"github.com/zhs007/tradingdb/trading"
)

var strNewChandleChunk = `mutation NewChandleChunk($code: String!, $name: String!, $startTime: Timestamp!, $endTime: Timestamp!) {
	newCandleChunk(code: $code, name: $name, startTime: $startTime, endTime: $endTime) {
	  keyID
	}
  }`

var strInsertChandle = `mutation InsertCandle($keyID: String!, $candle: CandleInput!) {
	insertCandle(keyID: $keyID, candle: $candle) {
	  keyID,
	  candles {
		curTime
	  }
	}
  }`

func importCSV(ctx context.Context, dbname string, code string, name string, filename string, c ankadbclient.AnkaClient) error {
	var lst [](map[string]interface{})

	trading.ForEachCSV(filename, func(mapval map[string]interface{}) error {
		if lst == nil {
			lst = append(lst, mapval)
		} else {
			lt := lst[len(lst)-1]["curTime"].(int64)
			ct := mapval["curTime"].(int64)
			if ct != lt+60 {
				st := lst[0]["curTime"].(int64)

				cc := make(map[string]interface{})

				cc["code"] = code
				cc["name"] = name
				cc["startTime"] = st
				cc["endTime"] = lt

				buf, err := json.Marshal(cc)
				if err != nil {
					return err
				}
				c.Query(ctx, dbname, strNewChandleChunk, string(buf))

				for _, v := range lst {
					buf, err := json.Marshal(v)
					if err != nil {
						return err
					}
					c.Query(ctx, dbname, strInsertChandle, string(buf))
				}

				lst = nil

				lst = append(lst, mapval)
			}
		}

		return nil
	})

	// fr, err := os.Open(filename)
	// if err != nil {
	// 	return err
	// }

	// csvr := csv.NewReader(fr)
	// lines := 0

	// for {
	// 	record, err := csvr.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	if err != nil {
	// 		return err
	// 	}

	// 	lines++
	// 	if lines == 1 {
	// 		continue
	// 	}

	// 	mapval := make(map[string]interface{})
	// 	mapval["code"] = record[0]

	// 	buf, err1 := json.Marshal(mapval)
	// 	if err1 != nil {
	// 		return err1
	// 	}

	// 	c.Query(ctx, dbname, strQuery, string(buf))
	// 	// fmt.Println(record)
	// }

	return nil
}

func main() {
	c := ankadbclient.NewClient()

	importCSV(context.Background(), "tradingdb", "pta", "pta1601", "TA601.csv", c)

	c.Start("0.0.0.0:7788")
}
