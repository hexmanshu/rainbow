// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package dbram

import (
	"strings"

	"github.com/teal-finance/rainbow/pkg/rainbow"
)

// DB client.
type DB struct {
	// optionsByProvider separates the options by provider
	optionsByProvider map[string][]rainbow.Option
}

func NewDB() *DB {
	return &DB{
		optionsByProvider: map[string][]rainbow.Option{},
	}
}

func (db *DB) InsertOptions(options []rainbow.Option) error {
	for _, o := range options {
		if _, ok := db.optionsByProvider[o.Provider]; !ok {
			db.optionsByProvider[o.Provider] = []rainbow.Option{}
		}

		db.insertOption(o)
	}
	return nil
}

func (db *DB) insertOption(o rainbow.Option) {
	for i, oldOpt := range db.optionsByProvider[o.Provider] {
		if oldOpt.Name == o.Name &&
			oldOpt.Type == o.Type &&
			oldOpt.Expiry == o.Expiry &&
			oldOpt.Strike == o.Strike {
			// Update the option attributes
			db.optionsByProvider[o.Provider][i] = o
			return
		}
	}

	// Append the new option
	db.optionsByProvider[o.Provider] = append(db.optionsByProvider[o.Provider], o)
}

func (db *DB) Options(args rainbow.StoreArgs) ([]rainbow.Option, error) {
	n := 0
	for provider, o := range db.optionsByProvider {
		if len(args.Providers) > 0 {
			if !in(provider, args.Providers) {
				continue
			}
		}
		n += len(o)
	}

	options := make([]rainbow.Option, 0, n)
	for provider, o := range db.optionsByProvider {
		if len(args.Providers) > 0 {
			if !in(provider, args.Providers) {
				continue
			}
		}
		options = append(options, o...)
	}

	i := 0
	for _, o := range options {
		if len(args.Assets) > 0 {
			if !contains(o.Asset, args.Assets) {
				continue
			}
		}

		if len(args.Expiries) > 0 {
			if !startsWith(o.Expiry, args.Expiries) {
				continue
			}
		}

		options[i] = o
		i++
	}

	return options[:i], nil
}

// TODO go1.18: use generics with constraint comparable.
func contains(asset string, subStrings []string) bool {
	for _, substr := range subStrings {
		// use of strings.Contains because we allow some cases like WETH as ETH
		// TODO: sanitize data before put it in db
		if strings.Contains(asset, substr) {
			return true
		}
	}
	return false
}

func in(provider string, wanted []string) bool {
	for _, w := range wanted {
		if w == provider {
			return true
		}
	}
	return false
}

func startsWith(expiry string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(expiry, prefix) {
			return true
		}
	}
	return false
}
