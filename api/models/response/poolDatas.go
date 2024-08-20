package response

import "pledge-backend-practise/db"

type PoolData struct {
	PoolID                 int    `json:"-"`
	SettleAmountLend       string `json:"settleAmountLend"`
	SettleAmountBorrow     string `json:"settleAmountBorrow"`
	FinishAmountLend       string `json:"finishAmountLend"`
	FinishAmountBorrow     string `json:"finishAmountBorrow"`
	LiquidationAmounLend   string `json:"liquidationAmounLend"`
	LiquidationAmounBorrow string `json:"liquidationAmounBorrow"`
}

type PoolDataRes struct {
	Index    int      `json:"index"`
	PoolData PoolData `json:"pool_data"`
}

func NewPoolData() *PoolData {
	return &PoolData{}
}

func (data *PoolData) PoolDataInfo(chainId int, res *[]PoolDataRes) error {
	var poolData []PoolData

	err := db.MySql.Table("pooldata").Where("chain_id = ?", chainId).Order("pool_id asc").Find(&poolData).Debug().Error
	if err != nil {
		return err
	}

	for _, v := range poolData {
		*res = append(*res, PoolDataRes{
			Index:    v.PoolID - 1,
			PoolData: v,
		})
	}
	return nil
}
