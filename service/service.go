package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, serviceName, host, port string, registerHandler func()) (context.Context, error) {
	registerHandler()
	ctx, err := startService(ctx, serviceName, host, port)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) (context.Context, error) {

	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println(serviceName, "started. Press any key to stop.")
		var s string
		_, _ = fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx, nil
}
