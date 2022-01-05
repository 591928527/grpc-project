package goods

import (
	"context"
	"fmt"
	"grpccli/services"
	"log"
	"time"

	"google.golang.org/grpc"
)

type CustomCredential struct{}

func ProdClient(conn *grpc.ClientConn) {
	prodClient := services.NewProdServiceClient(conn) //商品客户端

	prodRes, err := prodClient.GetProdInfo(context.Background(), &services.ProdRequest{ProdId: 20, ProdArea: services.ProdAres_C})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("respons:", prodRes.ProdName)

	// prodRes, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 20, ProdArea: services.ProdAres_C})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("respons:", prodRes.ProdStock)

	// prodlistRes, err := prodClient.GetProdStocks(context.Background(), &services.QuerySize{Size: 10})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("respons:", prodlistRes.Prodres[2].ProdStock)
}

func (c CustomCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

func (c CustomCredential) RequireTransportSecurity() bool {
	return true
}

// interceptor 客户端拦截器
func Interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("method=%s req=%v rep=%v duration=%s error=%v\n", method, req, reply, time.Since(start), err)
	return err
}
