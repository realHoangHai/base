package errors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(status int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: status,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewCustomError(status int, root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(status, root, msg, root.Error(), key)
	}
	return NewErrorResponse(status, errors.New(msg), msg, msg, key)
}

func ErrInvalidRequest(root error) *AppError {
	return NewErrorResponse(http.StatusBadRequest, root, "invalid request", root.Error(), "ErrInvalidRequest")
}

func ErrInternal(root error) *AppError {
	return NewErrorResponse(http.StatusInternalServerError, root, "something went wrong in server", root.Error(), "ErrInternal")
}

func ErrDB(root error) *AppError {
	return NewErrorResponse(http.StatusInternalServerError, root, "something went wrong with DB", root.Error(), "DB_ERROR")
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("Err%sAlreadyExists", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		http.StatusBadRequest,
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", strings.Title(entity)),
	)
}

func ErrUnauthorized(root error) *AppError {
	return NewCustomError(http.StatusUnauthorized, root, "unauthorized", "ErrUnauthorized")
}

func ErrTooManyRequests(root error) *AppError {
	return NewCustomError(http.StatusTooManyRequests, root, "too many requests", "TooManyRequests")
}

func ErrMethodNotAllowed(root error) *AppError {
	return NewCustomError(http.StatusMethodNotAllowed, root, "method not allowed", "MethodNotAllowed")
}
