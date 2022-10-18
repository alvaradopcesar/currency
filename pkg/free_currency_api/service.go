package free_currency_api

import "context"

type Service interface {
	GetCurrencyLatest(ctx context.Context) (*Data, *QueryInfo, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) GetCurrencyLatest(ctx context.Context) (*Data, *QueryInfo, error) {
	res, info, err := s.repo.GetCurrencyLatest(ctx)
	if err != nil {
		return nil, nil, err
	}

	return res, info, nil
}
