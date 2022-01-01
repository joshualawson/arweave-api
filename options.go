package arweave

func WithClient(c Client) Option {
	return func(arweave *Arweave) *Arweave {
		arweave.client = c
		return arweave
	}
}

func WithServerAndPort(server string, port string) Option {
	return func(arweave *Arweave) *Arweave {
		arweave.server = server
		arweave.port = port
		return arweave
	}
}
