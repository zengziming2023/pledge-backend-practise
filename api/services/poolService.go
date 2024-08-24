package services

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/response"
	"pledge-backend-practise/config"
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

func (s *poolService) UpdateAllPoolInfo() {
	testNet := config.Config.TestNet
	s.UpdatePoolInfo(testNet.PlgrAddress, testNet.NetUrl, testNet.ChainId)

	mainNet := config.Config.MainNet
	s.UpdatePoolInfo(mainNet.PlgrAddress, mainNet.NetUrl, mainNet.ChainId)
}

func (s *poolService) UpdatePoolInfo(contractAddress string, netUrl string, chainId string) {
	log.Logger.Info("update pool info" + contractAddress + ", " + netUrl + ", " + chainId)
	client, err := ethclient.Dial(netUrl)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	_ = client
}
