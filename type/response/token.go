package response

type Token struct {
	Name string
	Id []byte
	Introduction string

	Symbol string

	Owner []byte

	Decimals int

	TotalSupply string

	TransactionNumber string
}