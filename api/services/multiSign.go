package services

import (
	"encoding/json"
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/request"
	"pledge-backend-practise/api/models/response"
	"pledge-backend-practise/db"
)

type MultiSign struct {
}

func NewMultiSign() *MultiSign {
	return &MultiSign{}
}

func (m *MultiSign) SetMultiSign(req *request.SetMultiSign) int {
	// db operation.
	dbMultiSign := &response.DbMultiSign{}
	err := db.MySql.Table("multi_sign").Where("chain_id =? ", req.ChainId).First(dbMultiSign).Debug().Error
	if err != nil {
		return statuscode.CommonErrServerErr
	}

	multiSignAccountByteArr, _ := json.Marshal(req.MultiSignAccount)
	err = db.MySql.Table("multi_sign").Where("id=?", dbMultiSign.Id).Create(&response.DbMultiSign{
		SpName:           req.SpName,
		ChainId:          req.ChainId,
		SpToken:          req.SpToken,
		JpName:           req.JpName,
		JpToken:          req.JpToken,
		SpAddress:        req.SpAddress,
		JpAddress:        req.JpAddress,
		SpHash:           req.SpHash,
		JpHash:           req.JpHash,
		MultiSignAccount: string(multiSignAccountByteArr),
	}).Debug().Error
	if err != nil {
		return statuscode.CommonErrServerErr
	}

	return statuscode.CommonSuccess
}

func (m *MultiSign) GetMultiSign(req *request.GetMultiSign, rsp *response.MultiSign) int {
	// db operation
	dbMultiSign := &response.DbMultiSign{}
	err := db.MySql.Table("multi_sign").Where("chain_id =?", req.ChainId).First(dbMultiSign).Debug().Error
	if err != nil {
		return statuscode.CommonErrServerErr
	}

	var multiSignAccount []string
	_ = json.Unmarshal([]byte(dbMultiSign.MultiSignAccount), &multiSignAccount)

	rsp.SpName = dbMultiSign.SpName
	rsp.SpToken = dbMultiSign.SpToken
	rsp.JpName = dbMultiSign.JpName
	rsp.JpToken = dbMultiSign.JpToken
	rsp.SpAddress = dbMultiSign.SpAddress
	rsp.JpAddress = dbMultiSign.JpAddress
	rsp.SpHash = dbMultiSign.SpHash
	rsp.JpHash = dbMultiSign.JpHash
	rsp.MultiSignAccount = multiSignAccount

	return statuscode.CommonSuccess
}
