package trading

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	pb "github.com/zhs007/tradingdb/proto"
)

func makeTradingDataKeyID(nameid string) string {
	return "td:" + nameid
}

func makeCandleChunkKeyID(code string, name string, startTime int64, loc *time.Location) string {
	tm := time.Unix(startTime, 0).In(loc)

	if code == "pta" {
		ts := tm.Format("20060102")

		ch := tm.Hour()
		if ch >= 21 {
			return "cc:" + name + ":" + ts + "2"
		} else if ch >= 13 {
			return "cc:" + name + ":" + ts + "1"
		}

		return "cc:" + name + ":" + ts + "0"
	}

	return "cc:" + name + ":" + string(startTime)
}

func countCandleChunkKeyID(code string, name string, startTime int64, endTime int64, loc *time.Location) []string {
	var lst []string

	if endTime < startTime {
		return lst
	}

	if code == "pta" {
		stm := time.Unix(startTime, 0).In(loc)
		etm := time.Unix(startTime, 0).In(loc)
		hoff := int(etm.Sub(stm).Hours())
		doff := hoff / 24
		lst = make([]string, 0, doff*2)

		for startTime <= endTime {
			tm := time.Unix(startTime, 0).In(loc)
			ts := tm.Format("20060102")

			lst = append(lst, "cc:"+name+":"+ts+"0")
			lst = append(lst, "cc:"+name+":"+ts+"1")
			lst = append(lst, "cc:"+name+":"+ts+"2")

			startTime += 24 * 60 * 60
		}
	}

	return lst
}

// FuncForEachCSV - funcForEach(mapval map[string]interface{})
type FuncForEachCSV func(mapval map[string]interface{}) error

// ForEachCSV - for each csv file
func ForEachCSV(filename string, local string, funcForEach FuncForEachCSV) error {
	loc, err := time.LoadLocation(local)
	if err != nil {
		return err
	}

	fr, err := os.Open(filename)
	if err != nil {
		return err
	}

	csvr := csv.NewReader(fr)
	lines := 0
	mapHead := make(map[string]int)
	// csveof := false

	for {
		record, err := csvr.Read()
		if err == io.EOF {
			// csveof = true
			err1 := funcForEach(nil)
			if err1 != nil {
				return err1
			}

			break
		} else if err != nil {
			return err
		}

		lines++
		if lines == 1 {
			for i, v := range record {
				// if v == "" {
				// 	mapHead["curTime"] = i
				// } else {
				mapHead[v] = i
				// }
			}

			continue
		}

		tm2, err := time.ParseInLocation("2006-01-02 15:04:05", record[mapHead["curtime"]], loc)
		if err != nil {
			return err
		}

		open, err := strconv.ParseFloat(record[mapHead["open"]], 64)
		if err != nil {
			return err
		}

		close, err := strconv.ParseFloat(record[mapHead["close"]], 64)
		if err != nil {
			return err
		}

		high, err := strconv.ParseFloat(record[mapHead["high"]], 64)
		if err != nil {
			return err
		}

		low, err := strconv.ParseFloat(record[mapHead["low"]], 64)
		if err != nil {
			return err
		}

		volume, err := strconv.ParseFloat(record[mapHead["volume"]], 64)
		if err != nil {
			return err
		}

		oi, err := strconv.ParseFloat(record[mapHead["oi"]], 64)
		if err != nil {
			return err
		}

		// use CandleInput
		mapval := make(map[string]interface{})
		mapval["curTime"] = tm2.Unix()
		mapval["open"] = int64(open * 100)
		mapval["close"] = int64(close * 100)
		mapval["high"] = int64(high * 100)
		mapval["low"] = int64(low * 100)
		mapval["volume"] = int64(volume * 100)
		mapval["openInterest"] = int64(oi * 100)

		err1 := funcForEach(mapval)
		if err1 != nil {
			return err1
		}

		// if csveof {
		// 	break
		// }
	}

	return nil
}

// ForEachOrderCSV - for each order csv file
func ForEachOrderCSV(filename string, local string, funcForEach FuncForEachCSV) error {
	loc, err := time.LoadLocation(local)
	if err != nil {
		return err
	}

	fr, err := os.Open(filename)
	if err != nil {
		return err
	}

	csvr := csv.NewReader(fr)
	lines := 0
	mapHead := make(map[string]int)

	for {
		record, err := csvr.Read()
		if err == io.EOF {
			err1 := funcForEach(nil)
			if err1 != nil {
				return err1
			}

			break
		} else if err != nil {
			return err
		}

		lines++
		if lines == 1 {
			for i, v := range record {
				mapHead[v] = i
			}

			continue
		}

		cti, err := strconv.ParseInt(record[mapHead["createtime"]], 10, 64)
		if err != nil {
			return err
		}

		ct := time.Unix(cti, 0).In(loc)

		tti, err := strconv.ParseInt(record[mapHead["tradetime"]], 10, 64)
		if err != nil {
			return err
		}

		tt := time.Unix(tti, 0).In(loc)

		price, err := strconv.ParseFloat(record[mapHead["price"]], 64)
		if err != nil {
			return err
		}

		volume, err := strconv.ParseFloat(record[mapHead["volume"]], 64)
		if err != nil {
			return err
		}

		avgprice, err := strconv.ParseFloat(record[mapHead["avgprice"]], 64)
		if err != nil {
			return err
		}

		lastvolume, err := strconv.ParseFloat(record[mapHead["lastvolume"]], 64)
		if err != nil {
			return err
		}

		// use OrderInput
		mapval := make(map[string]interface{})
		mapval["orderID"] = record[mapHead["id"]]
		mapval["orderType"] = orderType2GraphEnum(str2OrderType(record[mapHead["ordertype"]]))
		mapval["orderSide"] = orderSide2GraphEnum(str2OrderSide(record[mapHead["orderside"]]))
		mapval["price"] = int64(price * 100)
		mapval["volume"] = int64(volume * 100)
		mapval["startTime"] = ct.Unix()
		mapval["avgPrice"] = int64(avgprice * 100)
		mapval["doneVolume"] = int64((volume - lastvolume) * 100)
		mapval["doneTime"] = tt.Unix()

		err1 := funcForEach(mapval)
		if err1 != nil {
			return err1
		}
	}

	return nil
}

// ForEachTradeCSV - for each trade csv file
func ForEachTradeCSV(filename string, local string, funcForEach FuncForEachCSV) error {
	loc, err := time.LoadLocation(local)
	if err != nil {
		return err
	}

	fr, err := os.Open(filename)
	if err != nil {
		return err
	}

	csvr := csv.NewReader(fr)
	lines := 0
	mapHead := make(map[string]int)

	for {
		record, err := csvr.Read()
		if err == io.EOF {
			err1 := funcForEach(nil)
			if err1 != nil {
				return err1
			}

			break
		} else if err != nil {
			return err
		}

		lines++
		if lines == 1 {
			for i, v := range record {
				mapHead[v] = i
			}

			continue
		}

		cti, err := strconv.ParseInt(record[mapHead["tradetime"]], 10, 64)
		if err != nil {
			return err
		}

		ct := time.Unix(cti, 0).In(loc)

		price, err := strconv.ParseFloat(record[mapHead["price"]], 64)
		if err != nil {
			return err
		}

		volume, err := strconv.ParseFloat(record[mapHead["volume"]], 64)
		if err != nil {
			return err
		}

		// use OrderInput
		mapval := make(map[string]interface{})
		mapval["tradeID"] = record[mapHead["id"]]
		mapval["orderID"] = record[mapHead["orderid"]]
		mapval["curTime"] = ct.Unix()
		mapval["price"] = int64(price * 100)
		mapval["volume"] = int64(volume * 100)

		err1 := funcForEach(mapval)
		if err1 != nil {
			return err1
		}
	}

	return nil
}

func getStringFromMapEx(m map[string]interface{}, k string, defval string) string {
	v, ok := m[k]
	if !ok {
		return defval
	}

	return v.(string)
}

func str2OrderType(str string) pb.ORDERTYPE {
	if str == "limit" {
		return pb.ORDERTYPE_LIMIT
	}

	return pb.ORDERTYPE_INVALIDTYPE
}

func str2OrderSide(str string) pb.ORDERSIDE {
	if str == "buy" {
		return pb.ORDERSIDE_BUY
	} else if str == "sell" {
		return pb.ORDERSIDE_SELL
	}

	return pb.ORDERSIDE_INVALIDSIDE
}

func orderType2GraphEnum(ot pb.ORDERTYPE) string {
	return pb.ORDERTYPE_name[int32(ot)]
}

func orderSide2GraphEnum(os pb.ORDERSIDE) string {
	return pb.ORDERSIDE_name[int32(os)]
}
