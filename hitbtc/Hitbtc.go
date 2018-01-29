package hitbtc

import (
	"errors"
	. "github.com/openbtc/RobotEa"
	//"log"
	"net/http"
	"time"
)

const (
	EXCHANGE_NAME = "hitbtc.com"

	API_BASE_URL = "https://api.hitbtc.com/"
	API_V2       = "api/2/"
	TICKER_URI   = "public/ticker/"
)

type Hitbtc struct {
	accessKey,
	secretKey string
	httpClient *http.Client
}

func New(client *http.Client, accessKey, secretKey string) *Hitbtc {
	return &Hitbtc{accessKey, secretKey, client}
}

func (hitbtc *Hitbtc) GetExchangeName() string {
	return EXCHANGE_NAME
}

func (hitbtc *Hitbtc) GetTicker(currency CurrencyPair) (*Ticker, error) {
	currency = hitbtc.adaptCurrencyPair(currency)
	curr := currency.ToSymbol("")
	tickerUri := API_BASE_URL + API_V2 + TICKER_URI + curr
	bodyDataMap, err := HttpGet(hitbtc.httpClient, tickerUri)
	//log.Println("Hitbtc bodyDataMap:", tickerUri, bodyDataMap, err)

	if err != nil {
		//log.Println(err)
		return nil, err
	}
	if result, isok := bodyDataMap["error"].(map[string]interface{}); isok == true {
		//log.Println("bodyDataMap[\"error\"]", result)
		return nil, errors.New(result["message"].(string) + ", " + result["description"].(string))
	}

	tickerMap := bodyDataMap
	var ticker Ticker

	timestamp := time.Now().Unix()
	ticker.Date = uint64(timestamp)
	ticker.Last = ToFloat64(tickerMap["last"])
	ticker.Buy = ToFloat64(tickerMap["bid"])
	ticker.Sell = ToFloat64(tickerMap["ask"])
	ticker.Low = ToFloat64(tickerMap["low"])
	ticker.High = ToFloat64(tickerMap["high"])
	ticker.Vol = ToFloat64(tickerMap["volume"])

	//log.Println("Hitbtc", currency, "ticker:", ticker)

	return &ticker, nil
}

func (hitbtc *Hitbtc) GetTickers(currency CurrencyPair) (*Ticker, error) {
	return hitbtc.GetTicker(currency)
}

func (hitbtc *Hitbtc) GetTickerInBuf(currency CurrencyPair) (*Ticker, error) {
	return hitbtc.GetTicker(currency)
}

func (hitbtc *Hitbtc) LimitBuy(amount, price string, currency CurrencyPair) (*Order, error) {
	panic("not implements")
}

func (hitbtc *Hitbtc) LimitSell(amount, price string, currency CurrencyPair) (*Order, error) {
	panic("not implements")
}

func (hitbtc *Hitbtc) MarketBuy(amount, price string, currency CurrencyPair) (*Order, error) {
	panic("not implements")
}

func (hitbtc *Hitbtc) MarketSell(amount, price string, currency CurrencyPair) (*Order, error) {
	panic("not implements")
}

func (hitbtc *Hitbtc) CancelOrder(orderId string, currency CurrencyPair) (bool, error) {
	panic("not implements")
}

func (hitbtc *Hitbtc) GetOneOrder(orderId string, currency CurrencyPair) (*Order, error) {
	panic("not implements")
}
func (hitbtc *Hitbtc) GetUnfinishOrders(currency CurrencyPair) ([]Order, error) {
	panic("not implements")
}

func (hitbtc *Hitbtc) GetOrderHistorys(currency CurrencyPair, currentPage, pageSize int) ([]Order, error) {
	panic("not implements")
}

func (hitbtc *Hitbtc) GetAccount() (*Account, error) {
	panic("not implements")
}

func (hitbtc *Hitbtc) GetDepth(size int, currency CurrencyPair) (*Depth, error) {
	panic("not implement")
}
func (hitbtc *Hitbtc) adaptCurrencyPair(pair CurrencyPair) CurrencyPair {
	var currencyA Currency
	var currencyB Currency

	if pair.CurrencyA == BCC {
		currencyA = BCH
	} else {
		currencyA = pair.CurrencyA
	}
	//currencyB = pair.BaseCurrency
	if pair.CurrencyB == USDT {
		currencyB = USD
	} else {
		currencyB = pair.CurrencyB
	}

	return NewCurrencyPair(currencyA, currencyB)
}

func (hitbtc *Hitbtc) GetKlineRecords(currency CurrencyPair, period, size, since int) ([]Kline, error) {
	panic("not implements")
}

//非个人，整个交易所的交易记录
func (hitbtc *Hitbtc) GetTrades(currencyPair CurrencyPair, since int64) ([]Trade, error) {
	panic("not implements")
}
