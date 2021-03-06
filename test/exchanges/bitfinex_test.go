package test

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

import (
	"log"
	"testing"

	"github.com/bitontop/gored/coin"
	"github.com/bitontop/gored/exchange"
	"github.com/bitontop/gored/pair"

	"github.com/bitontop/gored/exchange/bitfinex"
	"github.com/bitontop/gored/test/conf"
	// "../../exchange/bitfinex"
	// "../conf"
)

/********************Public API********************/
func Test_Bitfinex(t *testing.T) {
	e := InitBitfinex()

	pair := pair.GetPairByKey("BTC|ETH")

	// Test_Coins(e)
	// Test_Pairs(e)
	Test_Pair(e, pair)
	// Test_Orderbook(e, pair)
	// Test_ConstraintFetch(e, pair)
	// Test_Constraint(e, pair)

	Test_Balance(e, pair)
	// Test_Trading(e, pair, 0.0001000123, 100)
	// Test_Trading_Sell(e, pair, 0.06123456789, 0.02)
	// Test_OrderStatus(e, pair, "40797839810") // 40797839810 an open order
	// Test_CancelOrder(e, pair, "41094576119")
	// Test_Withdraw(e, pair.Base, 1, "ADDRESS")
	// log.Printf("Url: %v", e.GetTradingWebURL(pair))
	// Test_DoWithdraw(e, pair.Target, "1", "0x37E0Fc27C6cDB5035B2a3d0682B4E7C05A4e6C46", "tag")
}

func Test_BITFINEX_TradeHistory(t *testing.T) {
	e := InitBitfinex()
	p := pair.GetPairByKey("BTC|ETH")

	opTradeHistory := &exchange.PublicOperation{
		Type:      exchange.TradeHistory,
		EX:        e.GetName(),
		Pair:      p,
		DebugMode: true,
	}

	err := e.LoadPublicData(opTradeHistory)
	if err != nil {
		log.Printf("%v", err)
	}

	log.Printf("TradeHistory: %s::%s", opTradeHistory.EX, opTradeHistory.Pair.Name)

	for _, d := range opTradeHistory.TradeHistory {
		log.Printf(">> %+v ", d)
	}
}

func InitBitfinex() exchange.Exchange {
	coin.Init()
	pair.Init()
	// utils.GetCommonDataFromJSON("https://raw.githubusercontent.com/bitontop/gored/master/data")

	config := &exchange.Config{}
	config.Source = exchange.EXCHANGE_API
	// config.Source = exchange.JSON_FILE
	// config.SourceURI = "https://raw.githubusercontent.com/bitontop/gored/master/data"
	conf.Exchange(exchange.BITFINEX, config)

	ex := bitfinex.CreateBitfinex(config)
	log.Printf("Initial [ %v ] ", ex.GetName())

	config = nil
	return ex
}
