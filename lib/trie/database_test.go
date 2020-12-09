// Copyright 2019 The go-fdg Authors
// This file is part of the go-fdg library.
//
// The go-fdg library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-fdg library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-fdg library. If not, see <http://www.gnu.org/licenses/>.

package trie

import (
	"testing"

	"github.com/futuredigitalgames/fdg-chain/common"
	"github.com/futuredigitalgames/fdg-chain/lib/fdgdb/memorydb"
)

// Tests that the trie database returns a missing trie node error if attempting
// to retrieve the meta root.
func TestDatabaseMetarootFetch(t *testing.T) {
	db := NewDatabase(memorydb.New())
	if _, err := db.Node(common.Hash{}); err == nil {
		t.Fatalf("metaroot retrieval succeeded")
	}
}
