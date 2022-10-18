package handler

import (
	"bytes"
	"context"
	"errors"
	"github.com/currency/pkg/currency/repository"
	"github.com/currency/pkg/currency/service"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type serviceMock struct{}

func (s serviceMock) GetAll(ctx context.Context) ([]service.GetAllResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s serviceMock) GetByCurrency(ctx context.Context, f service.GetByCurrencyRequest, currency string) ([]service.GetByCurrencyResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s serviceMock) Insert(ctx context.Context, req *repository.InsertCurrency) error {
	// TODO implement me
	panic("implement me")
}

func (s serviceMock) InsertQuery(ctx context.Context, req *repository.InsertQuery) error {
	// TODO implement me
	panic("implement me")
}

type getAllHandlerMock struct {
	serviceMock
	GetAllCurrencyRes []service.GetAllResponse
	GetAllCurrencyErr error
}

func (m getAllHandlerMock) GetAll(ctx context.Context) ([]service.GetAllResponse, error) {
	return m.GetAllCurrencyRes, m.GetAllCurrencyErr
}

func TestHandler_GetAllHandler(t *testing.T) {
	errNotFound := errors.New("not found")
	tests := []struct {
		name       string
		statusCode int
		resp       []service.GetAllResponse
		err        error
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			resp:       []service.GetAllResponse{},
			err:        nil,
		},
		{
			name:       "failure",
			statusCode: http.StatusInternalServerError,
			resp:       nil,
			err:        errNotFound,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := getAllHandlerMock{
				GetAllCurrencyRes: test.resp,
				GetAllCurrencyErr: test.err,
			}

			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodGet, "/currency", nil)
			if err != nil {
				require.NoError(t, err)
			}

			h := NewHandler(m)

			mux := chi.NewMux()
			mux.Get("/currency", h.GetAllHandler)
			mux.ServeHTTP(w, r)

			statusCode := w.Result().StatusCode
			assert.Equal(t, test.statusCode, statusCode)
		})
	}
}

type getByCurrencyHandlerMock struct {
	serviceMock
	getByCurrencyRes []service.GetByCurrencyResponse
	getByCurrencyErr error
}

func (m getByCurrencyHandlerMock) GetByCurrency(ctx context.Context, f service.GetByCurrencyRequest, currency string) ([]service.GetByCurrencyResponse, error) {
	return m.getByCurrencyRes, m.getByCurrencyErr
}

func TestHandler_GetByCurrencyHandler(t *testing.T) {
	errNotFound := errors.New("not found")
	q := "?finit=2022-08-14&fend=2022-08-17"
	currency := "USD"
	tests := []struct {
		name       string
		statusCode int
		body       []byte
		f          repository.CurrencyFilter
		resp       []service.GetByCurrencyResponse
		err        error
	}{
		{
			name:       "success",
			body:       []byte(`{"currency":"USD"}`),
			statusCode: http.StatusOK,
			resp:       []service.GetByCurrencyResponse{},
			err:        nil,
		},
		{
			name:       "failure",
			statusCode: http.StatusInternalServerError,
			body:       nil,
			resp:       nil,
			err:        errNotFound,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := getByCurrencyHandlerMock{
				getByCurrencyRes: test.resp,
				getByCurrencyErr: test.err,
			}

			b := bytes.NewBuffer(test.body)

			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodGet, "/currency/"+currency+q, b)
			if err != nil {
				require.NoError(t, err)
			}

			init := r.URL.Query().Get("finit")
			init = "2022-02-03"

			if init != "" {
				test.f.FInit, err = time.Parse(dateFormat, init)
				if err != nil {
					t.Fatal(err)
				}
				test.f.FInit = test.f.FInit.Add(24 * time.Hour)
			}

			end := r.URL.Query().Get("fend")
			init = "2022-02-03"

			if init != "" {
				test.f.FEnd, err = time.Parse(dateFormat, end)
				if err != nil {
					t.Fatal(err)
				}
				test.f.FEnd = test.f.FEnd.Add(24 * time.Hour)
			}

			h := NewHandler(m)

			mux := chi.NewMux()
			mux.Get("/currency/{currency}", h.GetByCurrencyHandler)
			mux.ServeHTTP(w, r)

			statusCode := w.Result().StatusCode
			assert.Equal(t, test.statusCode, statusCode)
		})
	}
}
