package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"lottery/utils/common"
)

var NotifyChanMap map[string]chan interface{}
var NotifyChanMutex     sync.RWMutex

func init() {
	NotifyChanMap = make(map[string]chan interface{})
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NotifySocket(r *gin.Context) {
	ws,err := upGrader.Upgrade(r.Writer,r.Request,nil)
	if err != nil {
		r.JSON(http.StatusBadRequest,"upgrade websocket failed")
		return
	}
	defer ws.Close()
	notifyChan := make(chan interface{}, 1000)
	key := common.RandStringBytesMaskImprSrc(8)
	NotifyChanMap[key] = notifyChan
	for {
		data := <- notifyChan
		err = ws.WriteJSON(data)
		if err != nil {
			close(notifyChan)
			break
		}
	}
	NotifyChanMutex.Lock()
	defer NotifyChanMutex.Unlock()
	delete(NotifyChanMap,key)
}

func DispatchData(i interface{}) {
	NotifyChanMutex.RLock()
	defer NotifyChanMutex.RUnlock()
	for _,c := range NotifyChanMap {
		c <- i
	}
}