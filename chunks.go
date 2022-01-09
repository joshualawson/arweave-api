package arweave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Chunk struct {
	DataRoot string `json:"data_root"`
	DataSize string `json:"data_size"`
	DataPath string `json:"data_path"`
	Chunk    string `json:"chunk"`
	Offset   string `json:"offset"`
}

type ChunkOffset struct {
	Offset string `json:"offset"`
	Size   string `json:"size"`
}

// UploadChunk Upload Data Chunks
func (a *Arweave) UploadChunk(chunk Chunk) error {
	j, err := json.Marshal(chunk)
	if err != nil {
		return ErrorJsonMarshal(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/chunk", a.fqdn()), bytes.NewReader(j))
	if err != nil {
		return err
	}

	res, err := a.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusBadRequest {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		var errorMessage Error
		json.Unmarshal(body, &errorMessage)

		return ErrorBadRequest(errorMessage)
	}

	if res.StatusCode != http.StatusOK {
		return ErrorNotOk(res.StatusCode)
	}

	return nil
}
