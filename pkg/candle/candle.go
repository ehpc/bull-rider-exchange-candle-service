package candle

// Candle is a trading candle
type Candle struct {
	Exchange               Exchange
	Pair                   Pair
	Interval               Interval
	OpenTime, CloseTime    uint64
	Open, Close, High, Low float64
	Volume, QuoteVolume    float64
	TradesNumber           uint64
}
