package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"pledge-backend-practise/api/models/ws"
	"pledge-backend-practise/log"
	"pledge-backend-practise/utils"
	"strings"
	"time"
)

type PriceController struct{}

func NewPriceController() *PriceController {
	return &PriceController{}
}

func (p *PriceController) NewPrice(ctx *gin.Context) {
	defer func() {
		if recoverRes := recover(); recoverRes != nil {
			log.Logger.Sugar().Error("new price recover ", recoverRes)
		}
	}()

	conn, err := (&websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		CheckOrigin:      func(r *http.Request) bool { return true },
		HandshakeTimeout: 5 * time.Second,
	}).Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Logger.Sugar().Error("websocket request err:", err)
		return
	}

	randomId := ""
	remoteIp := ctx.RemoteIP()

	if len(remoteIp) > 0 {
		randomId = strings.ReplaceAll(remoteIp, ".", "_") + "_" + utils.GetRandomString(23)
	} else {
		randomId = utils.GetRandomString(32)
	}

	s := &ws.Service{
		Id:       randomId,
		Socket:   conn,
		Send:     make(chan []byte, 256),
		LastTime: time.Now().Unix(),
	}

	go s.ReadAndWrite()
}
