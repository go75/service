package main

import (
	"log"
	"service/gatewayservice"
	"service"
)

func main() {
	gatewayservice.Init()
	err := service.Run(&service.ServiceInfo{
		Name: "gateway",
		Addr: "127.0.0.1:20001",
		RequiredServices: []string{"log", "visist"},
	})

	if err != nil {
		log.Fatalln(err)
	}
}