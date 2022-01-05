package services

import (
	context "context"
)

type ProdService struct {
}

func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	if request.ProdArea == ProdAres_A {
		stock = 300
	} else if request.ProdArea == ProdAres_B {
		stock = 31
	} else {
		stock = 50
	}
	return &ProdResponse{ProdStock: stock}, nil
}

func (this *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	produc := []*ProdResponse{
		&ProdResponse{ProdStock: 29},
		&ProdResponse{ProdStock: 30},
		&ProdResponse{ProdStock: 31},
		&ProdResponse{ProdStock: 32},
	}
	return &ProdResponseList{
		Prodres: produc,
	}, nil
}

func (this *ProdService) GetProdInfo(ctx context.Context, request *ProdRequest) (*ProdModel, error) {

	prod := &ProdModel{
		ProdId:    1,
		ProdName:  "测试商品",
		ProdPrice: 13.0,
	}
	return prod, nil
}
