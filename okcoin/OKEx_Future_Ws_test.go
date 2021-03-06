package okcoin

import (
	"testing"
	"net/http"
	"github.com/openbtc/RobotEa"
	"log"
	"time"
)

var okexFuture = NewOKEx(http.DefaultClient, "", "")

func TestOKEx_GetDepthWithWs(t *testing.T) {
	okexFuture.GetDepthWithWs(goex.BTC_USD, goex.QUARTER_CONTRACT, func(depth *goex.Depth) {
		log.Print(depth)
	})
	time.Sleep(1 * time.Minute)
	okexFuture.ws.CloseWs()
}
