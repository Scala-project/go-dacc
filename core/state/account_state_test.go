// Copyright (C) 2017 go-dacc authors
//
// This file is part of the go-dacc library.
//
// the go-dacc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-dacc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-dacc library.  If not, see <http://www.gnu.org/licenses/>.
//

package state

import (
	"testing"

	"github.com/daccproject/go-dacc/common/trie"
	"github.com/daccproject/go-dacc/storage"
	"github.com/daccproject/go-dacc/util"
	"github.com/stretchr/testify/assert"
)

func TestAccount_ToBytes(t *testing.T) {
	stor, _ := storage.NewMemoryStorage()
	vars, _ := trie.NewTrie(nil, stor, false)
	acc := &account{
		balance:    util.NewUint128(),
		nonce:      0,
		variables:  vars,
		birthPlace: []byte("0x0"),
	}
	bytes, _ := acc.ToBytes()
	a := &account{}
	a.FromBytes(bytes, stor)
	assert.Equal(t, acc, a)
}

func TestAccountState(t *testing.T) {
	stor, err := storage.NewMemoryStorage()
	assert.Nil(t, err)
	as, err := NewAccountState(nil, stor)
	assert.Nil(t, err)
	accAddr1 := []byte("accAddr1")
	acc1, err := as.GetOrCreateUserAccount(accAddr1)
	assert.Nil(t, err)
	assert.Equal(t, acc1.Balance(), util.NewUint128())
	assert.Equal(t, acc1.Nonce(), uint64(0))
	value, _ := util.NewUint128FromInt(16)
	acc1.AddBalance(value)
	acc1.IncrNonce()
	acc1.Put([]byte("var0"), []byte("value0"))

	asClone, err := as.Clone()
	assert.Nil(t, err)
	acc1Clone, err := asClone.GetOrCreateUserAccount(accAddr1)
	assert.Nil(t, err)
	value0, err := acc1Clone.Get([]byte("var0"))
	assert.Nil(t, err)
	assert.Equal(t, value0, []byte("value0"))
	asRoot := as.RootHash()
	assert.Nil(t, err)
	asCloneRoot := asClone.RootHash()
	assert.Nil(t, err)
	assert.Equal(t, asRoot, asCloneRoot)
	assert.Equal(t, acc1Clone.VarsHash(), acc1.VarsHash())
	accAddr2 := []byte("accAddr2")
	acc2, err := as.GetOrCreateUserAccount(accAddr2)
	assert.Nil(t, err)
	acc2.Put([]byte("var1"), []byte("value1"))
	accAddr3 := []byte("accAddr3")
	acc3, err := as.GetOrCreateUserAccount(accAddr3)
	assert.Nil(t, err)
	acc3.Put([]byte("var2"), []byte("value2"))
}
