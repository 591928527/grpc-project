package services

import (
	context "context"
	"fmt"
)

type OderService struct {
}

func (this *OderService) NewOrder(ctx context.Context, orderRequest *OrderRequest) (*OrderResponse, error) {
	err := orderRequest.OrderMain.Validate()
	if err != nil {
		return &OrderResponse{
			Status:  "error",
			Message: err.Error(),
		}, nil
	}
	fmt.Println(orderRequest.OrderMain)
	fmt.Println(orderRequest.OrderMain.OrderDetails)
	resp := &OrderResponse{
		Status:  "ok",
		Message: "success",
	}
	return resp, nil
}
