package response

import "fmt"

const (
	ErrCodeInternalServerError        = "err_internal"
	ErrDescriptionInternalServerError = "Internal Server Error"

	ErrCodeBadRequest            = "err_bad_request"
	ErrDescriptionBadRequestURL  = "The URL In Request contains errors"
	ErrDescriptionBadRequestBody = "The provided body contains errors"

	ErrCodeNotFound        = "err_not_found"
	ErrDescriptionNotFound = "The requested resource was not found"

	ErrCodeUnprocessableEntity        = "err_unprocessable_entity"
	ErrDescriptionUnprocessableEntity = "The request could not be processed"
	ErrReasonIsRequired               = "Reason is required"

	ErrCodeForbidden        = "err_forbidden"
	ErrDescriptionForbidden = "You do not have permission to access this resource"
)

var (
	StandardInternalServerError = Error{ErrCodeInternalServerError, ErrDescriptionInternalServerError}
	StandardBadBodyRequest      = Error{ErrCodeBadRequest, ErrDescriptionBadRequestBody}
	StandardNotFoundError       = Error{ErrCodeNotFound, ErrDescriptionNotFound}
	StandardUnprocessableEntity = Error{ErrCodeUnprocessableEntity, ErrDescriptionUnprocessableEntity}
	StandardForbidden           = Error{ErrCodeForbidden, ErrDescriptionForbidden}
)

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s:%s", e.Code, e.Description)
}
