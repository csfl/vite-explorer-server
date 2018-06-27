package response

type AccountBlock struct {
	Height string

	AccountAddress string

	To []byte

	FromHash []byte

	PrevHash []byte

	Status int

	Balance string

	Amount string

	Data string

	SnapshotTimestamp []byte

	Signature []byte

	Nounce []byte

	Difficulty []byte

	FAmount []byte
}