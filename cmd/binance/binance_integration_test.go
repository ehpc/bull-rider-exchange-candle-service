// +build integration

package main

import (
	"testing"
	"math/rand"
	"strconv"
	"log"
	"os"

	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"

	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/binanceapi"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candle"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/candlemodel"
	"github.com/ehpc/bull-rider-exchange-candle-service/pkg/transport"
)

func TestMain(m *testing.M) {
	// Loading configuration
	err := godotenv.Load("../../.env")
	if err != nil {
		dir, _ := os.Getwd()
		log.Fatal(err, dir)
	}
	os.Exit(m.Run())
}

func TestIntegratedRESTFlow(t *testing.T) {
	// Creating API transport
	apiTransport := transport.NewHTTPTransport(os.Getenv("BINANCE_REST_API_URL"))
	defer apiTransport.Close()
	// Fetching data from Binance
	api := binanceapi.NewAPI(apiTransport, nil)
	candles, err := api.GetCandles(
		[]candle.Pair{candle.PairIOTAUSDT},
		[]candle.Interval{candle.Interval15m},
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, candles)
	assert.Greater(t, candles[0].OpenTime, int64(0))
	assert.Greater(t, candles[0].CloseTime, int64(0))
	// Creating model transport
	exchangeName := "test" + strconv.Itoa(rand.Int())
	modelTransport, err := transport.NewRabbitMQTransport(
		os.Getenv("MESSAGE_BROKER_URL"),
		exchangeName,
		exchangeName,
		transport.RabbitMQTransportOptions{
			Temporary: true,
		},
	)
	assert.NoError(t, err)
	// Pushing data to recipients
	model := candlemodel.NewCandleModel(modelTransport)
	result, err := model.AddCandles(candles)
	assert.NoError(t, err)
	assert.True(t, result)
	err = modelTransport.Close()
	assert.NoError(t, err)
}

func TestIntegratedWebsocketFlow(t *testing.T) {

}