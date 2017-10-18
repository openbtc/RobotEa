package unit

import (
	. "rest"
	"github.com/stretchr/testify/assert"
	"testing"
    "rest/okcoin"
)

func Test_OKCoin_CN(t *testing.T) {
    var api API;
    api = okcoin.New("okcoin_cn", "", "");
    tk, err := api.GetTicker(BTC_CNY);
    if err != nil{
        t.Logf("%s", err.Error());
    }
    assert.True(t, err == nil);
    t.Logf("last:%f buy:%f sell:%f high:%f low:%f vol:%f date:%d",
        tk.Last, tk.Buy, tk.Sell, tk.High, tk.Low, tk.Vol, tk.Date);
}