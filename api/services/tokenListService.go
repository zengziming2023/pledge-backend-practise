package services

import (
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models"
	"pledge-backend-practise/api/models/request"
	"pledge-backend-practise/api/models/response"
	"pledge-backend-practise/config"
)

type TokenList struct {
}

func NewTokenList() *TokenList {
	return &TokenList{}
}

func (t *TokenList) TokenList(req *request.TokenList, rsp *response.TokenList) int {
	err, tokenList := models.NewTokenInfo().GetTokenList(req)
	if err != nil {
		return statuscode.CommonErrServerErr
	}

	rsp.Name = "Pledge Token List"
	rsp.LogoURI = config.GetBaseUrl() + "storage/img/Pledge-project-logo.png"
	rsp.Version = response.Version{
		Major: 2,
		Minor: 16,
		Patch: 12,
	}

	for _, v := range tokenList {
		rsp.Tokens = append(rsp.Tokens, response.Token{
			Name:     v.Symbol,
			Symbol:   v.Symbol,
			Decimals: v.Decimals,
			Address:  v.Token,
			ChainID:  v.ChainId,
			LogoURI:  v.Logo,
		})
	}

	return statuscode.CommonSuccess
}

func (t *TokenList) DebtTokenList(req *request.TokenList, rsp *[]models.TokenInfo) int {
	err, tokenInfos := models.NewTokenInfo().DebtTokenList(req)
	if err != nil {
		return statuscode.CommonErrServerErr
	}

	for _, v := range tokenInfos {
		*rsp = append(*rsp, v)
	}
	return statuscode.CommonSuccess
}
