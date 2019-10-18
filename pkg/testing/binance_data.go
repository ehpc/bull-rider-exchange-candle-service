package testing

import (
	"github.com/golang/protobuf/proto"

	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	protoCandle "github.com/ehpc/bull-rider/protobuf/go/candle"
)

// BinanceCandleExampleRESTJSON is an example of Binance REST API candle
/*
1499040000000,      // Open time
"0.01634790",       // Open
"0.80000000",       // High
"0.01575800",       // Low
"0.01577100",       // Close
"148976.11427815",  // Volume
1499644799999,      // Close time
"2434.19055334",    // Quote asset volume
308,                // Number of trades
"1756.87402397",    // Taker buy base asset volume
"28.46694368",      // Taker buy quote asset volume
"17928899.62484339" // Ignore.
*/
const BinanceCandleExampleRESTJSON = `[
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

// BinanceCandleExampleWebsocketJSON is an example of Binance Websocket API candle
/*
{
	"e": "kline",     // Event type
	"E": 123456789,   // Event time
	"s": "BNBBTC",    // Symbol
	"k": {
	  "t": 123400000, // Kline start time
	  "T": 123460000, // Kline close time
	  "s": "BNBBTC",  // Symbol
	  "i": "1m",      // Interval
	  "f": 100,       // First trade ID
	  "L": 200,       // Last trade ID
	  "o": "0.0010",  // Open price
	  "c": "0.0020",  // Close price
	  "h": "0.0025",  // High price
	  "l": "0.0015",  // Low price
	  "v": "1000",    // Base asset volume
	  "n": 100,       // Number of trades
	  "x": false,     // Is this kline closed?
	  "q": "1.0000",  // Quote asset volume
	  "V": "500",     // Taker buy base asset volume
	  "Q": "0.500",   // Taker buy quote asset volume
	  "B": "123456"   // Ignore
	}
}
*/
const BinanceCandleExampleWebsocketJSON = `{
	"e": "kline",
	"E": 123456789,
	"s": "BNBBTC",
	"k": {
	  "t": 123400000,
	  "T": 123460000,
	  "s": "BNBBTC",
	  "i": "1m",
	  "f": 100,
	  "L": 200,
	  "o": "0.0010",
	  "c": "0.0020",
	  "h": "0.0025",
	  "l": "0.0015",
	  "v": "1000",
	  "n": 100,
	  "x": false,
	  "q": "1.0000",
	  "V": "500",
	  "Q": "0.500",
	  "B": "123456"
	}
}`

// BinanceIOTAUSDT15mCandleExampleProtobuf is an example of Binance
// candle converted to protobuf format
var BinanceIOTAUSDT15mCandleExampleProtobuf = protoCandle.Candle{
	Exchange:    string(candle.ExchangeBinance),
	Pair:        string(candle.PairIOTAUSDT),
	Interval:    string(candle.Interval15m),
	OpenTime:    1561622400000,
	Open:        0.42590000,
	High:        0.42680000,
	Low:         0.40050000,
	Close:       0.42210000,
	Volume:      2589747.41000000,
	CloseTime:   1561636799999,
	QuoteVolume: 1073262.29118300,
	TradesCount: 5370,
}

// BinanceIOTAUSDT15mCandleExampleProtobufMarshaled is an example of Binance candle
// converted to protobuf format and marshaled to []bytes
var BinanceIOTAUSDT15mCandleExampleProtobufMarshaled, _ = proto.Marshal(&BinanceIOTAUSDT15mCandleExampleProtobuf)

// GenerateCandlesJSON generates json with candles from Binance API
func GenerateCandlesJSON(CandleJSON string, count int) string {
	json := "["
	for i := 0; i < count; i++ {
		json += CandleJSON
		if i < count-1 {
			json += ","
		}
	}
	return json + "]"
}
