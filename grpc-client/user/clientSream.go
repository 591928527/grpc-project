package user

import (
	"context"
	"fmt"
	"grpccli/services"
	"log"

	"google.golang.org/grpc"
)

func ClientSream(conn *grpc.ClientConn) {
	var i int32
	userClient := services.NewUserServiceClient(conn) //用户客户端
	stream, err := userClient.GetUserScoreByClientSream(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for j := 0; j < 3; j++ {
		req := services.UserScoreRequest{}
		req.Users = make([]*services.UserInfo, 0)
		for i = 1; i < 6; i++ {
			req.Users = append(req.Users, &services.UserInfo{UserId: i})
		}
		err := stream.Send(&req)
		if err != nil {
			log.Fatalln(err)
		}

	}
	req, _ := stream.CloseAndRecv()
	fmt.Println(req.Users)
}
