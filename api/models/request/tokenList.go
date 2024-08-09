package request

type TokenList struct {
	ChainId int `json:"chainId" binding:"required"`
}
