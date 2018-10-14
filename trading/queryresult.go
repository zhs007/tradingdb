package trading

// ResultCandleChunks -
type ResultCandleChunks struct {
	Data struct {
		CandleChunks struct {
			Candles []struct {
				CurTime      int64 `json:"curTime"`
				Open         int64 `json:"open"`
				Close        int64 `json:"close"`
				High         int64 `json:"high"`
				Low          int64 `json:"low"`
				Volume       int64 `json:"volume"`
				OpenInterest int64 `json:"openInterest"`
			} `json:"candles"`

			StartTime int64 `json:"starttime"`
			EndTime   int64 `json:"endtime"`
		} `json:"candleChunks"`
	} `json:"data"`
}
