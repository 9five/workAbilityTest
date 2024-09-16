package domain

import "context"

type TradDepth struct {
	LastUpdateId int64      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type OrderBook struct {
	Bids    []Order `json:"bids"`
	Asks    []Order `json:"asks"`
	BidsSum float64 `json:"bidsSum"`
	AsksSum float64 `json:"asksSum"`
}

type Order struct {
	Qty   float64 `json:"qty"`
	Price float64 `json:"price"`
	Sum   float64
}

type TradingRepository interface {
	GetDepth(ctx context.Context, symbol string) (*TradDepth, error)
}

type TradingUsecase interface {
	GetOrderBook(ctx context.Context, symbol string) (*OrderBook, error)
	GetOrderBookForWss(ctx context.Context, symbol string) (*OrderBook, error)
}
