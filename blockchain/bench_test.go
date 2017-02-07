// Copyright (c) 2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package blockchain_test

import (
	"testing"

	"github.com/bitgo/rmgd/blockchain"
	"github.com/bitgo/rmgd/rmgutil"
)

// BenchmarkIsCoinBase performs a simple benchmark against the IsCoinBase
// function.
func BenchmarkIsCoinBase(b *testing.B) {
	tx, _ := rmgutil.NewBlock(&SomeBlock).Tx(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blockchain.IsCoinBase(tx)
	}
}

// BenchmarkIsCoinBaseTx performs a simple benchmark against the IsCoinBaseTx
// function.
func BenchmarkIsCoinBaseTx(b *testing.B) {
	tx := SomeBlock.Transactions[1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blockchain.IsCoinBaseTx(tx)
	}
}
