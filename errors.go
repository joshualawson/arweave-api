package arweave

import "fmt"

var (
	ErrorNotOk func(statusCode int) error = func(statusCode int) error {
		return fmt.Errorf("status code not 200 was [%v]", statusCode)
	}

	ErrorJsonMarshal func(e error) error = func(e error) error {
		return fmt.Errorf("unable to marshal type to json: %w", e)
	}

	ErrorJsonUnmarshal func(e error) error = func(e error) error {
		return fmt.Errorf("unable to unmarshal json to type: %w", e)
	}

	ErrorUnmarshalTextToBigInt func(e error) error = func(e error) error {
		return fmt.Errorf("unable to unmarshal text to big.Int: %w", e)
	}

	ErrorBadRequest func(e error) error = func(e error) error {
		return fmt.Errorf("bad request: %w", e)
	}
)

type Error struct {
	Message string `json:"error"`
}

func (e Error) Error() string {
	return e.Message
}
