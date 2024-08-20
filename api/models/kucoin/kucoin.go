package kucoin

import (
	"github.com/Kucoin/kucoin-go-sdk"
	"pledge-backend-practise/db"
	"pledge-backend-practise/log"
)

// ApiKeyVersionV2 is v2 api key version
const ApiKeyVersionV2 = "2"

var PlgrPrice = "0.0027"
var PlgrPriceChan = make(chan string, 2)

func GetExchangePrice() {
	log.Logger.Info("GetExchangePrice")

	// get plgr price from redis
	price, err := db.RedisGetString("plgr_price")
	if err != nil {
		log.Logger.Error("get price from redis error: " + err.Error())
	} else {
		PlgrPrice = price
	}

	s := kucoin.NewApiService(kucoin.ApiKeyOption("key"),
		kucoin.ApiSecretOption("secret"),
		kucoin.ApiPassPhraseOption("passphrase"),
		kucoin.ApiKeyVersionOption(ApiKeyVersionV2))

	tokenResp, err := s.WebSocketPublicToken()
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	token := &kucoin.WebSocketTokenModel{}
	err = tokenResp.ReadData(token)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	socketClient := s.NewWebSocketClient(token)
	socketMsgChan, socketErrChan, err := socketClient.Connect()
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	subscribMsg := kucoin.NewSubscribeMessage("/market/ticker:PLGR-USDT", false)
	unSubcribeMsg := kucoin.NewUnsubscribeMessage("/market/ticker:PLGR-USDT", false)

	err = socketClient.Subscribe(subscribMsg)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	select {
	case err := <-socketErrChan:
		if err != nil {
			log.Logger.Error(err.Error())
			socketClient.Stop()
			_ = socketClient.Unsubscribe(unSubcribeMsg)
			return
		}

	case msg := <-socketMsgChan:
		t := &kucoin.TickerLevel1Model{}
		err := msg.ReadData(t)
		if err != nil {
			log.Logger.Error(err.Error())
			return
		}
		PlgrPrice = t.Price
		PlgrPriceChan <- PlgrPrice

		// 更新 redis
		_ = db.RedisSetString("plgr_price", PlgrPrice, 0)
	}
}
