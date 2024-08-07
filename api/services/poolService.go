package services

import (
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/response"
	"pledge-backend-practise/log"
)

type poolService struct{}

func NewPool() *poolService {
	return &poolService{}
}

func (s *poolService) PoolBaseInfo(chainId int, res *[]response.PoolBaseInfoRes) int {
	err := response.NewPoolBases().PoolBaseInfo(chainId, res)
	if err != nil {
		log.Logger.Error(err.Error())
		return statuscode.CommonErrServerErr
	}
	return statuscode.CommonSuccess
}

func (s *poolService) PoolDataInfo(chainId int, res *[]response.PoolDataRes) int {
	err := response.NewPoolData().PoolDataInfo(chainId, res)
	if err != nil {
		log.Logger.Error(err.Error())
		return statuscode.CommonErrServerErr
	}

	return statuscode.CommonSuccess
}
