package main

import (
	"log"
	"service"
	"service/visistservice"
)

func main() {
	visistservice.Init()
	err := service.Run(&service.ServiceInfo{
		Name:      "visist",
		Addr:      "127.0.0.1:20003",
		RequiredServices: []string{"log"},
	})
	if err != nil {
		log.Fatalln(err)
	}
}
