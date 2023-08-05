package main

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	"fmt"
	stdlog "log"
)

func main() {
	// Run before start service to initialize log
	log.Run("./distributed.log")
	r := registry.Registration{
		ServiceName: registry.LogService,
		ServiceURL:  "http://localhost:4000",
	}
	ctx, err := service.Start(context.Background(),
		"localhost",
		"4000",
		r,
		log.RegisterHandlers)
	if err != nil {
		stdlog.Println(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
