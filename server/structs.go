package server

type Counters struct {
	Pending  int `json:"pending"`
	Verified int `json:"verified"`
}

type BlockCounts struct {
	Local  uint64 `json:"local"`
	Remote uint64 `json:"remote"`
}

type DetailReply struct {
	Chain               string      `json:"chain"`
	Mode                string      `json:"mode"`
	Blocks              BlockCounts `json:"blocks"`
	RPCs                uint64      `json:"rpcs"`
	Peers               BlockCounts `json:"peers"`
	TransactionCounters Counters    `json:"transactionCounters"`
	Difficulty          float64     `json:"difficulty"`
	Version             string      `json:"version"`
	Uptime              string      `json:"uptime"`
	PublicKey           string      `json:"publicKey"`
}
