package main

import (
	"grpcproto/hander"
	"grpcproto/httpserver"
	"grpcproto/services"
	"net"

	"google.golang.org/grpc"
)

func main() {

	// creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server.key")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// rpcServer := grpc.NewServer(grpc.Creds(hander.GetClientCreds()), grpc.UnaryInterceptor(interceptor.Interceptor))
	rpcServer := grpc.NewServer(grpc.Creds(hander.GetClientCreds()))

	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))  //商品服务
	services.RegisterOrderServiceServer(rpcServer, new(services.OderService)) //订单服务
	services.RegisterUserServiceServer(rpcServer, new(services.UserService))  //用户服务

	go httpserver.Gw()

	listen, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(listen)

	// 	mux := http.NewServeMux()
	// 	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

	// 		fmt.Println(r)
	// 		rpcServer.ServeHTTP(rw, r)
	// 	})

	// 	httpServer := &http.Server{
	// 		Addr:    ":8081",
	// 		Handler: mux,
	// 	}
	// 	httpServer.ListenAndServeTLS("keys/server.crt", "keys/server.key")
}
