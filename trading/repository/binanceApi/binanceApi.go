package binanceapi

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"tokenize/domain"
)

type binanceApiTradingRepository struct {
	url string
}

func NewBinanceApiTradingRepository(urlSelection int) domain.TradingRepository {
	url := "https://api.binance.com"
	switch urlSelection {
	case 1:
		url = "https://api1.binance.com"
	case 2:
		url = "https://api2.binance.com"
	case 3:
		url = "https://api3.binance.com"
	case 4:
		url = "https://api4.binance.com"
	}
	return &binanceApiTradingRepository{url: url}
}

func (b *binanceApiTradingRepository) GetDepth(ctx context.Context, symbol string) (*domain.TradDepth, error) {
	var convert domain.TradDepth
	req, err := http.NewRequest("GET", b.url+"/api/v3/depth?symbol="+symbol, nil)
	if err != nil {
		return &convert, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &convert, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &convert, err
	}

	if res.StatusCode != 200 {
		return &convert, errors.New(res.Status + " : " + string(body))
	}

	err = json.Unmarshal(body, &convert)
	res.Body.Close()
	return &convert, err
}
