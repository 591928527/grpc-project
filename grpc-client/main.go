package main

import (
	"grpccli/goods"
	"grpccli/hander"
	"log"

	"google.golang.org/grpc"
)

var OpenSTL = true

func main() {
	var opts []grpc.DialOption
	if OpenSTL {
		opts = append(opts, grpc.WithTransportCredentials(hander.GetClientCreds()))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithPerRPCCredentials(new(goods.CustomCredential)))
	// 指定客户端interceptor
	opts = append(opts, grpc.WithUnaryInterceptor(goods.Interceptor))

	conn, err := grpc.Dial(":8081", opts...)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	//用户客户端使用客户端流的方式
	// user.ClientSream(conn)

	//用户客户端使用服务端流的方式
	// user.ServerStream(conn)

	//用户客户端使用客户端流的方式
	// user.BtStream(conn)

	//订单客户端
	// order.OrderClint(conn)

	//商品客户端
	goods.ProdClient(conn)

}
