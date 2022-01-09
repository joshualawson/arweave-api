package arweave

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

// WalletBalance Get the balance for a given wallet. Unknown wallet addresses will simply return 0.
func (a *Arweave) WalletBalance(address string) (*big.Int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/wallet/%s/balance", a.fqdn(), address), nil)
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

	i := new(big.Int)
	if err := i.UnmarshalText(body); err != nil {
		return nil, ErrorUnmarshalTextToBigInt(err)
	}

	return i, nil
}

// WalletLastTransactionID Get the last outgoing transaction for the given wallet address.
func (a *Arweave) WalletLastTransactionID(address string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/wallet/%s/last_tx", a.fqdn(), address), nil)
	if err != nil {
		return "", err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", ErrorNotOk(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
