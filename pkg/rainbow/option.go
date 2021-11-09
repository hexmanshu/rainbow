// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package rainbow

import (
	"strconv"
)

type Option struct {
	Name          string // ASSET-DATE-Strike-OptionsType
	Type          string // CALL / PUT
	Asset         string // ETH, BTC, SOL
	Expiry        string // Expiry date in format 2021-12-31
	Strike        float64
	ExchangeType  string // CEX / DEX
	Chain         string // Ethereum, Solana and "–" for CEX (Deribit)
	Layer         string // L1, L2 and "–" for CEX (Deribit)
	Provider      string // Opyn, Lyra, Thales, Deribit, Psyoptions
	QuoteCurrency string // ETH, BTC, USDT...
	Bid           []Order
	Ask           []Order
}

type Order struct {
	Px   float64 `json:"px"`   // Px is an abbreviation for "Price" (e.g. FIX protocol).
	Size float64 `json:"size"` // Size is often used in lieu of the longer word "Quantity".
}

func BestLimitStr(option Option) (bestBidPx, bestBidSz, bestAskPx, bestAskSz string) {
	bestBidPx, bestBidSz = none, none
	bestAskPx, bestAskSz = none, none

	if len(option.Bid) > 0 {
		bestBidPx = alignFloatOnDecimalPoint(option.Bid[0].Px)
		bestBidSz = alignFloatOnDecimalPoint(option.Bid[0].Size)
	}

	if len(option.Ask) > 0 {
		bestAskPx = alignFloatOnDecimalPoint(option.Ask[0].Px)
		bestAskSz = alignFloatOnDecimalPoint(option.Ask[0].Size)
	}

	return
}

// none must contain the same number of runes as len(spaces)-1.
const none = "   —"

// these two variables avoid unnecessary allocation by alignFloatOnDecimalPoint().
var (
	spaces = [5]byte{32, 32, 32, 32, 32} // len(spaces) must be the max digits wanted before the decimal point
	b      = make([]byte, 0, 10)
)

func alignFloatOnDecimalPoint(f float64) string {
	b = strconv.AppendFloat(b[:0], f, 'f', -1, 64)

	var i int // position of the '.' (else i=len(b) if not found)
	for i = 1; i < len(b); i++ {
		if b[i] == '.' {
			break
		}
	}

	if i < len(spaces) {
		return string(append(spaces[:len(spaces)-i-1], b...))
	}

	return string(b)
}

// WIP new implementation
//
// type Order struct{ Px, Qty string }
// type Limit struct{ Bid, Ask Order }
//
// func NewLimit() Limit {
// 	return Limit{
// 		Bid: Order{
// 			Px:  none,
// 			Qty: none,
// 		},
// 		Ask: Order{
// 			Px:  none,
// 			Qty: none,
// 		},
// 	}
// }
