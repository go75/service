package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"service/config"
)

type Service interface {
	Init()
}

func Run(service *ServiceInfo) (err error) {
	connectRedis(config.REDIS_ADDR, config.REDIS_PASS, config.REDIS_DB)
	
	err = RegistService(service)
	if err != nil {
		return err
	}
	defer func() {
		err = errors.Join(err, UnregistService(service))
	}()

	srv := http.Server{Addr: service.Addr}

	go func() {
		fmt.Println("Press any key to stop.")
		var s string
		fmt.Scan(&s)
		srv.Shutdown(context.Background())
	}()

	err = srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
