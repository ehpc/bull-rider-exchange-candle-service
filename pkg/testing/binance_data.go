package testing

// BinanceCandleJSONElement is an example of Binance API candle
const BinanceCandleJSONElement = `[
	1561622400000,
	"0.42590000",
	"0.42680000",
	"0.40050000",
	"0.42210000",
	"2589747.41000000",
	1561636799999,
	"1073262.29118300",
	5370,
	"1185070.39000000",
	"492584.11707600",
	"0"
]`

// GenerateCandlesJSON generates json with candles from Binance API
func GenerateCandlesJSON(count int) string {
	json := "["
	for i := 0; i < count; i++ {
		json += BinanceCandleJSONElement
		if i < count - 1 {
			json += ","
		}
	}
	return json + "]"
}