package arweave

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Info struct {
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

type PeerList []string

// Info Get the current network information including height, current block, and other properties.
func (a *Arweave) Info() (Info, error) {
	req, err := http.NewRequest("GET", a.fqdn()+"/info", nil)
	if err != nil {
		return Info{}, err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return Info{}, err
	}

	if res.StatusCode != http.StatusOK {
		return Info{}, ErrorNotOk(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Info{}, err
	}

	var ir Info
	if err := json.Unmarshal(body, &ir); err != nil {
		return Info{}, ErrorJsonUnmarshal(err)
	}

	return ir, nil
}

// PeerList Get the list of peers from the node. Nodes can only respond with peers they currently know about, so this
// will not be an exhaustive or complete list of nodes on the network.
func (a *Arweave) PeerList() (PeerList, error) {
	req, err := http.NewRequest("GET", a.fqdn()+"/peers", nil)
	if err != nil {
		return PeerList{}, err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return PeerList{}, err
	}

	if res.StatusCode != http.StatusOK {
		return PeerList{}, ErrorNotOk(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return PeerList{}, err
	}

	var pl PeerList
	if err := json.Unmarshal(body, &pl); err != nil {
		return PeerList{}, ErrorJsonUnmarshal(err)
	}

	return pl, nil
}
