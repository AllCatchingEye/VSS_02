package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/paymentApi"
	"google.golang.org/grpc"
	"log"
	"net"
	"reflect"
	"time"
)

type server struct {
	paymentApi.PaymentServiceServer
	redis       *redis.Client
	nats        *nats.Conn
	payedOrders []uint32
}

func (state *server) PayMyOrder(ctx context.Context, req *paymentApi.PayMyOrderRequest) (*paymentApi.PayMyOrderReply, error) {
	//TODO: check customerId (call CustomerService)
	//TODO: check orderId (call OrderService)
	fmt.Println("PayMyOrder called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req.GetOrderId())
	err := state.nats.Publish("log.paymentApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.paymentApi: cannot publish event")
	}
	state.payedOrders = append(state.payedOrders, req.OrderId)
	return &paymentApi.PayMyOrderReply{OrderId: req.OrderId}, nil
}

func (state *server) IsOrderPayed(ctx context.Context, req *paymentApi.IsOrderPayedRequest) (*paymentApi.IsOrderPayedReply, error) {
	fmt.Println("IsOrderPayed called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req.GetOrderId())
	err := state.nats.Publish("log.paymentApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.paymentApi: cannot publish event")
	}
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
	flagHost := flag.String("host", "127.0.0.1", "address of shipmentApi service")
	flagPort := flag.String("port", "50053", "port of shipmentApi service")
	flagRedis := flag.String("redis", "127.0.0.1:6379", "address and port of Redis server")
	flagNATS := flag.String("nats", "127.0.0.1:4222", "address and port of NATS server")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *flagHost, *flagPort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	rdb := redis.NewClient(&redis.Options{
		Addr:     *flagRedis,
		Password: "",
	})

	go func() {
		fmt.Println("starting to update redis for paymentApi service")
		for {
			rdb.Set(context.TODO(), "service:paymentApi", address, 13*time.Second)
			time.Sleep(10 * time.Second)
		}
	}()

	nc, err := nats.Connect(*flagNATS)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	paymentApi.RegisterPaymentServiceServer(s, &server{redis: rdb, nats: nc, payedOrders: []uint32{}})
	fmt.Println("creating paymentApi service finished")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
