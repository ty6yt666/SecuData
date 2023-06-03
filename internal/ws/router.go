package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func InitRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {

}

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var dataChan = make(chan map[string]interface{}, 100)

func installController(g *gin.Engine) {
	//storeIns, _ := postgres.GetPostgresFactorOr(pkg_options.NewPostgresOptions())
	v1 := g.Group("v1")
	{
		wsInstrumentv1 := v1.Group("/ws")
		{
			//wsInstrumentController := instrument.NewInstrumentController(storeIns)
			wsInstrumentv1.GET("/market_data", func(c *gin.Context) {
				// 升级成ws连接
				ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
				if err != nil {
					log.Fatal(err)
				}
				// 完成时关闭连接释放
				defer ws.Close()
				go func() {
					<-c.Done()
					fmt.Println("ws lost connection")
				}()

				go func() {
					for i := 0; ; i++ {
						data := map[string]interface{}{
							"msg":      i,
							"meg_type": "ooo",
						}
						dataChan <- data
					}
				}()

				// 单工模式，有数据就推送
				for {
					message, _ := json.Marshal(<-dataChan)
					time.Sleep(300 * time.Microsecond)
					ws.WriteMessage(1, message)
				}

				// 双工模式，死循环监听，读取客户端发来的消息，如果没收到就会一直阻塞
				//for {
				//	// messageType, message
				//	mt, message, err := ws.ReadMessage()
				//	if err != nil {
				//		fmt.Println("read error")
				//		fmt.Println(err)
				//		break
				//	}
				//	switch string(message) {
				//	case "ping":
				//		message = []byte("pong")
				//	case "instrument":
				//		message = []byte("进到instrument")
				//	case "market_data":
				//		message = []byte("进到market_data")
				//	default:
				//		{
				//		}
				//	}
				//	err = ws.WriteMessage(mt, message)
				//	if err != nil {
				//		fmt.Println(err)
				//		break
				//	}
				//}
			})
		}
	}
}
