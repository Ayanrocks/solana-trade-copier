package tradeCopier

import (
	"time"
)

const (
	BuyTransaction  = "buy"
	SellTransaction = "sell"
)

type TradeCopier struct {
	Threshold        int                `json:"threshold"`
	TimeWindowInMins int                `json:"time_window_in_mins"`
	Trades           map[string][]Trade `json:"trades"`
}

// NewTradeCopier returns a new instance of TradeCopier with threshold and timeWindow set
func NewTradeCopier(threshold int, timeWindow int) *TradeCopier {
	return &TradeCopier{
		Threshold:        threshold,
		TimeWindowInMins: timeWindow,
		Trades:           make(map[string][]Trade),
	}
}

// Add each trade as they appear
func (tc *TradeCopier) AddTrade(trade Trade) {
	// if trade is not in the map, create a new slice
	if tc.Trades[trade.Token] == nil {
		tc.Trades[trade.Token] = []Trade{}
	}
	// append the trade to the slice
	tc.Trades[trade.Token] = append(tc.Trades[trade.Token], trade)
}

// IsEligible checks if the trade is eligible to be copied to the source wallet
func (tc *TradeCopier) CopyTrade(sourceWallet *Wallet, trade *Trade) {
	if trade.Action != BuyTransaction && trade.Action != SellTransaction {
		return
	}

	// check the eligibility of the trade
	if tc.IsEligible(trade) {
		// copy the trade to the wallet
		newTradeDetails := Trade{
			WalletAddr: sourceWallet.Address,
			Action:     trade.Action,
			Token:      trade.Token,
			Timestamp:  time.Now(),
			Price:      sourceWallet.BuyAmount,
		}
		sourceWallet.Trade = append(sourceWallet.Trade, newTradeDetails)
	}
}

// IsEligible checks if the trade is eligible to be copied to the source wallet based on certain condition
func (tc *TradeCopier) IsEligible(trade *Trade) bool {
	// check if the trade is eligible
	// get the token from
	tokenTrades := tc.Trades[trade.Token]

	// check how many unique wallet addresses have traded the token
	uniqueWallets := make(map[string]Trade)
	for _, t := range tokenTrades {
		uniqueWallets[t.WalletAddr] = Trade{
			Action:    t.Action,
			Timestamp: t.Timestamp,
		}
	}

	if len(uniqueWallets) >= tc.Threshold {
		// check if the trade is within the time window from last transaction
		for _, wallet := range uniqueWallets {
			// if current time - timeWindow is greater than last transaction time, then return false, meaning the last transaction was not within the time window
			if trade.Timestamp.Add(-1*time.Duration(tc.TimeWindowInMins)*time.Minute).Sub(wallet.Timestamp) > 0 {
				return false
			}
		}

		return true
	}

	return false
}
