package request

type SnapshotChainBlocklist struct {
	Paging Paging `form:"paging" binding:"required"`
}

type SnapshotChainBlock struct {
	BlockHash string `form:"blockHash" binding:"required"`
}
