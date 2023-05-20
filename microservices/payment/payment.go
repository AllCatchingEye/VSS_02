package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	paymentApi "gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/payment/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	paymentApi.PaymentServiceServer
	payedOrders []uint32
}

func (state *server) PayMyOrder(ctx context.Context, req *paymentApi.PayMyOrderRequest) (*paymentApi.PayMyOrderReply, error) {
	//TODO: check customerId (call CustomerService)
	//TODO: check orderId (call OrderService)
	state.payedOrders = append(state.payedOrders, req.OrderId)
	return &paymentApi.PayMyOrderReply{OrderId: req.OrderId}, nil
}

func (state *server) IsOrderPayed(ctx context.Context, req *paymentApi.IsOrderPayedRequest) (*paymentApi.IsOrderPayedReply, error) {
	fmt.Println("IsOrderPayed called")
	fmt.Println(req.GetOrderId())
	orderId := req.GetOrderId()
	//TODO: check orderId (call OrderService)
	for _, payedOrder := range state.payedOrders {
		if orderId == payedOrder {
			return &paymentApi.IsOrderPayedReply{IsPayed: true}, nil
		}
	}
	return &paymentApi.IsOrderPayedReply{IsPayed: false}, nil
}

func main() {
	flagHost := flag.String("host", "127.0.0.1", "address of shipment service")
	flagPort := flag.String("port", "50053", "port of shipment service")
	flagRedis := flag.String("redis", "127.0.0.1:6379", "address and port of Redis server")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *flagHost, *flagPort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	paymentApi.RegisterPaymentServiceServer(s, &server{})
	fmt.Println("creating payment service finished")

	rdb := redis.NewClient(&redis.Options{
		Addr:     *flagRedis,
		Password: "",
	})

	go func() {
		fmt.Println("starting to update redis for payment service")
		for {
			rdb.Set(context.TODO(), "service:payment", address, 13*time.Second)
			time.Sleep(10 * time.Second)
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
