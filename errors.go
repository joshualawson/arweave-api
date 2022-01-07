package arweave

import "fmt"

var (
	ErrorNotOk func(statusCode int) error = func(statusCode int) error {
		return fmt.Errorf("status code not 200 was [%v]", statusCode)
	}

	JsonMarshalError func(e error) error = func(e error) error {
		return fmt.Errorf("unable to mashal type to json: %w", e)
	}
)
