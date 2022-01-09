package arweave

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestArweave_Info(t *testing.T) {
	var tests = []struct {
		name       string
		response   Info
		json       string
		statusCode int
		err        error
	}{
		{
			name: "get info should return 200 ok",
			response: Info{
				Network:          "arweave.N.1",
				Version:          5,
				Release:          43,
				Height:           551511,
				Current:          "XIDpYbc3b5iuiqclSl_Hrx263Sd4zzmrNja1cvFlqNWUGuyymhhGZYI4WMsID1K3",
				Blocks:           97375,
				Peers:            64,
				QueueLength:      0,
				NodeStateLatency: 18,
			},
			json: `{
		 "network": "arweave.N.1",
		 "version": 5,
		 "release": 43,
		 "height": 551511,
		 "current": "XIDpYbc3b5iuiqclSl_Hrx263Sd4zzmrNja1cvFlqNWUGuyymhhGZYI4WMsID1K3",
		 "blocks": 97375,
		 "peers": 64,
		 "queue_length": 0,
		 "node_state_latency": 18
		}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "get info should return 400 Bad Request",
			response:   Info{},
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

			res, err := arweave.Info()

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.response, res)

			c.ExpectedCalls = []*mock.Call{}
		})
	}
}

func TestArweave_PeerList(t *testing.T) {

	var tests = []struct {
		name       string
		response   PeerList
		json       string
		statusCode int
		err        error
	}{
		{
			name: "get peer list should return 200 OK",
			response: PeerList{
				"localhost:1984",
				"127.0.0.1:1984",
			},
			json:       `["localhost:1984", "127.0.0.1:1984"]`,
			statusCode: http.StatusOK,
		},
		{
			name:       "get peer list should return 400 Bad Request",
			response:   PeerList{},
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

			res, err := arweave.PeerList()

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.response, res)

			c.ExpectedCalls = []*mock.Call{}
		})
	}
}
