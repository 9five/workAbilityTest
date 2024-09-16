package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"tokenize/config"
	_tradingHandler "tokenize/trading/delivery/handler"
	_tradingWssHandler "tokenize/trading/delivery/wssHandler"
	_tradingRepo "tokenize/trading/repository/binanceApi"
	_tradingUsecase "tokenize/trading/usecase"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	setupRouter(r)
	return r
}

func setupRouter(router *gin.Engine) {
	setupCORS(router)
	setupTrading(router)
	setupTradingWebSocket(router)
}

func setupCORS(router *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.AllowOrigin
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	router.Use(cors.New(corsConfig))

	err := router.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}
	router.Use(gin.Recovery())
}

func setupTrading(router *gin.Engine) {
	tradingRepo := _tradingRepo.NewBinanceApiTradingRepository(1)
	tradingUsecase := _tradingUsecase.NewTradingUsecase(tradingRepo)
	_tradingHandler.NewTradingHandler(router.Group("/task"), tradingUsecase)
}

func setupTradingWebSocket(router *gin.Engine) {
	tradingRepo := _tradingRepo.NewBinanceApiTradingRepository(1)
	tradingUsecase := _tradingUsecase.NewTradingUsecase(tradingRepo)
	_tradingWssHandler.NewWssTradingHandler(router.Group("/sockettask"), tradingUsecase)
}
