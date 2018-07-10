package request

type SnapshotChainBlocklist struct {
	Paging
}

type SnapshotChainBlock struct {
	BlockHash string `form:"blockHash" binding:"required"`
}
