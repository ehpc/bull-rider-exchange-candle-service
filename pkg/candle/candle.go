package candle

// Candle is a trading candle
type Candle struct {
	Exchange               Exchange
	Pair                   Pair
	Interval               Interval
	OpenTime, CloseTime    int64
	Open, Close, High, Low float64
	Volume, QuoteVolume    float64
	TradesCount            int64
}

// NewCandle is a Candle constructor
func NewCandle(exchange Exchange, pair Pair, interval Interval,
	openTime int64, closeTime int64, open float64, close float64,
	high float64, low float64, volume float64, quoteVolume float64,
	tradesCount int64) Candle {
	return Candle{
		Exchange: exchange,
		Pair: pair,
		Interval: interval,
		OpenTime: openTime,
		CloseTime: closeTime,
		Open: open,
		Close: close,
		High: high,
		Low: low,
		Volume: volume,
		QuoteVolume: quoteVolume,
		TradesCount: tradesCount,
	}
}
