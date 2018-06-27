package response

type SnapshotBlock struct {
	PrevHash []byte
	Height string
	Producer []byte
	Snapshot map[string][]byte
	Amount string
}