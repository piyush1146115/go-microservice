package server

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"github.com/piyush1146115/go-microservice/currency/data"
	protos "github.com/piyush1146115/go-microservice/currency/protos/currency"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	c := &Currency{r, l, make(map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest), protos.UnimplementedCurrencyServer{}}

	go c.handleUpdates()

	return c
}

func (c *Currency) handleUpdates() {
	ru := c.rates.MonitorRates(5 * time.Second)

	for range ru {
		c.log.Info("Got updated rates")

		// loop over subscribed clients
		for k, v := range c.subscriptions {

			// loop over subscribed rates
			for _, rr := range v {
				r, err := c.rates.GetRates(rr.GetBase().String(), rr.GetDestination().String())
				if err != nil {
					c.log.Error("Unable to get updated rates", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
				}

				err = k.Send(&protos.StreamingRateResponse{
					Message: &protos.StreamingRateResponse_RateResponse{
						RateResponse: &protos.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: r},
					},
				})

				if err != nil {
					c.log.Error("Unable to send update rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
				}
			}
		}
	}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())

	if rr.Base == rr.Destination {
		err := status.Newf(
			codes.InvalidArgument,
			"Base currency %s can not be same as the destination currency %s",
			rr.Base.String(),
			rr.Destination.String(),
		)

		err, wde := err.WithDetails(rr)
		if wde != nil {
			return nil, wde
		}

		return nil, err.Err()
	}

	rate, err := c.rates.GetRates(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: rate}, nil
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

		// check if already in the subscribe list and return a gRPC error
		for _, r := range rrs {
			// if we already have subscribed to thhis currency, return an error
			if r.Base == rr.Base && r.Destination == rr.Destination {
				c.log.Error("Subscription already active", "base", rr.Base.String(), "dest", rr.Destination.String())

				grpcError := status.New(codes.InvalidArgument, "Subscription already active for rate")

				grpcError, err = grpcError.WithDetails(rr)
				if err != nil {
					c.log.Error("Unable to add metadata to error message", "error", err)
					continue
				}

				// Can't return error as that will terminate the connection, instead must send an error whichh
				// can be handled by the client Recv stream
				rrs := &protos.StreamingRateResponse_Error{Error: grpcError.Proto()}
				src.Send(&protos.StreamingRateResponse{Message: rrs})
			}
		}

		rrs = append(rrs, rr)
		c.subscriptions[src] = rrs
	}

	return nil
}
