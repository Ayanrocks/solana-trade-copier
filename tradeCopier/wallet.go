package tradeCopier

import (
	"fmt"
	"time"
)

type Wallet struct {
	Address   string  `json:"address"`
	BuyAmount float64 `json:"buy_amount"`
	Trade     []Trade `json:"trade"`
}

func (w *Wallet) PrintTrades() {
	for _, trade := range w.Trade {
		fmt.Printf("Executing Trade on %s: %s token %s for amount: %.2f at %s\n", trade.WalletAddr, trade.Action, trade.Token, trade.Price, trade.Timestamp.Format(time.StampMilli))
	}
}
