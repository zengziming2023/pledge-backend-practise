package services

import (
	"fmt"
	"pledge-backend-practise/api/common/statuscode"
	"pledge-backend-practise/api/models/request"
	"pledge-backend-practise/api/models/response"
)

type Search struct {
}

func NewSearch() *Search {
	return &Search{}
}

func (s *Search) Search(req *request.Search, rsp *response.Search) int {
	whereCondition := fmt.Sprintf(`chain_id='%v'`, req.ChainID)
	if req.LendTokenSymbol != "" {
		whereCondition += fmt.Sprintf(` and lend_token_symbol='%v'`, req.LendTokenSymbol)
	}
	if req.State != "" {
		whereCondition += fmt.Sprintf(` and state='%v'`, req.State)
	}

	err, count, pools := response.NewPool().Pagination(req, whereCondition)
	if err != nil {
		return statuscode.CommonErrServerErr
	}

	rsp.Count = count
	rsp.Rows = pools

	return statuscode.CommonSuccess
}
