package server

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"github.com/piyush1146115/go-microservice/currency/data"
	protos "github.com/piyush1146115/go-microservice/currency/protos/currency"
	"io"
	"time"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Currency struct {
	rates         *data.ExchangeRates
	log           hclog.Logger
	subscriptions map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest
	protos.UnimplementedCurrencyServer
}

// NewCurrency creates a new Currency server
func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	go func() {
		ru := r.MonitorRates(5 * time.Second)

		for range ru {
			l.Info("Got updated rates")
		}
	}()

	return &Currency{r, l, make(map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest), protos.UnimplementedCurrencyServer{}}
}

func (c *Currency) handleUpdates() {
	ru := c.rates.MonitorRates(5 * time.Second)

	for range ru {
		c.log.Info("Got updated rates")

		// loop over subscribed clients
		for k, v := range c.subscriptions {
			
		}
	}
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

// SubscribeRates implements the gRPC bidirectional streaming method for the server
func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {
	// handle client messages
	for {
		rr, err := src.Recv() // Recv is a blocking method which returns on client data
		// io.EOF signals that the client has closed the connection
		if err == io.EOF {
			c.log.Info("Client has closed connection")
			break
		}

		// any other error means the transport between the server and client is unavailable
		if err != nil {
			c.log.Error("Unable to read from client", "error", err)
			return err
		}

		c.log.Info("Handle client request", "request_base", rr.GetBase(), "request_dest", rr.GetDestination())

		rrs, ok := c.subscriptions[src]
		if !ok {
			rrs = []*protos.RateRequest{}
		}

		rrs = append(rrs, rr)
		c.subscriptions[src] = rrs
	}

	return nil
}
