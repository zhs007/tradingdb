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
