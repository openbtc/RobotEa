package aex

import (
	"github.com/openbtc/RobotEa"
	"net/http"
	"testing"
)

var acx = New(http.DefaultClient, "", "", "")

func TestAcx_GetTicker(t *testing.T) {
	ticker, err := acx.GetTicker(goex.ETH_BTC)
	t.Log("err=>", err)
	t.Log("ticker=>", ticker)
}
