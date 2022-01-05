package httpserver

import (
	"context"
	"grpcproto/hander"
	"grpcproto/services"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func Gw() {
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(hander.GetClientCreds())}
	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:8081", opt)
	if err != nil {
		log.Fatal(err)
	}

	err = services.RegisterOrderServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:8081", opt)
	if err != nil {
		log.Fatal(err)
	}

	httpserver := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	httpserver.ListenAndServe()
}
