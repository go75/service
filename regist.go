package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"service/packet"
	"service/info"
)

func (e *Engine) registerMonitor() {
	msgChan := rds.Subscribe(e.serviceInfo.Name).Channel()
	go func() {
		for msg := range msgChan {
			msg := packet.UnPack([]byte(msg.String()))
			e.RouteTable.Process(msg)
		}
	}()

	http.HandleFunc("/heart-beat", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func RegistService(serviceInfo *info.ServiceInfo) error {
	data, err := json.Marshal(serviceInfo)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://"+SERVICE_ADDR+"/services", "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("regist %s error with code %d", serviceInfo.Name, resp.StatusCode)
	}

	err = provider.ParseServiceInfos(resp.Body)
	if err != nil {
		return err
	}

	provider.Dump()

	data = packet.Pack(packet.NewMessage(packet.ADD, data))

	err = rds.Publish(serviceInfo.Name, data).Err()

	return err
}

func UnregistService(serviceInfo *info.ServiceInfo) error {
	data, err := packet.JsonMarshal(packet.ADD, serviceInfo)
	if err != nil {
		return err
	}

	err = rds.Publish(serviceInfo.Name, data).Err()

	return err
}

var provider = info.NewServiceTable()

func Get(serviceName string) string {
	service := provider.Get(serviceName)
	if service != nil {
		return service.Addr
	}
	return ""
}
