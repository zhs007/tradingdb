package trading

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"
)

func makeKeyID(code string, name string, startTime int64) string {
	tm := time.Unix(startTime, 0)

	if code == "pta" {
		ts := tm.Format("20060102")

		ch := tm.Hour()
		if ch >= 21 {
			return name + ts + "1"
		}

		return name + ts + "0"
	}

	return name + string(startTime)
}

func countKeyID(code string, name string, startTime int64, endTime int64) []string {
	var lst []string

	if endTime < startTime {
		return lst
	}

	if code == "pta" {
		stm := time.Unix(startTime, 0)
		etm := time.Unix(startTime, 0)
		hoff := int(etm.Sub(stm).Hours())
		doff := hoff / 24
		lst = make([]string, 0, doff*2)

		for startTime <= endTime {
			tm := time.Unix(startTime, 0)
			ts := tm.Format("20060102")

			lst = append(lst, name+ts+"0")
			lst = append(lst, name+ts+"1")

			startTime += 24 * 60 * 60
		}
	}

	return lst
}

// FuncForEachCSV - funcForEach(mapval map[string]interface{})
type FuncForEachCSV func(mapval map[string]interface{}) error

// ForEachCSV - for each csv file
func ForEachCSV(filename string, funcForEach FuncForEachCSV) error {
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
			break
		}

		if err != nil {
			return err
		}

		lines++
		if lines == 1 {
			for i, v := range record {
				if v == "" {
					mapHead["curTime"] = i
				} else {
					mapHead[v] = i
				}
			}

			continue
		}

		tm2, err := time.Parse("2006-01-02 15:04:05", record[mapHead["curTime"]])
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
	}

	return nil
}
