package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pledge-backend-practise/api/common/statuscode"
)

type PageResponse struct {
	Code  int         `json:"code"`
	Msg   string      `json:"message"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func JsonResponse(c *gin.Context, code int, data interface{}, httpStatusArray ...int) {
	lang := statuscode.LangEn
	if clang, ok := c.Get("lang"); ok {
		lang = clang.(int)
	}

	httpStatus := http.StatusOK
	if len(httpStatusArray) > 0 {
		httpStatus = httpStatusArray[0]
	}

	c.JSON(httpStatus, Response{
		Code: code,
		Msg:  statuscode.GetMsg(code, lang),
		Data: data,
	})
}

func JsonPagedResponse(c *gin.Context, code int, totalCount int, data interface{}, httpStatusArray ...int) {
	lang := statuscode.LangEn
	if clang, ok := c.Get("lang"); ok {
		lang = clang.(int)
	}

	httpStatus := http.StatusOK
	if len(httpStatusArray) > 0 {
		httpStatus = httpStatusArray[0]
	}

	c.JSON(httpStatus, PageResponse{
		Code:  code,
		Msg:   statuscode.GetMsg(code, lang),
		Total: totalCount,
		Data:  data,
	})

}
