package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type getAllRepositoryMock struct {
	GetAllCurrencyRes []Currency
	GetAllCurrencyErr error
}

func (p getAllRepositoryMock) GetAll(ctx context.Context) ([]Currency, error) {
	return p.GetAllCurrencyRes, nil
}

func (p getAllRepositoryMock) GetByID(ctx context.Context, f CurrencyFilter) ([]Currency, error) {
	return p.GetAllCurrencyRes, nil
}

func (p getAllRepositoryMock) Insert(ctx context.Context, req *InsertCurrency) error {
	return p.GetAllCurrencyErr
}

func (p getAllRepositoryMock) InsertQuery(ctx context.Context, req *InsertQuery) error {
	return p.GetAllCurrencyErr
}

func TestPostgres_GetAll(t *testing.T) {
	currency := []Currency{
		{
			Code:      "USD",
			Value:     1,
			CreatedAt: time.Now(),
		},
	}

	repo := getAllRepositoryMock{
		GetAllCurrencyRes: currency,
		GetAllCurrencyErr: nil,
	}
	getallResponse, err := repo.GetAll(context.Background())
	require.Equal(t, err, nil)
	require.Equal(t, currency, getallResponse)
}

func TestPostgres_GetByID(t *testing.T) {
	currency := []Currency{
		{
			CustomerID: 1,
			Code:       "USD",
			Value:      1,
			CreatedAt:  time.Now(),
		},
	}
	f := CurrencyFilter{
		Code:  "USD",
		FInit: time.Now(),
		FEnd:  time.Now(),
	}
	repo := getAllRepositoryMock{
		GetAllCurrencyRes: currency,
		GetAllCurrencyErr: nil,
	}
	getByIDResponse, err := repo.GetByID(context.Background(), f)
	require.Equal(t, err, nil)
	require.Equal(t, currency, getByIDResponse)
}

func TestPostgres_Insert(t *testing.T) {
	currency := &InsertCurrency{
		Code:  "USD",
		Value: 1,
	}
	repo := getAllRepositoryMock{
		GetAllCurrencyErr: nil,
	}
	err := repo.Insert(context.Background(), currency)
	require.Equal(t, err, nil)
}

func TestPostgres_InsertQuery(t *testing.T) {
	query := &InsertQuery{
		Method:  "GET",
		Address: "/currency/usd",
		Code:    200,
		Time:    0.9312,
	}
	repo := getAllRepositoryMock{
		GetAllCurrencyErr: nil,
	}
	err := repo.InsertQuery(context.Background(), query)
	require.Equal(t, err, nil)
}
