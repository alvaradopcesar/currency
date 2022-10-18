package free_currency_api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/currency/internal/request"
)

// maxRequestTime is timeout for http request.
const maxRequestTime = 15000

type Repository interface {
	GetCurrencyLatest(ctx context.Context) (*Data, *QueryInfo, error)
}

type repository struct {
	req  request.Request
	conf Config
}

func NewRepository(req request.Request, conf Config) Repository {
	return &repository{
		req:  req,
		conf: conf,
	}
}

func (r repository) GetCurrencyLatest(ctx context.Context) (*Data, *QueryInfo, error) {
	query := map[string]string{
		"apikey": r.conf.FreeCurrencyApiKey,
	}

	url := fmt.Sprintf("/v1/latest")

	req, err := request.MakeHttpRequest(http.MethodGet, url, nil, query, nil)
	if err != nil {
		return nil, nil, err
	}

	timeInit := time.Now()

	resp, t, err := r.req.Do(ctx, req, maxRequestTime)
	if err != nil {
		return nil, nil, err
	}

	duration := time.Since(timeInit)

	var currency Data
	err = json.Unmarshal(resp, &currency)
	if err != nil {
		return nil, nil, err
	}

	info := &QueryInfo{
		Method:  req.Method,
		Address: req.URL.Path,
		Code:    t,
		Time:    duration.Seconds(),
	}

	return &currency, info, nil
}
