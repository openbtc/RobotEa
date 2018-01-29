package hitbtc

import (
	"github.com/openbtc/RobotEa"
	"net/http"
	"testing"
)

var htb = New(http.DefaultClient, "", "")

func TestWex_GetTicker(t *testing.T) {
	ticker, err := htb.GetTicker(goex.BTC_USD)
	t.Log("err=>", err)
	t.Log("ticker=>", ticker)
}
