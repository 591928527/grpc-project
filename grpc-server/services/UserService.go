package services

import (
	context "context"
	"io"
	"log"
	"time"
)

type UserService struct {
}

//普通方法
func (this *UserService) GetUserScore(ctx context.Context, userScoreRequest *UserScoreRequest) (*UserScoreResponse, error) {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for _, user := range userScoreRequest.Users {
		user.UserScore = score
		score++
		users = append(users, user)
	}
	return &UserScoreResponse{Users: users}, nil
}

// 服务端流
func (this *UserService) GetUserScoreByServerSream(in *UserScoreRequest, stream UserService_GetUserScoreByServerSreamServer) error {

	var score int32 = 101
	users := make([]*UserInfo, 0)
	for index, user := range in.Users {
		user.UserScore = score
		score++
		users = append(users, user)

		if (index+1)%2 == 0 && index > 0 {
			err := stream.Send(&UserScoreResponse{Users: users})
			if err != nil {
				return err
			}
			users = (users)[0:0]
		}
		time.Sleep(time.Second * 2)

	}
	if len(users) > 0 {
		err := stream.Send(&UserScoreResponse{Users: users})
		if err != nil {
			return err
		}
	}
	return nil
}

// 客户端流
func (this *UserService) GetUserScoreByClientSream(stream UserService_GetUserScoreByClientSreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&UserScoreResponse{Users: users})
		}
		if err != nil {
			return err
		}

		for _, user := range req.Users {
			user.UserScore = score
			score++
			users = append(users, user)
		}
	}
}

//服务端、客户端双向流
func (this *UserService) GetUserScoreByTWS(stream UserService_GetUserScoreByTWSServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		for _, user := range req.Users {
			user.UserScore = score
			score++
			users = append(users, user)
		}
		err = stream.Send(&UserScoreResponse{Users: users})
		if err != nil {
			log.Fatalln(err)
		}
		users = (users)[0:0]
	}
}
