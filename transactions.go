package arweave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

type Transaction struct {
	Format    int    `json:"format"`
	Id        string `json:"id"`
	LastTx    string `json:"last_tx"`
	Owner     string `json:"owner"`
	Tags      Tags   `json:"tags"`
	Target    string `json:"target"`
	Quantity  string `json:"quantity"`
	DataRoot  string `json:"data_root"`
	Data      string `json:"data"`
	DataSize  string `json:"data_size"`
	Reward    string `json:"reward"`
	Signature string `json:"signature"`
	Pending   bool   `json:"-"`
}

type TransactionStatus struct {
	BlockHeight           int    `json:"block_height"`
	BlockIndepHash        string `json:"block_indep_hash"`
	NumberOfConfirmations int    `json:"number_of_confirmations"`
}

type TransactionOffset struct {
	Offset string `json:"offset"`
	Size   string `json:"size"`
}

// Transaction Get a transaction by its ID.
func (a *Arweave) Transaction(id string) (Transaction, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/tx/%s", a.fqdn(), id), nil)
	if err != nil {
		return Transaction{}, err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return Transaction{}, err
	}

	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusAccepted {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return Transaction{}, err
		}

		var t Transaction
		if res.StatusCode == http.StatusOK {
			if err := json.Unmarshal(body, &t); err != nil {
				return Transaction{}, ErrorJsonUnmarshal(err)
			}
		} else {
			t.Pending = true
		}

		return t, nil
	}

	return Transaction{}, ErrorNotOk(res.StatusCode)
}

// TransactionStatus Get the status of a transaction
func (a *Arweave) TransactionStatus(id string) (TransactionStatus, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/tx/%s/status", a.fqdn(), id), nil)
	if err != nil {
		return TransactionStatus{}, err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return TransactionStatus{}, err
	}

	if res.StatusCode != http.StatusOK {
		return TransactionStatus{}, ErrorNotOk(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return TransactionStatus{}, err
	}

	var ts TransactionStatus
	if err := json.Unmarshal(body, &ts); err != nil {
		return TransactionStatus{}, ErrorJsonUnmarshal(err)
	}

	return ts, nil
}

// TransactionField Get a single field from a transaction.
func (a *Arweave) TransactionField(id string, field string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/tx/%s/%s", a.fqdn(), id, field), nil)
	if err != nil {
		return "", err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK || res.StatusCode != http.StatusAccepted {
		return "", ErrorNotOk(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Data returns the data that is with the transaction, along with the content type
func (a *Arweave) Data(id string) ([]byte, string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", a.fqdn(), id), nil)
	if err != nil {
		return nil, "", err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return nil, "", err
	}

	if res.StatusCode != http.StatusOK {
		return nil, "", ErrorNotOk(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, "", err
	}

	ct := res.Header.Get("content-type")

	return body, ct, nil
}

// TransactionData The endpoint serves data regardless of how it was uploaded
func (a *Arweave) TransactionData(id string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/tx/%s/data", a.fqdn(), id), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, ErrorNotOk(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (a *Arweave) TransactionOffset(id string) (TransactionOffset, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/tx/%s/offset", a.fqdn(), id), nil)
	if err != nil {
		return TransactionOffset{}, err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return TransactionOffset{}, err
	}

	if res.StatusCode != http.StatusOK {
		return TransactionOffset{}, ErrorNotOk(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return TransactionOffset{}, err
	}

	var o TransactionOffset
	if err := json.Unmarshal(body, &o); err != nil {
		return TransactionOffset{}, ErrorJsonUnmarshal(err)
	}

	return o, nil
}

// TransactionPrice This endpoint is used to calculate the minimum fee (reward) for a transaction of a specific size,
// and possibly to a specific address.This endpoint should always be used to calculate transaction fees as closely to
// the submission time as possible. Pricing is dynamic and determined by the network, so it's not always possible to
// accurately calculate prices offline or ahead of time.Transactions with a fee that's too low will simply be rejected.
func (a *Arweave) TransactionPrice(bytes string, target *string) (*big.Int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/price/%s/%s", a.fqdn(), bytes, target), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, ErrorNotOk(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	p := new(big.Int)
	if err := p.UnmarshalText(body); err != nil {
		return nil, ErrorUnmarshalTextToBigInt(err)
	}

	return p, nil
}

// SubmitTransaction Submit a new transaction to the network.The request body should be a JSON object with the
// attributes described in Transaction Format.
func (a *Arweave) SubmitTransaction(t Transaction) (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return "", ErrorJsonMarshal(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/tx", a.fqdn()), bytes.NewReader(b))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := a.client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK || res.StatusCode != http.StatusAlreadyReported {
		return "", ErrorNotOk(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
