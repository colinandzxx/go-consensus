/*
 * Copyright (c) 2019
 *
 * This project is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * @File: interface.go
 * @LastModified: 2019-10-08 17:31:11
 */

package consensus

import (
	"github.com/colinandzxx/go-consensus/types"
	"github.com/tinylib/msgp/msgp"
	"math/big"
)

// ================= chain interface =================

type Header interface {
	//Marshal() ([]byte, error)

	GetHeight() uint64
	GetHash() types.Byte32
	GetParentHash() types.Byte32
	GetNonce() uint64
	GetTimestamp() uint64
	GetCoinbase() types.Bytes
	GetConsensusData() Data
}

type Block interface{
	//Marshal() ([]byte, error)

	GetHeader() Header
	GetTotalDifficulty() *big.Int
}

type ChainReader interface {
	GetCurrentHeader() Header
	GetHeader(hash types.Byte32, height uint64) Header
	GetHeaderByHeight(height uint64) Header
	GetHeaderByHash(hash types.Byte32) Header
	GetBlock(hash types.Byte32, height uint64) Block
}

// ================= consensus data interface =================

type DataWrapper interface {
	Wrap(chain ChainReader, height uint64) ([]byte, error)
}

type DataUnWrapper interface {
	UnWrap([]byte) (Data, error)
}

type Data interface {
	//encoding.TextMarshaler
	//encoding.TextUnmarshaler
	//json.Unmarshaler
	//json.Marshaler
	//fmt.Stringer
	//fmt.Formatter

	// for msgp
	msgp.Decodable
	msgp.Encodable
	msgp.Marshaler
	msgp.Unmarshaler
	msgp.Sizer

	DataWrapper
	DataUnWrapper
}

// ================= engine interface =================

type Engine interface {
	Stuffer
	Verifier
	Forger
}

type Stuffer interface {
	CalculateDifficulty(chain ChainReader, height uint64) *big.Int
}

type Verifier interface {
	// VerifyHeader checks whether a header conforms to the consensus rules of a
	// given engine.
	VerifyHeader(chain ChainReader, header Header) error
	VerifyHeaderWithoutForge(chain ChainReader, header Header) error

	// VerifySeal checks whether the crypto seal on a header is valid according to
	// the consensus rules of the given engine.
	VerifyForge(chain ChainReader, header Header) error
}

type Forger interface {
	Forge(chain ChainReader, header Header) (Data, error)
}
