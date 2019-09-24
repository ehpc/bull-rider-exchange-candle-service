package candle

// Interval is an interval for a candle
type Interval uint8

// All possible intervals
const (
	Interval1m Interval = iota
	Interval3m
	Interval5m
	Interval15m
	Interval30m
	Interval1h
	Interval2h
	Interval4h
	Interval6h
	Interval8h
	Interval12h
	Interval1d
	Interval3d
	Interval1w
	Interval1M
)
