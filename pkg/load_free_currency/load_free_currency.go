package load_free_currency

import (
	"context"
	"github.com/currency/pkg/currency/repository"
	"github.com/currency/pkg/currency/service"
	"github.com/currency/pkg/free_currency_api"
	"time"

	"github.com/currency/internal/logger"
	"github.com/go-co-op/gocron"
)

type Provider struct {
	freeCurrencyTime    int
	servFreeCurrencyAPI free_currency_api.Service
	servCurrency        service.Service
	log                 logger.Logger
}

func NewProvider(freeCurrencyTime int, servFree free_currency_api.Service, servCurrency service.Service, log logger.Logger) *Provider {
	return &Provider{
		freeCurrencyTime:    freeCurrencyTime,
		servFreeCurrencyAPI: servFree,
		servCurrency:        servCurrency,
		log:                 log,
	}
}

func (c Provider) Load() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(c.freeCurrencyTime).Minute().Do(c.saveDB)
	s.StartBlocking()
}

func (c Provider) saveDB() {
	ctx := context.Background()
	data, info, err := c.servFreeCurrencyAPI.GetCurrencyLatest(ctx)
	if err != nil {
		c.log.Error(err)
		return
	}

	for key, value := range data.Data {
		err = c.servCurrency.Insert(ctx, &repository.InsertCurrency{
			Code:  key,
			Value: value,
		})
		if err != nil {
			c.log.Error(err)
		}
		if err != nil {
			continue
		}
	}

	err = c.servCurrency.InsertQuery(ctx, &repository.InsertQuery{
		Method:  info.Method,
		Address: info.Address,
		Code:    info.Code,
		Time:    info.Time,
	})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info("Loading data ok")
}
