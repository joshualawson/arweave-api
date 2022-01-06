package arweave

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"math/big"
	"net/http"
	"testing"
)

func TestArweave_WalletBalance(t *testing.T) {
	var tests = []struct {
		name       string
		address    string
		response   *big.Int
		json       string
		statusCode int
		err        error
	}{
		{
			name:       "get wallet balance should return 200 ok",
			address:    "*address*",
			response:   big.NewInt(200),
			json:       `200`,
			statusCode: http.StatusOK,
		},
		{
			name:       "get wallet balance should return 400 Bad Request",
			address:    "*address*",
			response:   nil,
			json:       ``,
			statusCode: http.StatusBadRequest,
			err:        ErrorNotOk(http.StatusBadRequest),
		},
	}
	c := &MockClient{}

	arweave := New(
		WithClient(c),
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.On("Do", mock.Anything).Return(&http.Response{StatusCode: tt.statusCode, Body: ioutil.NopCloser(bytes.NewReader([]byte(tt.json)))}, nil)

			res, err := arweave.WalletBalance(tt.address)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.response, res)

			c.ExpectedCalls = []*mock.Call{}
		})
	}
}

func TestArweave_WalletLastTransactionID(t *testing.T) {
	var tests = []struct {
		name       string
		address    string
		response   string
		json       string
		statusCode int
		err        error
	}{
		{
			name:       "get wallet balance should return 200 ok",
			address:    "*address*",
			response:   "*last_transaction_id*",
			json:       `*last_transaction_id*`,
			statusCode: http.StatusOK,
		},
		{
			name:       "get wallet balance should return 400 Bad Request",
			address:    "*address*",
			response:   "",
			json:       ``,
			statusCode: http.StatusBadRequest,
			err:        ErrorNotOk(http.StatusBadRequest),
		},
	}
	c := &MockClient{}

	arweave := New(
		WithClient(c),
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.On("Do", mock.Anything).Return(&http.Response{StatusCode: tt.statusCode, Body: ioutil.NopCloser(bytes.NewReader([]byte(tt.json)))}, nil)

			res, err := arweave.WalletLastTransactionID(tt.address)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.response, res)

			c.ExpectedCalls = []*mock.Call{}
		})
	}
}
