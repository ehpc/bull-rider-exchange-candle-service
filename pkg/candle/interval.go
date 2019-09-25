package candle

// Interval is an interval for a candle in seconds
type Interval string

// All possible intervals
const (
	Interval1m Interval = "1m"
	Interval3m = "3m"
	Interval5m = "5m"
	Interval15m = "15m"
	Interval30m = "30m"
	Interval1h = "1h"
	Interval2h = "2h"
	Interval4h = "4h"
	Interval6h = "6h"
	Interval8h = "8h"
	Interval12h = "12h"
	Interval1d = "1d"
	Interval3d = "3d"
	Interval1w = "1w"
	Interval1M = "1M"
)
