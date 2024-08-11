package services

import (
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/request"
	"pledge-backend-practise/api/models/response"
)

type MultiSign struct {
}

func NewMultiSign() *MultiSign {
	return &MultiSign{}
}

func (m *MultiSign) SetMultiSign(req *request.SetMultiSign) int {
	// TODO: db operation.
	return statuscode.CommonSuccess
}

func (m *MultiSign) GetMultiSign(req *request.GetMultiSign, rsp *response.MultiSign) int {
	// TODO: db operation
	return statuscode.CommonSuccess
}
