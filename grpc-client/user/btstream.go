package user

import (
	"context"
	"fmt"
	"grpccli/services"
	"io"
	"log"

	"google.golang.org/grpc"
)

func BtStream(conn *grpc.ClientConn) {
	var i int32
	userClient := services.NewUserServiceClient(conn) //用户客户端
	stream, err := userClient.GetUserScoreByTWS(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var uuid int32 = 1
	for j := 1; j <= 3; j++ {
		req := services.UserScoreRequest{}
		req.Users = make([]*services.UserInfo, 0)
		for i = 1; i < 6; i++ {
			req.Users = append(req.Users, &services.UserInfo{UserId: uuid})
			uuid++
		}
		err := stream.Send(&req)
		if err != nil {
			log.Fatalln(err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("第%d次接收:%v\n", j, res.Users)

	}
	// req, _ := stream.CloseAndRecv()
	// fmt.Println(req.Users)

}
