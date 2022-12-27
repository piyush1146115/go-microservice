package server

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"github.com/piyush1146115/go-microservice/currency/data"
	protos "github.com/piyush1146115/go-microservice/currency/protos/currency"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Currency struct {
	rates *data.ExchangeRates
	log   hclog.Logger
	protos.UnimplementedCurrencyServer
}

// NewCurrency creates a new Currency server
func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{r, l, protos.UnimplementedCurrencyServer{}}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())

	rate, err := c.rates.GetRates(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Rate: rate}, nil
}

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {
	return src.Send(&protos.RateResponse{Rate: 12.1})
}
