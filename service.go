package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"service/config"
	"service/info"
	"service/route"
)

type Engine struct {
	serviceInfo *info.ServiceInfo
	route.RouteTable
}

func New(serviceInfo *info.ServiceInfo) *Engine {
	return &Engine{
		serviceInfo: serviceInfo,
		RouteTable: make(route.RouteTable, 0),
	}
}

func (e *Engine) Run() (err error) {
	connectRedis(config.REDIS_ADDR, config.REDIS_PASS, config.REDIS_DB)
	e.registerMonitor()

	err = RegistService(e.serviceInfo)
	if err != nil {
		return err
	}

	defer func() {
		err = errors.Join(err, UnregistService(e.serviceInfo))
	}()

	server := http.Server{Addr: e.serviceInfo.Addr}

	go func() {
		fmt.Println("Press any key to stop.")
		var s string
		fmt.Scan(&s)
		server.Shutdown(context.Background())
	}()

	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
