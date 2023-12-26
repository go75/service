package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"service/packet"
)

func registerMonitorHandler() {
	http.HandleFunc("/heart-beat", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func RegistService(service *ServiceInfo) error {
	registerMonitorHandler()

	data, err := json.Marshal(service)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://"+SERVICE_ADDR+"/services", "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("regist %s error with code %d", service.Name, resp.StatusCode)
	}

	err = provider.parseServiceInfos(resp.Body)
	if err != nil {
		return err
	}

	provider.dump()

	data, err = packet.JsonMarshal(packet.ADD, service)
	if err != nil {
		return err
	}

	err = rds.Publish(service.Name, data).Err()

	return err
}

func UnregistService(service *ServiceInfo) error {
	data, err := packet.JsonMarshal(packet.ADD, service)
	if err != nil {
		return err
	}

	err = rds.Publish(service.Name, data).Err()

	return err
}

var provider = newServiceTable()

func Get(serviceName string) string {
	service := provider.get(serviceName)
	if service != nil {
		return service.Addr
	}
	return ""
}
