package ws

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"pledge-backend-practise/api/models/kucoin"
	"pledge-backend-practise/config"
	"pledge-backend-practise/log"
	"sync"
	"time"
)

const SuccessCode = 0
const PongCode = 1
const ErrorCode = -1

type Service struct {
	sync.Mutex
	Id       string
	Socket   *websocket.Conn
	Send     chan []byte
	LastTime int64 // last send time
}

type ServerManager struct {
	Servers    sync.Map
	Broadcast  chan []byte
	Register   chan *Service
	UnRegister chan *Service
}

type Message struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

var manager = ServerManager{}
var UserPingPongDurTime = config.Config.Env.WssTimeoutDuration

func (s *Service) SendToClient(code int, data string) {
	s.Lock()
	defer s.Unlock()

	bytes, err := json.Marshal(Message{Code: code, Data: data})
	if err != nil {
		log.Logger.Sugar().Error(s.Id+" SendToClient json marshal err ", err)
		return
	}

	err = s.Socket.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		log.Logger.Sugar().Error(s.Id+" SendToClient write message err ", err)
	}
}

func (s *Service) ReadAndWrite() {
	errChan := make(chan error)

	manager.Servers.Store(s.Id, s)
	defer func() {
		manager.Servers.Delete(s.Id)
		_ = s.Socket.Close()
		close(s.Send)
	}()

	//write
	go func() {
		for {
			select {
			case messageBytes, ok := <-s.Send:
				if !ok {
					errChan <- errors.New("write message error")
					return
				}
				s.SendToClient(SuccessCode, string(messageBytes))
			}
		}

	}()

	//read
	go func() {
		for {
			_, messageBtyes, err := s.Socket.ReadMessage()
			if err != nil {
				log.Logger.Sugar().Error(s.Id+" ReadMessage err ", err)
				errChan <- err
				return
			}

			//update heartbeat time
			if string(messageBtyes) == "ping" {
				s.LastTime = time.Now().Unix()
				s.SendToClient(PongCode, "pong")
				return
			}

			s.LastTime = time.Now().Unix()
			s.Send <- messageBtyes
		}
	}()

	//check heartbeat
	go func() {
		select {
		case <-time.After(time.Second):
			if time.Now().Unix()-s.LastTime > UserPingPongDurTime {
				s.SendToClient(ErrorCode, "heartbeat timeout")
				return
			}
		case err := <-errChan:
			log.Logger.Sugar().Error(s.Id, " ReadAndWrite returned ", err)
			return
		}
	}()
}

func StartServer() {
	log.Logger.Info("start ws server.")
	for {
		select {
		case price, ok := <-kucoin.PlgrPriceChan:
			if ok {
				manager.Servers.Range(func(k, v interface{}) bool {
					v.(*Service).SendToClient(SuccessCode, price)
					return true
				})
			}
		}
	}
}
