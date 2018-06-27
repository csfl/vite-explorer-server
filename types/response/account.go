package response

type AccountToken struct {
	Balance string
}

type Account struct {
	AccountAddress string

	BlockHeight string

	TokenList []*AccountToken
}