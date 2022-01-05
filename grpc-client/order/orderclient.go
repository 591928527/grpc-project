package order

import (
	"context"
	"fmt"
	"grpccli/services"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

func OrderClint(conn *grpc.ClientConn) {

	orderClient := services.NewOrderServiceClient(conn) //订单客户端
	t := timestamp.Timestamp{Seconds: time.Now().Unix()}

	resp, err := orderClient.NewOrder(context.Background(), &services.OrderRequest{
		OrderMain: &services.OrderMain{
			OrderId:    123,
			OrderNo:    "123456",
			OrderMoney: 12354,
			OrderTime:  &t,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
