package service

import (
	"context"
	"github.com/currency/pkg/currency/repository"
	"github.com/currency/pkg/free_currency_api"
	"time"
)

type Service interface {
	GetAll(ctx context.Context) ([]GetAllResponse, error)
	GetByCurrency(ctx context.Context, f GetByCurrencyRequest, currency string) ([]GetByCurrencyResponse, error)
	Insert(ctx context.Context, req *repository.InsertCurrency) error
	InsertQuery(ctx context.Context, req *repository.InsertQuery) error
}

type service struct {
	repo repository.Repository
	serv free_currency_api.Service
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

type GetAllResponse struct {
	Code      string
	Value     float64
	CreatedAt time.Time
}

func (s service) GetAll(ctx context.Context) ([]GetAllResponse, error) {
	getAlls, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var getAllResponses []GetAllResponse
	for _, getAll := range getAlls {
		getAllResponses = append(getAllResponses, GetAllResponse{
			Code:      getAll.Code,
			Value:     getAll.Value,
			CreatedAt: getAll.CreatedAt,
		})

	}

	return getAllResponses, nil
}

type GetByCurrencyResponse struct {
	CustomerID int
	Code       string
	Value      float64
	CreatedAt  time.Time
}

type GetByCurrencyRequest struct {
	FInit time.Time `json:"finit"`
	FEnd  time.Time `json:"fend"`
}

func (s service) GetByCurrency(ctx context.Context, f GetByCurrencyRequest, currency string) ([]GetByCurrencyResponse, error) {
	if !f.FInit.IsZero() && f.FEnd.IsZero() {
		f.FInit = time.Now().Add(24 * time.Hour)
	}

	filter := repository.CurrencyFilter{
		Code:  currency,
		FInit: f.FInit,
		FEnd:  f.FEnd,
	}

	getByIDs, err := s.repo.GetByID(ctx, filter)
	if err != nil {
		return nil, err
	}
	var getByCurrencyResponses []GetByCurrencyResponse
	for _, getByID := range getByIDs {
		getByCurrencyResponses = append(getByCurrencyResponses, GetByCurrencyResponse{
			CustomerID: getByID.CustomerID,
			Code:       getByID.Code,
			Value:      getByID.Value,
			CreatedAt:  getByID.CreatedAt,
		})
	}

	return getByCurrencyResponses, nil
}

func (s service) Insert(ctx context.Context, req *repository.InsertCurrency) error {
	err := s.repo.Insert(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s service) InsertQuery(ctx context.Context, req *repository.InsertQuery) error {
	err := s.repo.InsertQuery(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
