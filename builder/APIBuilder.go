package builder

import (
	"context"
	. "github.com/openbtc/RobotEa"
	"github.com/openbtc/RobotEa/binance"
	"github.com/openbtc/RobotEa/bitfinex"
	"github.com/openbtc/RobotEa/bitstamp"
	"github.com/openbtc/RobotEa/btcbox"
	"github.com/openbtc/RobotEa/chbtc"
	"github.com/openbtc/RobotEa/coincheck"
	"github.com/openbtc/RobotEa/huobi"
	"github.com/openbtc/RobotEa/kraken"
	"github.com/openbtc/RobotEa/okcoin"
	"github.com/openbtc/RobotEa/poloniex"
	"github.com/openbtc/RobotEa/yunbi"
	"github.com/openbtc/RobotEa/zaif"
	"net"
	"net/http"
	"net/url"
	"time"
	"github.com/openbtc/RobotEa/bittrex"
	"github.com/openbtc/RobotEa/bithumb"
	"github.com/openbtc/RobotEa/gdax"
	"github.com/openbtc/RobotEa/gateio"
	"github.com/openbtc/RobotEa/wex"
	"github.com/openbtc/RobotEa/zb"
)

type APIBuilder struct {
	client      *http.Client
	httpTimeout time.Duration
	apiKey      string
	secretkey   string
	clientId    string
}

func NewAPIBuilder() (builder *APIBuilder) {
	_client := http.DefaultClient
	transport := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 4 * time.Second,
	}
	_client.Transport = transport
	return &APIBuilder{client: _client}
}

func NewCustomAPIBuilder(client *http.Client) (builder *APIBuilder) {
	return &APIBuilder{client: client}
}

func (builder *APIBuilder) APIKey(key string) (_builder *APIBuilder) {
	builder.apiKey = key
	return builder
}

func (builder *APIBuilder) APISecretkey(key string) (_builder *APIBuilder) {
	builder.secretkey = key
	return builder
}

func (builder *APIBuilder) HttpProxy(proxyUrl string) (_builder *APIBuilder) {
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		return
	}
	transport := builder.client.Transport.(*http.Transport)
	transport.Proxy = http.ProxyURL(proxy)
	return builder
}

func (builder *APIBuilder) ClientID(id string) (_builder *APIBuilder) {
	builder.clientId = id
	return builder
}

func (builder *APIBuilder) HttpTimeout(timeout time.Duration) (_builder *APIBuilder) {
	builder.httpTimeout = timeout
	builder.client.Timeout = timeout
	transport := builder.client.Transport.(*http.Transport)
	if transport != nil {
		transport.ResponseHeaderTimeout = timeout
		transport.TLSHandshakeTimeout = timeout
		transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		}
	}
	return builder
}

func (builder *APIBuilder) Build(exName string) (api API) {
	var _api API
	switch exName {
	case "okcoin.cn":
		_api = okcoin.New(builder.client, builder.apiKey, builder.secretkey)
	case "huobi.com":
		_api = huobi.New(builder.client, builder.apiKey, builder.secretkey)
	case "chbtc.com":
		_api = chbtc.New(builder.client, builder.apiKey, builder.secretkey)
	case "yunbi.com":
		_api = yunbi.New(builder.client, builder.apiKey, builder.secretkey)
	case "poloniex.com":
		_api = poloniex.New(builder.client, builder.apiKey, builder.secretkey)
	case "okcoin.com":
		_api = okcoin.NewCOM(builder.client, builder.apiKey, builder.secretkey)
	case "coincheck.com":
		_api = coincheck.New(builder.client, builder.apiKey, builder.secretkey)
	case "zaif.jp":
		_api = zaif.New(builder.client, builder.apiKey, builder.secretkey)
	case "bitstamp.net":
		_api = bitstamp.NewBitstamp(builder.client, builder.apiKey, builder.secretkey, builder.clientId)
	case "huobi.pro":
		_api = huobi.NewHuobiPro(builder.client, builder.apiKey, builder.secretkey, builder.clientId)
	case "okex.com":
		_api = okcoin.NewOKExSpot(builder.client, builder.apiKey, builder.secretkey)
	case "bitfinex.com":
		_api = bitfinex.New(builder.client, builder.apiKey, builder.secretkey)
	case "kraken.com":
		_api = kraken.New(builder.client, builder.apiKey, builder.secretkey)
	case "binance.com":
		_api = binance.New(builder.client, builder.apiKey, builder.secretkey)
	case "btcbox.co.jp":
		_api = btcbox.New(builder.client, builder.apiKey, builder.secretkey)
	case "bittrex.com":
		_api = bittrex.New(builder.client, builder.apiKey, builder.secretkey)
	case "bithumb.com":
		_api = bithumb.New(builder.client, builder.apiKey, builder.secretkey)
	case "gdax.com":
		_api = gdax.New(builder.client , builder.apiKey , builder.secretkey)
	case "gate.io":
		_api = gateio.New(builder.client , builder.apiKey , builder.secretkey)
	case "wex.nz":
		_api = wex.New(builder.client , builder.apiKey ,builder.secretkey)
	case "zb.com":
		_api = zb.New(builder.client , builder.apiKey, builder.secretkey)
	default:
		panic("exchange name error ["+exName+"].")

	}
	return _api
}
