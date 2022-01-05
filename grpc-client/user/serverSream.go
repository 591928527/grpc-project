package user

import (
	"context"
	"fmt"
	"grpccli/services"
	"io"
	"log"

	"google.golang.org/grpc"
)

func ServerStream(conn *grpc.ClientConn) {
	var i int32
	userClient := services.NewUserServiceClient(conn) //用户客户端
	req := services.UserScoreRequest{}
	req.Users = make([]*services.UserInfo, 0)
	for i = 1; i < 6; i++ {
		req.Users = append(req.Users, &services.UserInfo{UserId: i})
	}

	// 使用服务端流的方式
	stream, err := userClient.GetUserScoreByServerSream(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp.Users)
	}

}
