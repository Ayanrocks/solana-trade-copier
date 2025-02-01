package tradeCopier

import "time"

type Trade struct {
	WalletAddr string    `json:"wallet_addr"`
	Action     string    `json:"action"`
	Token      string    `json:"token"`
	Timestamp  time.Time `json:"timestamp"`
	Price      float64   `json:"price"`
}
