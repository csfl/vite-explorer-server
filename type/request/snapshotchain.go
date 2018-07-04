package request

type SnapshotChainBlocklist struct {
	Paging
}

type SnapshotChainBlock struct {
	BlockHash string `form:"blockHahs" binding:"required"`
}
