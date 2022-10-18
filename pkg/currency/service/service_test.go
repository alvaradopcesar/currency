package service

import (
	"context"
	"errors"
	"github.com/currency/pkg/currency/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

type repositoryMock struct{}

func (r repositoryMock) GetAll(ctx context.Context) ([]GetAllResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (r repositoryMock) GetByID(ctx context.Context, f GetByCurrencyRequest) ([]GetByCurrencyResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (r repositoryMock) Insert(ctx context.Context, req *repository.InsertCurrency) error {
	// TODO implement me
	panic("implement me")
}

func (r repositoryMock) InsertQuery(ctx context.Context, req *repository.InsertQuery) error {
	// TODO implement me
	panic("implement me")
}

type GetAllCurrencyMock struct {
	repositoryMock
	GetAllCurrencyRes []GetAllResponse
	GetAllCurrencyErr error
}

func (m GetAllCurrencyMock) GetAll(ctx context.Context) ([]GetAllResponse, error) {
	return m.GetAllCurrencyRes, m.GetAllCurrencyErr
}

func TestService_GetAll(t *testing.T) {
	allCurrency := []GetAllResponse{
		{
			Code:  "USD",
			Value: 1,
		},
		{
			Code:  "TRY",
			Value: 17.958154,
		},
	}
	errorMock := errors.New("not found")
	tests := []struct {
		name              string
		getAllCurrencyRes []GetAllResponse
		getAllCurrencyErr error
		GetAllRes         []GetAllResponse
		GetAllErr         error
	}{
		{
			name:              "success",
			getAllCurrencyRes: allCurrency,
			getAllCurrencyErr: nil,
			GetAllRes:         allCurrency,
			GetAllErr:         nil,
		},
		{
			name:              "fail",
			getAllCurrencyRes: nil,
			getAllCurrencyErr: errorMock,
			GetAllRes:         nil,
			GetAllErr:         errorMock,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := GetAllCurrencyMock{
				GetAllCurrencyRes: test.getAllCurrencyRes,
				GetAllCurrencyErr: test.getAllCurrencyErr,
			}
			resp, err := s.GetAll(ctx)
			assert.Equal(t, test.GetAllRes, resp)
			assert.Equal(t, test.GetAllErr, err)
		})
	}
}

type getByCurrencyMock struct {
	repositoryMock
	GetByIDCurrencyRes []GetByCurrencyResponse
	GetByIDCurrencyErr error
}

func (m getByCurrencyMock) GetByCurrency(ctx context.Context, f GetByCurrencyRequest, currency string) ([]GetByCurrencyResponse, error) {
	return m.GetByIDCurrencyRes, m.GetByIDCurrencyErr
}

func TestService_GetByCurrency(t *testing.T) {
	currency := []GetByCurrencyResponse{
		{
			CustomerID: 1,
			Code:       "USD",
			Value:      1.2,
			CreatedAt:  time.Now(),
		},
	}
	filter := GetByCurrencyRequest{
		FInit: time.Now(),
		FEnd:  time.Now(),
	}
	f := repository.CurrencyFilter{
		Code:  "USD",
		FInit: time.Now(),
		FEnd:  time.Now(),
	}
	errorMock := errors.New("not found")
	tests := []struct {
		name                string
		filter              GetByCurrencyRequest
		currencyID          string
		getByIDReq          repository.CurrencyFilter
		getByIDCurrencyResp []GetByCurrencyResponse
		getByIDCurrencyErr  error
		getByIDRes          []GetByCurrencyResponse
		getByIDErr          error
	}{
		{
			name:                "success",
			filter:              filter,
			currencyID:          "USD",
			getByIDReq:          f,
			getByIDCurrencyResp: currency,
			getByIDCurrencyErr:  nil,
			getByIDRes:          currency,
			getByIDErr:          nil,
		},
		{
			name:                "fail filter",
			filter:              GetByCurrencyRequest{},
			currencyID:          "",
			getByIDReq:          repository.CurrencyFilter{},
			getByIDCurrencyResp: nil,
			getByIDCurrencyErr:  errorMock,
			getByIDRes:          nil,
			getByIDErr:          errorMock,
		},
		{
			name:                "fail",
			filter:              GetByCurrencyRequest{},
			currencyID:          "USD",
			getByIDReq:          repository.CurrencyFilter{},
			getByIDCurrencyResp: nil,
			getByIDCurrencyErr:  errorMock,
			getByIDRes:          nil,
			getByIDErr:          errorMock,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := getByCurrencyMock{
				GetByIDCurrencyRes: test.getByIDCurrencyResp,
				GetByIDCurrencyErr: test.getByIDCurrencyErr,
			}
			resp, err := s.GetByCurrency(ctx, test.filter, test.currencyID)
			assert.Equal(t, test.getByIDRes, resp)
			assert.Equal(t, test.getByIDErr, err)
		})
	}
}

type insertCurrency struct {
	repositoryMock
	InsertCurrencyErr error
}

func (m insertCurrency) Insert(ctx context.Context, req *repository.InsertCurrency) error {
	return m.InsertCurrencyErr
}

func TestService_Insert(t *testing.T) {
	errorMock := errors.New("not found")
	tests := []struct {
		name              string
		insertCurrencyReq *repository.InsertCurrency
		insertCurrencyErr error
		insertErr         error
	}{
		{
			name: "success",
			insertCurrencyReq: &repository.InsertCurrency{
				Code:  "USD",
				Value: 1,
			},
			insertCurrencyErr: nil,
			insertErr:         nil,
		},
		{
			name:              "fail",
			insertCurrencyReq: &repository.InsertCurrency{},
			insertCurrencyErr: errorMock,
			insertErr:         errorMock,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := insertCurrency{
				InsertCurrencyErr: test.insertCurrencyErr,
			}
			err := s.Insert(ctx, test.insertCurrencyReq)
			assert.Equal(t, test.insertErr, err)
		})
	}
}

type insertQueryMock struct {
	repositoryMock
	insertQueryErr error
}

func (m insertQueryMock) InsertQuery(ctx context.Context, req *repository.InsertQuery) error {
	return m.insertQueryErr
}

func TestService_InsertQuery(t *testing.T) {
	errorMock := errors.New("not found")
	tests := []struct {
		name           string
		insertQueryReq *repository.InsertQuery
		insertQueryErr error
		insertErr      error
	}{
		{
			name: "success",
			insertQueryReq: &repository.InsertQuery{
				Method:  "GET",
				Address: "/test",
				Code:    200,
				Time:    0.12,
			},
			insertQueryErr: nil,
			insertErr:      nil,
		},
		{
			name:           "fail",
			insertQueryReq: &repository.InsertQuery{},
			insertQueryErr: errorMock,
			insertErr:      errorMock,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := insertQueryMock{
				insertQueryErr: test.insertQueryErr,
			}
			err := s.InsertQuery(ctx, test.insertQueryReq)
			assert.Equal(t, test.insertErr, err)
		})
	}
}
