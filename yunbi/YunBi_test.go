package yunbi

import (
	"testing"
	"net/http"
	. "github.com/openbtc/RobotEa"
)

var (
	yb = New(http.DefaultClient, "", "")
)

func TestYunBi_GetTicker(t *testing.T) {
	t.Log(yb.GetTicker(BTS_CNY))
	t.Log(yb.GetTicker(SC_CNY))
	t.Log(yb.GetTicker(EOS_CNY))
}
