package arweave

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestArweave_BlocksByID(t *testing.T) {
	var tests = []struct {
		name       string
		blockHash  string
		response   Block
		json       string
		statusCode int
		err        error
	}{
		{
			name:      "get block should return 200 ok",
			blockHash: "*blockHash*",
			response: Block{
				Nonce:          "W3Jy4wp2LVbDFhGX_hUjRQZCkTdEbKxz45E5OVe52Lo",
				PreviousBlock:  "YuTyalVBTNB9t5KhuRezcIgxVz9PbQsbrcY4Tpkiu8XBPgglGM_Yql5qZd0c9PVG",
				Timestamp:      1586440919,
				LastRetarget:   1586440919,
				Diff:           "115792089039110416381168389782714091630053560834545856346499935466490404274176",
				Height:         422250,
				Hash:           "_____8422fLZnBsEsxtwEdpi8GZDHVT-aFlqroQDG44",
				IndepHash:      "5VTARz7bwDO4GqviCSI9JXm8_JOtoQwF-QCZm0Gt2gVgwdzSY3brOtOD46bjMz09",
				Txs:            []string{"IRPCjc_ws7aS5GWp4mwR2k-HuQy-zT_GWrgR6kRdbmI"},
				TxRoot:         "lsoo-p3Tj7oblZ-54WVPHoVguqgw5rA9Jf3lLH6H8zY",
				TxTree:         []string{},
				WalletList:     "N5NJtXhgH9bPmXoSopehcr_zqwyPjjg3igel0V8G1DdLk_BYdoRVIBsqjVA9JmFc",
				RewardAddr:     "Oox7m4HIcVhUtMd6AUuGtlaOoSCmREUNPyyKQCbz4d4",
				Tags:           []string{},
				RewardPool:     3026104059201252,
				WeaveSize:      407672420044,
				BlockSize:      937455,
				CumulativeDiff: "99416580392277",
				HashListMerkle: "akSjDrBKPuepJMOhO_S9C-iFp5zn9Glv57HGdN_WPqEToWC0Ukb37Gzs4PDA7oLU",
				Poa: Poa{
					Option:   "1",
					TxPath:   "xZ6vhVXw_0BlD-Xkv3KtfnJeLXykjkjUrwcPsXw2JUnie021At7I-fMZkt5EF_xOHtcdq4RIqXto1gwFAM5eZgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAfDSbuKpWzKZ9HP_N2I4gX6cUujNsJtelJULjHmbZp0XzmkBljlK4S1PMlSrTePIjfJdRfqvFNE8idpnj69X1P0zAfwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAn4ybxD6lgdArqnPJzs7t8bU-7KfEb1YqpAOvbr6q3vmP-MWnCTWZJKTL90azeYZmHrTMx-iutuT6bP6CUC7zgHAfGgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAmTpFIGvz18gKl5rZ6p2Ve4yVeRzWNwibyVTKz80HSBYprfIpVJk9oRG3E5q1xRn5wErqyH2vFLbsLxDqKcR0vLunBwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAfDwBRWXT_vDxcaBxGmihJwlU_n_PFBCOsP-Lx3hSG6H6UGesIMAEYMmd2c5QixR-fCimhm_9S582cLzSUffsrAHliQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAmP-RTrBhY9xCC1yywyehB7X6EmlBjyQBqm0y1L9Ex_dkswkf50rG-LE29UJP4st0bzFthHukfHvvWZY3bgIiog3L7gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAfD3YxQguhfH8daMBAQrveQq3MMp4iKB3khk5mbU34Ckl1q8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJj_kQ",
					DataPath: "bTVpffiN3SSDeqBEJpKiXegQGKKnprS_AFMh6zz4QRIU-8dJuvFzyKxqjkDHQvtKl0Eajfm18yZsjaAJkNhbAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAOH0cuoLq1CTbSelF9C59C-fcO3a3ywoceaNxRl4nQQH1BuwcpiNdDdZvEz6Pfk5wKbnsF_VwVIgrfcLZgsxoKwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAefOoaNyW7ORmrzbZ5O7midzLByHooxjM5oEMJfZbQsY9mKS14G9fUEFmFaCPPJX6EXVGrUwROzDIWfHf8oHErAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAktmxYyC7BSV-MULrjzgdJJYfJY7lDFcKe3mo_EX19xoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAA",
					Chunk:    "*chunk*",
				},
			},
			json: `{
    "nonce": "W3Jy4wp2LVbDFhGX_hUjRQZCkTdEbKxz45E5OVe52Lo",
    "previous_block": "YuTyalVBTNB9t5KhuRezcIgxVz9PbQsbrcY4Tpkiu8XBPgglGM_Yql5qZd0c9PVG",
    "timestamp": 1586440919,
    "last_retarget": 1586440919,
    "diff": "115792089039110416381168389782714091630053560834545856346499935466490404274176",
    "height": 422250,
    "hash": "_____8422fLZnBsEsxtwEdpi8GZDHVT-aFlqroQDG44","indep_hash":"5VTARz7bwDO4GqviCSI9JXm8_JOtoQwF-QCZm0Gt2gVgwdzSY3brOtOD46bjMz09","txs":["IRPCjc_ws7aS5GWp4mwR2k-HuQy-zT_GWrgR6kRdbmI"],
    "tx_root": "lsoo-p3Tj7oblZ-54WVPHoVguqgw5rA9Jf3lLH6H8zY",
    "tx_tree":[],
    "wallet_list":"N5NJtXhgH9bPmXoSopehcr_zqwyPjjg3igel0V8G1DdLk_BYdoRVIBsqjVA9JmFc","reward_addr":"Oox7m4HIcVhUtMd6AUuGtlaOoSCmREUNPyyKQCbz4d4",
    "tags":[],
    "reward_pool":3026104059201252,
    "weave_size": 407672420044,
    "block_size": 937455,
    "cumulative_diff": "99416580392277",
    "hash_list_merkle": "akSjDrBKPuepJMOhO_S9C-iFp5zn9Glv57HGdN_WPqEToWC0Ukb37Gzs4PDA7oLU",
    "poa": {
        "option":"1",
        "tx_path": "xZ6vhVXw_0BlD-Xkv3KtfnJeLXykjkjUrwcPsXw2JUnie021At7I-fMZkt5EF_xOHtcdq4RIqXto1gwFAM5eZgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAfDSbuKpWzKZ9HP_N2I4gX6cUujNsJtelJULjHmbZp0XzmkBljlK4S1PMlSrTePIjfJdRfqvFNE8idpnj69X1P0zAfwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAn4ybxD6lgdArqnPJzs7t8bU-7KfEb1YqpAOvbr6q3vmP-MWnCTWZJKTL90azeYZmHrTMx-iutuT6bP6CUC7zgHAfGgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAmTpFIGvz18gKl5rZ6p2Ve4yVeRzWNwibyVTKz80HSBYprfIpVJk9oRG3E5q1xRn5wErqyH2vFLbsLxDqKcR0vLunBwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAfDwBRWXT_vDxcaBxGmihJwlU_n_PFBCOsP-Lx3hSG6H6UGesIMAEYMmd2c5QixR-fCimhm_9S582cLzSUffsrAHliQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAmP-RTrBhY9xCC1yywyehB7X6EmlBjyQBqm0y1L9Ex_dkswkf50rG-LE29UJP4st0bzFthHukfHvvWZY3bgIiog3L7gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAfD3YxQguhfH8daMBAQrveQq3MMp4iKB3khk5mbU34Ckl1q8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJj_kQ",
        "data_path": "bTVpffiN3SSDeqBEJpKiXegQGKKnprS_AFMh6zz4QRIU-8dJuvFzyKxqjkDHQvtKl0Eajfm18yZsjaAJkNhbAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAOH0cuoLq1CTbSelF9C59C-fcO3a3ywoceaNxRl4nQQH1BuwcpiNdDdZvEz6Pfk5wKbnsF_VwVIgrfcLZgsxoKwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAefOoaNyW7ORmrzbZ5O7midzLByHooxjM5oEMJfZbQsY9mKS14G9fUEFmFaCPPJX6EXVGrUwROzDIWfHf8oHErAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAktmxYyC7BSV-MULrjzgdJJYfJY7lDFcKe3mo_EX19xoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAA",
        "chunk":"*chunk*"
    }
}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "get block should return 400 Bad Request",
			blockHash:  "*blockHash*",
			response:   Block{},
			json:       ``,
			statusCode: http.StatusNotFound,
			err:        ErrorNotOk(http.StatusNotFound),
		},
	}
	c := &MockClient{}

	arweave := New(
		WithClient(c),
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c.On("Do", mock.Anything).Return(&http.Response{StatusCode: tt.statusCode, Body: ioutil.NopCloser(bytes.NewReader([]byte(tt.json)))}, nil)

			res, err := arweave.Block(tt.blockHash)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.response, res)

			c.ExpectedCalls = []*mock.Call{}
		})
	}
}
