package response

import "pledge-backend-practise/api/models/request"

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
	pools := make([]Pool, 0)

	// TODO : query db and fill up data.
	return nil, total, pools
}
