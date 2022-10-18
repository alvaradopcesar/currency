package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/currency/internal/config"
)

const versionHeader = "Api-Version"

type BaseResponse struct {
	Result interface{} `json:"result,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

func newBaseResponseWithData(data interface{}) BaseResponse {
	return BaseResponse{
		Result: data,
	}
}

func newBaseResponseWithError(err interface{}) BaseResponse {
	return BaseResponse{
		Error: err,
	}
}

func RespondWithData(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(versionHeader, config.GetVersion())
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(newBaseResponseWithData(data))
}

func RespondWithError(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(versionHeader, config.GetVersion())
	w.WriteHeader(statusCodeFromError(err))
	return json.NewEncoder(w).Encode(newBaseResponseWithError(viewModelFromError(err)))
}

func statusCodeFromError(err error) int {
	vErr := &Error{}
	if errors.As(err, vErr) {
		switch vErr.Code {
		case ErrCodeBadRequest:
			return http.StatusBadRequest
		case ErrCodeNotFound:
			return http.StatusNotFound
		case ErrCodeUnprocessableEntity:
			return http.StatusUnprocessableEntity
		case ErrCodeForbidden:
			return http.StatusForbidden
		default:
			return http.StatusInternalServerError
		}
	}

	return http.StatusInternalServerError
}

func viewModelFromError(err error) Error {
	vErr := Error{}
	if errors.As(err, &vErr) {
		return vErr
	}
	return StandardInternalServerError
}
