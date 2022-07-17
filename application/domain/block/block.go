package block

type Block struct {
	Number        uint64 `json:"block_number"`
	Hash          string `json:"block_hash"`
	Time          uint64 `json:"block_time"`
	ParentHash    string `json:"parent_hash"`
	Confirmations uint64 `json:"confirmations"`
}

type BlockDetail struct {
	Number        uint64   `json:"block_number"`
	Hash          string   `json:"block_hash"`
	Time          uint64   `json:"block_time"`
	ParentHash    string   `json:"parent_hash"`
	Confirmations uint64   `json:"confirmations"`
	Transactions  []string `json:"transactions"`
}

type Transaction struct {
	Hash  string           `json:"tx_hash"`
	From  string           `json:"from"`
	To    string           `json:"to"`
	Nonce uint64           `json:"nonce"`
	Value string           `json:"value"`
	Data  string           `json:"data"`
	Logs  []TransactionLog `json:"logs"`
}

type TransactionLog struct {
	Index int    `json:"index"`
	Data  string `json:"data"`
}
