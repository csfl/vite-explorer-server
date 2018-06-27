package response

import "math/big"

type SnapshotBlock struct {
	PrevHash []byte
	Height *big.Int
	Producer []byte
	Snapshot map[string][]byte
	Amount *big.Int
}