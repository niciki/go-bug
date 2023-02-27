package main

import (
	"fmt"
	"os"
	"os/signal"
	"service/internal/service"
	"service/internal/transport"

	"github.com/rs/zerolog/log"
)

func main() {
	service := service.FouService{}
	servicesOpts := []transport.Option{
		transport.Service(transport.NewFouService(service)),
	}

	srv := transport.New(log.Logger, servicesOpts...).WithLog()

	go func() {
		if err := srv.Fiber().Listen(":8080"); err != nil {
			fmt.Println(err.Error())
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
