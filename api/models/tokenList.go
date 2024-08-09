package models

import "pledge-backend-practise/api/models/request"

type TokenInfo struct {
	Id      int32  `json:"-" gorm:"column:id;primaryKey"`
	Symbol  string `json:"symbol" gorm:"column:symbol"`
	Token   string `json:"token" gorm:"column:token"`
	ChainId int    `json:"chain_id" gorm:"column:chain_id"`
}

type TokenList struct {
	Id       int32  `json:"-" gorm:"column:id;primaryKey"`
	Symbol   string `json:"symbol" gorm:"column:symbol"`
	Decimals int    `json:"decimals" gorm:"column:decimals"`
	Token    string `json:"token" gorm:"column:token"`
	Logo     string `json:"logo" gorm:"column:logo"`
	ChainId  int    `json:"chain_id" gorm:"column:chain_id"`
}

func NewTokenInfo() *TokenInfo {
	return &TokenInfo{}
}

func (t *TokenInfo) GetTokenList(req *request.TokenList) (error, []TokenList) {
	tokenList := make([]TokenList, 0)
	//err := db.Mysql.Table("token_info").Where("chain_id", req.ChainId).Find(&tokenList).Debug().Error

	// query from db
	return nil, tokenList
}
