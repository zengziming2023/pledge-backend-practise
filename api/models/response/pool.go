package response

import (
	"encoding/json"
	"pledge-backend-practise/api/models/request"
	"pledge-backend-practise/db"
	"pledge-backend-practise/db/models"
)

type Pool struct {
	PoolID                 int      `json:"pool_id"`
	SettleTime             string   `json:"settleTime"`
	EndTime                string   `json:"endTime"`
	InterestRate           string   `json:"interestRate"`
	MaxSupply              string   `json:"maxSupply"`
	LendSupply             string   `json:"lendSupply"`
	BorrowSupply           string   `json:"borrowSupply"`
	MartgageRate           string   `json:"martgageRate"`
	LendToken              string   `json:"lendToken"`
	LendTokenSymbol        string   `json:"lend_token_symbol"`
	BorrowToken            string   `json:"borrowToken"`
	BorrowTokenSymbol      string   `json:"borrow_token_symbol"`
	State                  string   `json:"state"`
	SpCoin                 string   `json:"spCoin"`
	JpCoin                 string   `json:"jpCoin"`
	AutoLiquidateThreshold string   `json:"autoLiquidateThreshold"`
	Pooldata               PoolData `json:"pooldata"`
}

func NewPool() *Pool {
	return &Pool{}
}

func (p *Pool) Pagination(req *request.Search, whereCondition string) (error, int64, []Pool) {
	var total int64
	poolBases := make([]models.PoolBase, 0)

	pools := make([]Pool, 0)

	// query db and fill up data.
	db.MySql.Table("poolbases").Where(whereCondition).Count(&total)

	err := db.MySql.Table("poolbases").Where(whereCondition).Order("pool_id desc").Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&poolBases).Debug().Error
	if err != nil {
		return err, 0, pools
	}

	poolData := PoolData{}
	err = db.MySql.Table("pooldata").Where("chain_id=?", req.ChainID).First(&poolData).Debug().Error
	if err != nil {
		return err, 0, pools
	}

	for _, poolBase := range poolBases {
		var lendToken LendTokenInfo
		_ = json.Unmarshal([]byte(poolBase.LendToken), &lendToken)
		var borrowToken BorrowTokenInfo
		_ = json.Unmarshal([]byte(poolBase.BorrowToken), &borrowToken)

		pools = append(pools, Pool{
			PoolID:                 poolBase.PoolId,
			SettleTime:             poolBase.SettleTime,
			EndTime:                poolBase.EndTime,
			InterestRate:           poolBase.InterestRate,
			MaxSupply:              poolBase.MaxSupply,
			LendSupply:             poolBase.LendSupply,
			BorrowSupply:           poolBase.BorrowSupply,
			MartgageRate:           poolBase.MartgageRate,
			LendToken:              lendToken.TokenName,
			BorrowToken:            borrowToken.TokenName,
			BorrowTokenSymbol:      poolBase.BorrowTokenSymbol,
			State:                  poolBase.State,
			SpCoin:                 poolBase.SpCoin,
			JpCoin:                 poolBase.JpCoin,
			AutoLiquidateThreshold: poolBase.AutoLiquidateThreshold,
			Pooldata:               poolData,
		})
	}

	return nil, total, pools
}
