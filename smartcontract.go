package sc

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"encoding/json"
)

type StartContract struct{
	Action string `json:"action"`
	Contractid string `json:"contractid"`
	Path string `json:"path"`
	Script string `json:"script"`
	Type string `json:"type"`
	Onwer string `json:"onwer"`
}

type ExecuteContract struct {
	Action string `json:"action"`
	ContractID string `json:"contractID"`
	Arg string `json:"arg"`
	Requester string `json:"requester"`
}

type KillContractProcess struct {
	Action string `json:"action"`
	Id string  `json:"id"`
}

type Response struct{
	Action string `json:"action"`
	Data string `json:"data"`
	ExecuteTime int64 `json:"executeTime"`
}

type Handle interface {
	OnStartContract(res Response) error
	OnExecuteResult(res Response) error
	OnListContractProcess(res Response) error
	OnKillContractProcess(res Response) error
	OnOutputStream(res Response) error
}

type WebSocketClient struct {
	conn *websocket.Conn
	URL  url.URL
	Handler  Handle
}

// example:
// addr: localhost:8080
// path: /SCIDE/SCExecutor
/*func InitWebSocket(addr string, path string) WebSocketClient {
	u := url.URL{Scheme: "ws", Host: addr, Path: path}

	wsClient := WebSocketClient{
		url: u,
	}

	return wsClient
}*/

func (wsc *WebSocketClient) Start() {
	log.Printf("connecting to %s", wsc.URL.String())

	var err error
	wsc.conn, _, err = websocket.DefaultDialer.Dial(wsc.URL.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	// recieve message loop
	go func() {
		for {
			_, message, err := wsc.conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
			go func() {
				if handler := wsc.Handler; handler != nil {
					response := Response{}
					json.Unmarshal([]byte(message), &response)
					switch response.Action {
					case "onStartContract":
						handler.OnStartContract(response)
					case "onExecuteResult":
						handler.OnExecuteResult(response)
					case "onListContractProcess":
						handler.OnListContractProcess(response)
					case "onKillContractProcess":
						handler.OnKillContractProcess(response)
					case "onOutputStream":
						handler.OnOutputStream(response)
					}
				}
			}()
		}
	}()
}

func (wsc *WebSocketClient) Send(data interface{}) (err error) {
	jsons, err := json.Marshal(data)
	log.Println("send message: \n", string(jsons))
	err = wsc.conn.WriteMessage(websocket.TextMessage, []byte(jsons))
	if err != nil {
		log.Println("write:", err)
		return
	}
	return
}

func (wsc *WebSocketClient) Close() {
	wsc.conn.Close()
}

func (wsc *WebSocketClient) Handle(handle Handle) error {
	wsc.Handler = handle
	return nil
}