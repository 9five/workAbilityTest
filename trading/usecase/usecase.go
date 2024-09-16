package usecase

import (
	"context"
	"strconv"
	"tokenize/domain"
)

type tradingUsecase struct {
	tradingRepo domain.TradingRepository
}

func NewTradingUsecase(tradingRepo domain.TradingRepository) domain.TradingUsecase {
	return &tradingUsecase{
		tradingRepo: tradingRepo,
	}
}

func (t *tradingUsecase) GetOrderBook(ctx context.Context, symbol string) (*domain.OrderBook, error) {
	var orderBook domain.OrderBook
	depth, err := t.tradingRepo.GetDepth(ctx, symbol)
	if err != nil {
		return &orderBook, err
	}

	for _, v := range depth.Bids {
		price, err := strconv.ParseFloat(string(v[0]), 64)
		if err != nil {
			return &orderBook, err
		}
		qty, err := strconv.ParseFloat(string(v[1]), 64)
		if err != nil {
			return &orderBook, err
		}

		if len(orderBook.Bids) == 0 || price < orderBook.Bids[len(orderBook.Bids)-1].Price {
			orderBook.Bids = append(orderBook.Bids, domain.Order{Qty: qty, Price: price, Sum: qty * price})
		} else if price > orderBook.Bids[len(orderBook.Bids)-1].Price {
			orderBook.Bids = append([]domain.Order{{Qty: qty, Price: price, Sum: qty * price}}, orderBook.Bids...)
		}

	}

	for _, v := range depth.Asks {
		price, err := strconv.ParseFloat(string(v[0]), 64)
		if err != nil {
			return &orderBook, err
		}
		qty, err := strconv.ParseFloat(string(v[1]), 64)
		if err != nil {
			return &orderBook, err
		}

		if len(orderBook.Asks) == 0 || price > orderBook.Asks[len(orderBook.Asks)-1].Price {
			orderBook.Asks = append(orderBook.Asks, domain.Order{Qty: qty, Price: price, Sum: qty * price})
		} else if price < orderBook.Asks[len(orderBook.Asks)-1].Price {
			orderBook.Asks = append([]domain.Order{{Qty: qty, Price: price, Sum: qty * price}}, orderBook.Asks...)
		}
	}
	return &orderBook, nil
}

func (t *tradingUsecase) GetOrderBookForWss(ctx context.Context, symbol string) (*domain.OrderBook, error) {
	var orderBook domain.OrderBook
	depth, err := t.tradingRepo.GetDepth(ctx, symbol)
	if err != nil {
		return &orderBook, err
	}

	for _, v := range depth.Bids {
		price, err := strconv.ParseFloat(string(v[0]), 64)
		if err != nil {
			return &orderBook, err
		}
		qty, err := strconv.ParseFloat(string(v[1]), 64)
		if err != nil {
			return &orderBook, err
		}
		sum := qty * price

		if orderBook.BidsSum+sum >= 5 {
			break
		}

		orderBook.BidsSum += sum
		if len(orderBook.Bids) == 0 || price < orderBook.Bids[len(orderBook.Bids)-1].Price {
			orderBook.Bids = append(orderBook.Bids, domain.Order{Qty: qty, Price: price, Sum: sum})
		} else if price > orderBook.Bids[len(orderBook.Bids)-1].Price {
			orderBook.Bids = append([]domain.Order{{Qty: qty, Price: price, Sum: sum}}, orderBook.Bids...)
		}
	}

	for _, v := range depth.Asks {
		price, err := strconv.ParseFloat(string(v[0]), 64)
		if err != nil {
			return &orderBook, err
		}
		qty, err := strconv.ParseFloat(string(v[1]), 64)
		if err != nil {
			return &orderBook, err
		}
		sum := qty * price

		if orderBook.AsksSum+sum >= 150 {
			break
		}

		orderBook.AsksSum += sum
		if len(orderBook.Asks) == 0 || price > orderBook.Asks[len(orderBook.Asks)-1].Price {
			orderBook.Asks = append(orderBook.Asks, domain.Order{Qty: qty, Price: price, Sum: sum})
		} else if price < orderBook.Asks[len(orderBook.Asks)-1].Price {
			orderBook.Asks = append([]domain.Order{{Qty: qty, Price: price, Sum: sum}}, orderBook.Asks...)
		}
	}
	return &orderBook, nil
}
