package services

import (
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models"
	"pledge-backend-practise/log"
)

type poolService struct{}

func NewPool() *poolService {
	return &poolService{}
}

func (s *poolService) PoolBaseInfo(chainId int, res *[]models.PoolBaseInfoRes) int {
	err := models.NewPoolBases().PoolBaseInfo(chainId, res)
	if err != nil {
		log.Logger.Error(err.Error())
		return statuscode.CommonErrServerErr
	}
	return statuscode.CommonSuccess
}
