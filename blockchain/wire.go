package blockchain

import (
	"github.com/tendermint/go-amino"
	"github.com/blockchainworkers/conch/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockchainMessages(cdc)
	types.RegisterBlockAmino(cdc)
}
