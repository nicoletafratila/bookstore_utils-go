package rest_errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "message: this is the message - status: 500 - error: internal_server_error - causes: [ [database error] ]", err.Error())

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "message: this is the message - status: 400 - error: bad_request - causes: [ [] ]", err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "message: this is the message - status: 404 - error: not_found - causes: [ [] ]", err.Error())
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "message: this is the message - status: 401 - error: unauthorized - causes: [ [] ]", err.Error())
}

func TestNewRestError(t *testing.T) {
	err := NewRestError("this is the message", 600, "error type", nil)
	assert.NotNil(t, err)
	assert.EqualValues(t, 600, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "message: this is the message - status: 600 - error: error type - causes: [ [] ]", err.Error())
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 0, len(err.Causes()))
}
