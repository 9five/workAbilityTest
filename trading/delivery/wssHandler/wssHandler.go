package wsshandler

import (
	"net/http"
	"time"
	"tokenize/domain"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WssTradingHandler struct {
	tradingUsecase domain.TradingUsecase
}

func NewWssTradingHandler(router *gin.RouterGroup, tradingUsecase domain.TradingUsecase) {
	handler := &WssTradingHandler{
		tradingUsecase: tradingUsecase,
	}

	router.GET("", handler.WssOrderBookGen)
}

func (t *WssTradingHandler) WssOrderBookGen(ctx *gin.Context) {
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer ws.Close()

	for {
		orderBook, err := t.tradingUsecase.GetOrderBookForWss(ctx, "ETHBTC")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = ws.WriteJSON(orderBook)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		time.Sleep(time.Second * 30)
	}
}
