package binanceapi_test

import (
	"context"
	"testing"

	"tokenize/domain"
	repo "tokenize/trading/repository/binanceApi"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TradingBinanceApiSuite struct {
	suite.Suite
	repo domain.TradingRepository
}

func TestStart(t *testing.T) {
	suite.Run(t, &TradingBinanceApiSuite{})
}

func (s *TradingBinanceApiSuite) SetupTest() {
	s.repo = repo.NewBinanceApiTradingRepository(1)
}

func (s *TradingBinanceApiSuite) TestGetDepth_Success() {
	_, err := s.repo.GetDepth(context.TODO(), "ETHBTC")
	assert.NoError(s.Suite.T(), err)
}
