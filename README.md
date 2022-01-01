# Arweave API

Go implementation of the Arweave API

## Todo
A list of endpoints that are complete or need to be done, if the endpoint is marked as done it means it includes
unit tests and an example

### Transactions
- [ ] Get Transaction by ID
- [ ] Get Transaction Status
- [ ] Get Transaction Field
- [ ] Get Decoded Transaction Data (http://arweave.net/{id})
- [ ] Get Transaction Data Encoded (http://arweave.net/tx/{id}/data.{extension})
- [ ] Get Transaction Price
- [ ] Submit Transaction

### Wallets
- [x] Get Wallet Balance
- [x] Get Last Transaction ID

### Blocks
- [x] Get Block by ID

### Network and Node State
- [x] Network Info
- [x] Peer List

### Chunks
#### Upload Chunks
- [ ] Upload Chunks
#### Download Chunks
- [ ] Get Transaction Data
- [ ] Get Transaction Offset and Size
