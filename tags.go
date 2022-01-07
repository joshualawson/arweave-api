package arweave

type Tags []Tag

type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
