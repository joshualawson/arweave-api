package arweave

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type InfoResponse struct {
	Network          string `json:"network"`
	Version          int    `json:"version"`
	Release          int    `json:"release"`
	Height           int    `json:"height"`
	Current          string `json:"current"`
	Blocks           int    `json:"blocks"`
	Peers            int    `json:"peers"`
	QueueLength      int    `json:"queue_length"`
	NodeStateLatency int    `json:"node_state_latency"`
}

// Info Get the current network information including height, current block, and other properties.
func (a *Arweave) Info() (InfoResponse, error) {
	req, err := http.NewRequest("GET", a.fqdn()+"/info", nil)
	if err != nil {
		return InfoResponse{}, err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return InfoResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		return InfoResponse{}, ErrorNotOk(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return InfoResponse{}, err
	}

	var ir InfoResponse
	json.Unmarshal(body, &ir)

	return ir, nil
}

type PeerListResponse []string

// PeerList Get the list of peers from the node. Nodes can only respond with peers they currently know about, so this
// will not be an exhaustive or complete list of nodes on the network.
func (a *Arweave) PeerList() (PeerListResponse, error) {
	req, err := http.NewRequest("GET", a.fqdn()+"/peers", nil)
	if err != nil {
		return PeerListResponse{}, err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return PeerListResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		return PeerListResponse{}, ErrorNotOk(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return PeerListResponse{}, err
	}

	var pl PeerListResponse
	json.Unmarshal(body, &pl)

	return pl, nil
}
