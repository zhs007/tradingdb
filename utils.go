package main

import "time"

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
