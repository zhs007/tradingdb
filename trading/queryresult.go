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

// ResultTradingData -
type ResultTradingData struct {
	Data struct {
		TradingData struct {
			KeyID string `json:"keyID"`

			Orders []struct {
				OrderID   string `json:"orderID"`
				OrderType string `json:"orderType"`
				OrderSide string `json:"orderSide"`
				Price     int64  `json:"price"`
				Volume    int64  `json:"volume"`
				StartTime int64  `json:"startTime"`

				AvgPrice   int64 `json:"avgPrice"`
				DoneVolume int64 `json:"doneVolume"`
				DoneTime   int64 `json:"doneTime"`
			} `json:"orders"`

			Trades []struct {
				TradeID string `json:"tradeID"`
				OrderID string `json:"orderID"`
				CurTime int64  `json:"curTime"`
				Price   int64  `json:"price"`
				Volume  int64  `json:"volume"`
			} `json:"trades"`
		} `json:"tradingData"`
	} `json:"data"`
}
