package handler

import (
	"github.com/currency/pkg/currency/service"
	"net/http"
	"time"

	"github.com/currency/internal/response"
	"github.com/go-chi/chi"
)

const dateFormat = "2006-01-02"

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return Handler{
		service: service,
	}
}

func (h *Handler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp, err := h.service.GetAll(ctx)
	if err != nil {
		_ = response.RespondWithError(w, response.Error{
			Code:        response.ErrDescriptionInternalServerError,
			Description: err.Error(),
		})
		return
	}

	_ = response.RespondWithData(w, http.StatusOK, resp)
}

func (h *Handler) GetByCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currency := chi.URLParam(r, "currency")
	dateStart := r.URL.Query().Get("finit")
	end := r.URL.Query().Get("fend")

	var (
		req service.GetByCurrencyRequest
		err error
	)

	if dateStart != "" {
		req.FInit, err = time.Parse(dateFormat, dateStart)
		if err != nil {
			_ = response.RespondWithError(w, response.Error{
				Code:        response.ErrCodeBadRequest,
				Description: err.Error(),
			})
			return
		}
		req.FInit = req.FInit.Add(24 * time.Hour)
	}

	if end != "" {
		req.FEnd, err = time.Parse(dateFormat, end)
		if err != nil {
			_ = response.RespondWithError(w, response.Error{
				Code:        response.ErrCodeBadRequest,
				Description: err.Error(),
			})
			return
		}
		req.FEnd = req.FEnd.Add(24 * time.Hour)
	}

	resp, err := h.service.GetByCurrency(ctx, req, currency)
	if err != nil {
		_ = response.RespondWithError(w, response.Error{
			Code:        response.ErrDescriptionInternalServerError,
			Description: err.Error(),
		})
		return
	}

	_ = response.RespondWithData(w, http.StatusOK, resp)
}
