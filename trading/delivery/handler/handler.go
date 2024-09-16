package handler

import (
	"net/http"
	"tokenize/domain"

	"github.com/gin-gonic/gin"
)

type TradingHandler struct {
	tradingUsecase domain.TradingUsecase
}

func NewTradingHandler(router *gin.RouterGroup, tradingUsecase domain.TradingUsecase) {
	handler := &TradingHandler{
		tradingUsecase: tradingUsecase,
	}

	router.GET("", handler.OrderBookGen)
}

func (t *TradingHandler) OrderBookGen(ctx *gin.Context) {
	orderBook, err := t.tradingUsecase.GetOrderBook(ctx, "ETHBTC")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, orderBook)
}
