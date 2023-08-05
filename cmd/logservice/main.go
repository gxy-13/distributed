package main

import (
	"context"
	"distributed/log"
	"distributed/service"
	"fmt"
	stdlog "log"
)

func main() {
	// Run before start service to initialize log
	log.Run("./distributed.log")
	ctx, err := service.Start(context.Background(),
		"Log Service",
		"localhost",
		"4000",
		log.RegisterHandlers)
	if err != nil {
		stdlog.Println(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
