package main

import (
	"go.uber.org/dig"

	"kbtg.tech/template/httpsvr"
	"kbtg.tech/template/services"
	"kbtg.tech/template/share"
	golog "log"
	"os"
	"os/signal"
	"runtime"
)

var buildstamp string
var githash string

func NewConfig() *share.Config {
	return &share.Config{
		HttpPort:   "8000",
		SocketPort: "27305",
	}
}

func NewHttpServer(conf *share.Config) *httpsvr.HttpServer {
	return &httpsvr.HttpServer{
		Config: conf,
	}
}
func NewSocketService(conf *share.Config) *services.SocketService {
	return &services.SocketService{
		Filename: "",
		Filepath: "",
	}
}
//func NewSocketServer(conf *share.Config, svc *services.SocketService) *socketsvr.SocketServer {
//	return &socketsvr.SocketServer{
//		Config:        conf,
//		SocketService: svc,
//	}
//}

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(NewConfig)
	container.Provide(NewHttpServer)
	container.Provide(NewSocketService)
	//container.Provide(NewSocketServer)
	return container
}

func main() {
	runtime.GOMAXPROCS(1)
	container := BuildContainer()
	//	err := container.Invoke(func(server *socketsvr.SocketServer) {
	//		server.Run()
	//	})
	//	if err != nil {
	//		panic(err)
	//	}
	err := container.Invoke(func(server *httpsvr.HttpServer) {
		server.Run()
	})
	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	golog.Println("Server exiting")
}
