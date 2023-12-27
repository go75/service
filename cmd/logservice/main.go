package main

import (
	"log"
	"service"
	"service/info"
	"service/logservice"
)

func main() {
	logservice.Init("./services.log")
	engine := service.New(&info.ServiceInfo{
		Name:      "log",
		Addr:      "127.0.0.1:20002",
		RequiredServices: make([]string, 0),
	})

	err := engine.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
