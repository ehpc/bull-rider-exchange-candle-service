package testing

import (
	"github.com/golang/protobuf/proto"
	
	protoCandle "github.com/ehpc/bull-rider/protobuf/go/candle"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
)

// BinanceCandleExampleJSON is an example of Binance API candle
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
const BinanceCandleExampleJSON = `[
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

// BinanceIOTAUSDT15mCandleExampleProtobuf is an example of Binance
// candle converted to protobuf format
var BinanceIOTAUSDT15mCandleExampleProtobuf = protoCandle.Candle{
	Exchange: string(candle.ExchangeBinance),
	Pair: string(candle.PairIOTAUSDT),
	Interval: string(candle.Interval15m),
	OpenTime: 1561622400000,
	Open: 0.42590000,
	High: 0.42680000,
	Low: 0.40050000,
	Close: 0.42210000,
	Volume: 2589747.41000000,
	CloseTime: 1561636799999,
	QuoteVolume: 1073262.29118300,
	TradesCount: 5370,
}

// BinanceIOTAUSDT15mCandleExampleProtobufMarshaled is an example of Binance candle 
// converted to protobuf format and marshaled to []bytes
var BinanceIOTAUSDT15mCandleExampleProtobufMarshaled, _ = proto.Marshal(&BinanceIOTAUSDT15mCandleExampleProtobuf)

// GenerateCandlesJSON generates json with candles from Binance API
func GenerateCandlesJSON(candleJSON string, count int) string {
	json := "["
	for i := 0; i < count; i++ {
		json += candleJSON
		if i < count-1 {
			json += ","
		}
	}
	return json + "]"
}
