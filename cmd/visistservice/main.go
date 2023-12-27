package main

import (
	"log"
	"service"
	"service/info"
	"service/visistservice"
)

func main() {
	visistservice.Init()
	engine := service.New(&info.ServiceInfo{
		Name:      "visist",
		Addr:      "127.0.0.1:20003",
		RequiredServices: []string{"log"},
	})

	err := engine.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
