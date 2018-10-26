package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/zhs007/ankadb/client"
	"github.com/zhs007/tradingdb/trading"
)

var strNewChandleChunk = `mutation NewChandleChunk($code: String!, $name: String!, $startTime: Timestamp!, $endTime: Timestamp!, $timeZone: String) {
	newCandleChunk(code: $code, name: $name, startTime: $startTime, endTime: $endTime, timeZone: $timeZone) {
	  keyID
	}
  }`

var strInsertChandle = `mutation InsertCandle($keyID: ID!, $candle: CandleInput!) {
	insertCandle(keyID: $keyID, candle: $candle) {
	  keyID,
	  candles {
		curTime
	  }
	}
  }`

var strInsertChandles = `mutation InsertCandles($keyID: ID!, $candles: [CandleInput]!) {
	insertCandles(keyID: $keyID, candles: $candles) {
	  keyID,
	  candles {
		curTime
	  }
	}
  }`

var strSetTradingData = `mutation SetTradingData($name: ID!, $orders: [OrderInput]!, $trades: [TradeInput]!) {
	setTradingData(name: $name, orders: $orders, trades: $trades) {
	  keyID,
	  orders {
		orderID
	  },
	  trades {
		tradeID
	  }	  
	}
  }`

func insCandles(ctx context.Context, code string, name string, lst [](map[string]interface{}), tz string, c ankadbclient.AnkaClient, lt int64) error {
	st := lst[0]["curTime"].(int64)

	cc := make(map[string]interface{})

	cc["code"] = code
	cc["name"] = name
	cc["startTime"] = st
	cc["endTime"] = lt
	cc["timeZone"] = tz

	buf, err := json.Marshal(cc)
	if err != nil {
		return err
	}

	queryReply, err := c.Query(ctx, strNewChandleChunk, string(buf))
	if err != nil {
		return err
	}

	fmt.Print(queryReply.Result)

	if queryReply.Err != "" {
		return errors.New(queryReply.Err)
	}

	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(queryReply.Result), &mapResult); err != nil {
		return err
	}

	var mapData map[string]interface{}
	var ok bool
	if mapData, ok = mapResult["data"].(map[string]interface{}); !ok {
		return errors.New("result no data")
	}

	var mapNewCC map[string]interface{}
	if mapNewCC, ok = mapData["newCandleChunk"].(map[string]interface{}); !ok {
		return errors.New("result data invalid")
	}

	var curKeyID string
	if curKeyID, ok = mapNewCC["keyID"].(string); !ok {
		return errors.New("result data invalid")
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

	queryReply1, err1 := c.Query(ctx, strInsertChandles, string(buf))

	if queryReply1.Err != "" {
		return errors.New(queryReply1.Err)
	}

	fmt.Printf("queryReply " + queryReply1.Result + "\n")
	fmt.Printf("key " + curKeyID + " ok!\n")

	return nil
}

func importCSV(ctx context.Context, code string, name string, local string, filename string, c ankadbclient.AnkaClient) error {
	var lst [](map[string]interface{})

	trading.ForEachCSV(filename, local, func(mapval map[string]interface{}) error {
		if mapval == nil {
			if lst == nil {
				return nil
			}

			lt := lst[len(lst)-1]["curTime"].(int64)

			err := insCandles(ctx, code, name, lst, local, c, lt)
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

				err := insCandles(ctx, code, name, lst, local, c, lt)
				if err != nil {
					return err
				}

				lst = nil
			}

			lst = append(lst, mapval)
		}

		return nil
	})

	return nil
}

func insTradingData(ctx context.Context, name string, lstorder [](map[string]interface{}), lsttrade [](map[string]interface{}), c ankadbclient.AnkaClient) error {
	cc := make(map[string]interface{})

	cc["name"] = name
	cc["orders"] = lstorder
	cc["trades"] = lsttrade

	buf, err := json.Marshal(cc)
	if err != nil {
		return err
	}

	queryReply, err := c.Query(ctx, strSetTradingData, string(buf))
	if err != nil {
		return err
	}

	fmt.Print(queryReply.Result)

	if queryReply.Err != "" {
		return errors.New(queryReply.Err)
	}

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

	// mapInsC := make(map[string]interface{})
	// mapInsC["keyID"] = curKeyID
	// mapInsC["candles"] = lst

	// buf, err1 := json.Marshal(mapInsC)
	// if err1 != nil {
	// 	return err1
	// }

	// queryReply1, err1 := c.Query(ctx, strInsertChandles, string(buf))

	// if queryReply1.Code != ankadbpb.CODE_OK {
	// 	return ankadberr.NewError(queryReply1.Code)
	// }

	fmt.Printf("queryReply " + queryReply.Result + "\n")
	// fmt.Printf("key " + curKeyID + " ok!\n")

	return nil
}

func importTradingData(ctx context.Context, name string, local string, orderfn string, tradefn string, c ankadbclient.AnkaClient) error {
	var lstorder [](map[string]interface{})
	var lsttrade [](map[string]interface{})

	err := trading.ForEachOrderCSV(orderfn, local, func(mapval map[string]interface{}) error {
		if mapval != nil {
			lstorder = append(lstorder, mapval)
		}

		return nil
	})
	if err != nil {
		fmt.Print(err)
	}

	err = trading.ForEachTradeCSV(tradefn, local, func(mapval map[string]interface{}) error {
		if mapval != nil {
			lsttrade = append(lsttrade, mapval)
		}

		return nil
	})
	if err != nil {
		fmt.Print(err)
	}

	insTradingData(ctx, name, lstorder, lsttrade, c)

	return nil
}

func main() {
	c := ankadbclient.NewClient()

	// c.Start("0.0.0.0:7788")
	c.Start("47.90.46.159:7788")

	// importCSV(context.Background(), "pta", "pta1601", "Asia/Shanghai", "TA601F.csv", c)
	// importCSV(context.Background(), "pta", "pta1605", "Asia/Shanghai", "TA605F.csv", c)
	// importCSV(context.Background(), "pta", "pta1609", "Asia/Shanghai", "TA609F.csv", c)
	// importCSV(context.Background(), "pta", "pta1701", "Asia/Shanghai", "TA701F.csv", c)
	// importCSV(context.Background(), "pta", "pta1705", "Asia/Shanghai", "TA705F.csv", c)
	// importCSV(context.Background(), "pta", "pta1709", "Asia/Shanghai", "TA709F.csv", c)
	// importCSV(context.Background(), "pta", "pta1801", "Asia/Shanghai", "TA801F.csv", c)
	// importCSV(context.Background(), "pta", "pta1805", "Asia/Shanghai", "TA805F.csv", c)
	// importCSV(context.Background(), "pta", "pta1809", "Asia/Shanghai", "TA809F.csv", c)
	// importCSV(context.Background(), "pta", "pta1901", "Asia/Shanghai", "TA901F.csv", c)

	importTradingData(context.Background(), "zhs007-001", "Asia/Shanghai", "order.csv", "trade.csv", c)
}
