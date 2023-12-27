package main

import (
	"log"
	"service/gatewayservice"
	"service"
	"service/info"
)

func main() {
	gatewayservice.Init()
	engine := service.New(&info.ServiceInfo{
		Name: "gateway",
		Addr: "127.0.0.1:20001",
		RequiredServices: []string{"log", "visist"},
	})

	
	err := engine.Run()

	if err != nil {
		log.Fatalln(err)
	}
}