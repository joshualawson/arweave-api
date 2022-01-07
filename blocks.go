package arweave

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Block struct {
	Nonce          string   `json:"nonce"`
	PreviousBlock  string   `json:"previous_block"`
	Timestamp      int64    `json:"timestamp"`
	LastRetarget   int64    `json:"last_retarget"`
	Diff           string   `json:"diff"`
	Height         int64    `json:"height"`
	Hash           string   `json:"hash"`
	IndepHash      string   `json:"indep_hash"`
	Txs            []string `json:"txs"`
	TxRoot         string   `json:"tx_root"`
	TxTree         []string `json:"tx_tree"`
	WalletList     string   `json:"wallet_list"`
	RewardAddr     string   `json:"reward_addr"`
	Tags           []string `json:"tags"`
	RewardPool     int64    `json:"reward_pool"`
	WeaveSize      int64    `json:"weave_size"`
	BlockSize      int64    `json:"block_size"`
	CumulativeDiff string   `json:"cumulative_diff"`
	HashListMerkle string   `json:"hash_list_merkle"`
	Poa            Poa      `json:"poa"`
}

type Poa struct {
	Option   string `json:"option"`
	TxPath   string `json:"tx_path"`
	DataPath string `json:"data_path"`
	Chunk    string `json:"chunk"`
}

// Block Get a block by its hash (idep_hash).
func (a *Arweave) Block(blockHash string) (Block, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/block/hash/%s", a.fqdn(), blockHash), nil)
	if err != nil {
		return Block{}, err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return Block{}, err
	}

	if res.StatusCode != http.StatusOK {
		return Block{}, ErrorNotOk(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Block{}, err
	}

	var b Block
	json.Unmarshal(body, &b)

	return b, nil
}
