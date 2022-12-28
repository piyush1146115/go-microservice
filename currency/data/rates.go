package data

import (
	"encoding/xml"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

func (e *ExchangeRates) GetRates(base, dest string) (float64, error) {
	br, ok := e.rates[base]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", base)
	}

	dr, ok := e.rates[dest]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", dest)
	}

	return dr / br, nil
}

// MonitorRates checks the rates in the ECB API every interval and sends a message to
// the returned channel when there are changes
//
// Note: the ECB API only returns data once a day, this function only simulates the changes
// in rates for demonstration purpose
func (e *ExchangeRates) MonitorRates(interval time.Duration) chan struct{} {
	ret := make(chan struct{})

	go func() {
		ticker := time.NewTicker(interval)

		for {
			select {
			case <-ticker.C:
				// just add a random difference to the rate and return it
				// this simulates the fluctuations in currency rates

				for k, v := range e.rates {
					// change can be 10% of the original value
					change := (rand.Float64() / 10)
					// is this a positive or negative change
					direction := rand.Intn(1)

					if direction == 0 {
						// new rate will be 90% of the old rate
						change = 1 - change
					} else {
						// new rate will be 110% of the old
						change = 1 + change
					}

					// modify the rate
					e.rates[k] = v * change
				}

				// notify updates, this will block unless there is a listener on the other end
				ret <- struct{}{}
			}
		}
	}()

	return ret
}

func NewRates(l hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{log: l, rates: map[string]float64{}}

	err := er.getRates()

	return er, err
}

func (e *ExchangeRates) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status code 200, got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	md := &Cubes{}
	if err := xml.NewDecoder(resp.Body).Decode(&md); err != nil {
		return fmt.Errorf("failed to decode currency rates: %w", err)
	}

	for _, c := range md.CubeData {
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}
		e.rates[c.Currency] = r
	}
	e.rates["EUR"] = 1

	return nil
}

type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
