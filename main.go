package main

import (
	"fmt"
	"time"
	"trade-copier-solana/tradeCopier"
)

func main() {

	threshold := 2
	timeWindow := 15

	fmt.Println("TestCase One")
	TestCase1(threshold, timeWindow)

	fmt.Println("TestCase Two")
	TestCase2(threshold, timeWindow)

	fmt.Println("TestCase Three")
	TestCase3(threshold, timeWindow)

	fmt.Println("TestCase Three")
	TestCase4(threshold, timeWindow)
}

func TestCase1(threshold, timeWindow int) {
	tc := tradeCopier.NewTradeCopier(threshold, timeWindow)

	walletToCopy := tradeCopier.Wallet{
		Address:   "C",
		BuyAmount: 0.05,
		Trade:     make([]tradeCopier.Trade, 0),
	}

	// upstream
	trades := []tradeCopier.Trade{
		{
			WalletAddr: "addr1",
			Action:     tradeCopier.BuyTransaction,
			Token:      "xyz",
			Timestamp:  time.Now().Add(time.Minute * 10),
			Price:      0.4,
		},
		{
			WalletAddr: "addr1",
			Action:     tradeCopier.BuyTransaction,
			Token:      "xyz",
			Timestamp:  time.Now().Add(time.Minute * 14),
			Price:      0.2,
		},
		{
			WalletAddr: "addr1",
			Action:     tradeCopier.SellTransaction,
			Token:      "xyz",
			Timestamp:  time.Now().Add(time.Minute * 18),
			Price:      0.5,
		},
		{
			WalletAddr: "addr1",
			Action:     tradeCopier.BuyTransaction,
			Token:      "abc",
			Timestamp:  time.Now().Add(time.Minute * 20),
			Price:      0.2,
		},
		{
			WalletAddr: "addr2",
			Action:     tradeCopier.BuyTransaction,
			Token:      "xyz",
			Timestamp:  time.Now().Add(time.Minute * 26),
			Price:      0.05,
		},
		{
			WalletAddr: "addr2",
			Action:     tradeCopier.BuyTransaction,
			Token:      "abc",
			Timestamp:  time.Now().Add(time.Minute * 30),
			Price:      0.15,
		},
	}

	// acts as a listener to upstream to get new trades
	for _, t := range trades {
		tc.AddTrade(t)
		tc.CopyTrade(&walletToCopy, &t)
	}

	walletToCopy.PrintTrades()
}

func TestCase2(threshold, timeWindow int) {
	baseTime := time.Now()
	tc := tradeCopier.NewTradeCopier(threshold, timeWindow)

	walletToCopy := tradeCopier.Wallet{
		Address:   "D",
		BuyAmount: 0.05,
		Trade:     make([]tradeCopier.Trade, 0),
	}

	// Create a slice of Trade structs for the two activities.
	trades := []tradeCopier.Trade{
		{
			WalletAddr: "Addr1",
			Action:     "buy",
			Token:      "xyz",
			Timestamp:  baseTime.Add(10 * time.Minute), // min 10
			Price:      0.4,
		},
		{
			WalletAddr: "Addr2",
			Action:     "buy",
			Token:      "xyz",
			Timestamp:  baseTime.Add(14 * time.Minute), // min 14
			Price:      0.2,
		},
	}

	// acts as a listener to upstream to get new trades
	for _, t := range trades {
		tc.AddTrade(t)
		tc.CopyTrade(&walletToCopy, &t)
	}

	walletToCopy.PrintTrades()
}

func TestCase3(threshold, timeWindow int) {
	baseTime := time.Now()
	tc := tradeCopier.NewTradeCopier(threshold, timeWindow)
	walletToCopy := tradeCopier.Wallet{

		Address:   "E",
		BuyAmount: 0.05,
		Trade:     make([]tradeCopier.Trade, 0),
	}

	trades := []tradeCopier.Trade{
		{
			WalletAddr: "Addr1",
			Action:     "buy",
			Token:      "xyz",
			Timestamp:  baseTime.Add(2 * time.Minute),
			Price:      0.5,
		},
		{
			WalletAddr: "Addr2",
			Action:     "sell",
			Token:      "xyz",
			Timestamp:  baseTime.Add(3 * time.Minute),
			Price:      0.45,
		},
		{
			WalletAddr: "Addr3",
			Action:     "buy",
			Token:      "abc",
			Timestamp:  baseTime.Add(4 * time.Minute),
			Price:      1.2,
		},
		{
			WalletAddr: "Addr1",
			Action:     "sell",
			Token:      "abc",
			Timestamp:  baseTime.Add(6 * time.Minute),
			Price:      1.1,
		},
	}

	for _, t := range trades {
		tc.AddTrade(t)
		tc.CopyTrade(&walletToCopy, &t)
	}

	walletToCopy.PrintTrades()
}

func TestCase4(threshold, timeWindow int) {
	baseTime := time.Now()
	tc := tradeCopier.NewTradeCopier(threshold, timeWindow)
	walletToCopy := tradeCopier.Wallet{

		Address:   "F",
		BuyAmount: 0.05,
		Trade:     make([]tradeCopier.Trade, 0),
	}

	sameTime := baseTime.Add(20 * time.Minute)
	trades := []tradeCopier.Trade{
		{
			WalletAddr: "Addr1",
			Action:     "buy",
			Token:      "def",
			Timestamp:  sameTime,
			Price:      0.75,
		},
		{
			WalletAddr: "Addr2",
			Action:     "buy",
			Token:      "def",
			Timestamp:  sameTime,
			Price:      0.80,
		},
	}
	for _, t := range trades {
		tc.AddTrade(t)
		tc.CopyTrade(&walletToCopy, &t)
	}

	walletToCopy.PrintTrades()
}
