package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zhs007/ankadb/client"
	"github.com/zhs007/ankadb/err"
	"github.com/zhs007/ankadb/proto"
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

var strInsertChandles = `mutation InsertCandles($keyID: String!, $candles: [CandleInput]!) {
	insertCandles(keyID: $keyID, candles: $candles) {
	  keyID,
	  candles {
		curTime
	  }
	}
  }`

func insCSVData(ctx context.Context, dbname string, code string, name string, lst [](map[string]interface{}), c ankadbclient.AnkaClient, lt int64) error {
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

	queryReply, err := c.Query(ctx, dbname, strNewChandleChunk, string(buf))
	if err != nil {
		return err
	}

	if queryReply.Code != ankadbpb.CODE_OK {
		return ankadberr.NewError(queryReply.Code)
	}

	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(queryReply.Result), &mapResult); err != nil {
		return err
	}

	var mapData map[string]interface{}
	var ok bool
	if mapData, ok = mapResult["data"].(map[string]interface{}); !ok {
		return ankadberr.NewError(ankadbpb.CODE_RESULT_NO_DATA)
	}

	var mapNewCC map[string]interface{}
	if mapNewCC, ok = mapData["newCandleChunk"].(map[string]interface{}); !ok {
		return ankadberr.NewError(ankadbpb.CODE_RESULT_DATA_INVALID)
	}

	var curKeyID string
	if curKeyID, ok = mapNewCC["keyID"].(string); !ok {
		return ankadberr.NewError(ankadbpb.CODE_RESULT_DATA_INVALID)
	}

	// curKeyID := retNewCC.(string)

	// fmt.Print(queryReply)

	// for _, v := range lst {
	// 	mapInsC := make(map[string]interface{})
	// 	mapInsC["keyID"] = curKeyID
	// 	mapInsC["candle"] = v

	// 	buf, err := json.Marshal(mapInsC)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	queryReply, err := c.Query(ctx, dbname, strInsertChandle, string(buf))

	// 	if queryReply.Code != ankadbpb.CODE_OK {
	// 		return ankadberr.NewError(queryReply.Code)
	// 	}
	// }

	mapInsC := make(map[string]interface{})
	mapInsC["keyID"] = curKeyID
	mapInsC["candles"] = lst

	buf, err1 := json.Marshal(mapInsC)
	if err1 != nil {
		return err1
	}

	queryReply1, err1 := c.Query(ctx, dbname, strInsertChandles, string(buf))

	if queryReply1.Code != ankadbpb.CODE_OK {
		return ankadberr.NewError(queryReply1.Code)
	}

	fmt.Printf("queryReply " + queryReply1.Result + "\n")
	fmt.Printf("key " + curKeyID + " ok!\n")

	return nil
}

func importCSV(ctx context.Context, dbname string, code string, name string, local string, filename string, c ankadbclient.AnkaClient) error {
	var lst [](map[string]interface{})

	trading.ForEachCSV(filename, local, func(mapval map[string]interface{}) error {
		if mapval == nil {
			if lst == nil {
				return nil
			}

			lt := lst[len(lst)-1]["curTime"].(int64)

			err := insCSVData(ctx, dbname, code, name, lst, c, lt)
			if err != nil {
				return err
			}

			lst = nil
		} else if lst == nil {
			lst = append(lst, mapval)
		} else {
			lt := lst[len(lst)-1]["curTime"].(int64)
			ct := mapval["curTime"].(int64)
			if ct != lt+60 {

				err := insCSVData(ctx, dbname, code, name, lst, c, lt)
				if err != nil {
					return err
				}

				// st := lst[0]["curTime"].(int64)

				// cc := make(map[string]interface{})

				// cc["code"] = code
				// cc["name"] = name
				// cc["startTime"] = st
				// cc["endTime"] = lt

				// buf, err := json.Marshal(cc)
				// if err != nil {
				// 	return err
				// }

				// queryReply, err := c.Query(ctx, dbname, strNewChandleChunk, string(buf))
				// if err != nil {
				// 	return err
				// }

				// if queryReply.Code != ankadbpb.CODE_OK {
				// 	return ankadberr.NewError(queryReply.Code)
				// }

				// var mapResult map[string]interface{}
				// if err := json.Unmarshal([]byte(queryReply.Result), &mapResult); err != nil {
				// 	return err
				// }

				// var mapData map[string]interface{}
				// var ok bool
				// if mapData, ok = mapResult["data"].(map[string]interface{}); !ok {
				// 	return ankadberr.NewError(ankadbpb.CODE_RESULT_NO_DATA)
				// }

				// var mapNewCC map[string]interface{}
				// if mapNewCC, ok = mapData["newCandleChunk"].(map[string]interface{}); !ok {
				// 	return ankadberr.NewError(ankadbpb.CODE_RESULT_DATA_INVALID)
				// }

				// var curKeyID string
				// if curKeyID, ok = mapNewCC["keyID"].(string); !ok {
				// 	return ankadberr.NewError(ankadbpb.CODE_RESULT_DATA_INVALID)
				// }

				// // curKeyID := retNewCC.(string)

				// // fmt.Print(queryReply)

				// // for _, v := range lst {
				// // 	mapInsC := make(map[string]interface{})
				// // 	mapInsC["keyID"] = curKeyID
				// // 	mapInsC["candle"] = v

				// // 	buf, err := json.Marshal(mapInsC)
				// // 	if err != nil {
				// // 		return err
				// // 	}

				// // 	queryReply, err := c.Query(ctx, dbname, strInsertChandle, string(buf))

				// // 	if queryReply.Code != ankadbpb.CODE_OK {
				// // 		return ankadberr.NewError(queryReply.Code)
				// // 	}
				// // }

				// mapInsC := make(map[string]interface{})
				// mapInsC["keyID"] = curKeyID
				// mapInsC["candles"] = lst

				// buf, err1 := json.Marshal(mapInsC)
				// if err1 != nil {
				// 	return err1
				// }

				// queryReply1, err1 := c.Query(ctx, dbname, strInsertChandles, string(buf))

				// if queryReply1.Code != ankadbpb.CODE_OK {
				// 	return ankadberr.NewError(queryReply1.Code)
				// }

				// fmt.Printf("queryReply " + queryReply1.Result + "\n")
				// fmt.Printf("key " + curKeyID + " ok!\n")

				lst = nil
			}

			lst = append(lst, mapval)
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

	c.Start("0.0.0.0:7788")

	importCSV(context.Background(), "tradingdb", "pta", "pta1601", "Asia/Shanghai", "TA601F.csv", c)
}
