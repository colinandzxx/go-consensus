package consensus

import (
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/tinylib/msgp/msgp"
	"math/big"
)

// ================= chain interface =================

type Header interface {
	//Marshal() ([]byte, error)

	GetHeight() uint64
	GetHash() Byte32
	GetParentHash() Byte32
	GetNonce() uint64
	GetTimestamp() uint64
	GetCoinbase() Bytes
	GetConsensusData() Data
}

type Block interface{
	//Marshal() ([]byte, error)

	GetHeader() Header
	GetTotalDifficulty() *big.Int
}

type ChainReader interface {
	GetCurrentHeader() Header
	GetHeader(hash Byte32, height uint64) Header
	GetHeaderByHeight(height uint64) Header
	GetHeaderByHash(hash Byte32) Header
	GetBlock(hash Byte32, height uint64) Block
}

// ================= consensus data interface =================

type DataWrapper interface {
	Wrap(chain ChainReader, height uint64) ([]byte, error)
}

type DataUnWrapper interface {
	UnWrap([]byte) (Data, error)
}

type Data interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
	json.Unmarshaler
	json.Marshaler
	fmt.Stringer
	fmt.Formatter

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
