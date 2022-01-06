// Copyright (C) 2019-2022 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package compactcert

import (
	"github.com/algorand/go-algorand/crypto"
)

// HashType/ hashSize relate to the type of hash this package uses.
const (
	HashType = crypto.Sumhash
	HashSize = crypto.SumhashDigestSize
)

const (
	// MaxReveals is a bound on allocation and on numReveals to limit log computation
	MaxReveals = 1024

	// MaxProofDigests is a bound on allocation on number of proofs
	MaxProofDigests = 20 * MaxReveals
)