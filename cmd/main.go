package main


import (
	"context"
	"fmt"
	"gin-demo/internal/dao"
	"gin-demo/internal/models/handler"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"

	"gin-demo/internal/configs"
	"gin-demo/internal/routers"

)


func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	configs.Init()
	dao.Init()
	// 创建一个通道，用于在协程之间传递消息
	msgChan := make(chan string)

	// 创建一个消息处理函数，将接收到的消息发送到通道
	f := func(client mqtt.Client, msg mqtt.Message) {
		msgChan <- string(msg.Payload())
	}

	// 创建一个客户端
	opts := mqtt.NewClientOptions().AddBroker(configs.Cfg.Mqtt.Url).SetClientID("gotrivial")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	if token := c.Subscribe("test", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// set a new goroutine to receive message
	go func() {
		for msg := range msgChan {
			fmt.Println("Received message:", msg)
			// 在这里处理消息
		}
	}()

	// set router and start http server
	r := gin.Default()
	routers.SetupRouter(r)

	<-ctx.Done()
	stop()
	log.Infoln("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// close handler
	handler.Close()

}
