package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	orderApi "gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/orderApi"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
)

type server struct {
	orderApi.OrderServiceServer
	orders map[uint32]*orderApi.Order
}

func (state *server) NewOrder(ctx context.Context, req *orderApi.NewOrderRequest) (*orderApi.NewOrderReply, error) {
	fmt.Println("NewOrder called")
	fmt.Println(req)
	//TODO: check if customerApi exists (call customerApi service)
	order := &orderApi.Order{
		Customer:       req.GetCustomerId(),
		Products:       req.GetProducts(),
		OrderStatus:    false,
		PaymentStatus:  false,
		DeliveryStatus: orderApi.DELIVERY_STATUS(0),
	}
	orderId := generateUniqueOrderID(state.orders)
	state.orders[orderId] = order
	return &orderApi.NewOrderReply{OrderId: orderId, Order: order}, nil
}

func (state *server) GetOrder(ctx context.Context, req *orderApi.GetOrderRequest) (*orderApi.GetOrderReply, error) {
	fmt.Println("GetOrder called")
	fmt.Println(req)
	orderID := req.GetOrderId()
	order, ok := state.orders[orderID]
	if !ok {
		return nil, fmt.Errorf("orderApi not found")
	}
	return &orderApi.GetOrderReply{OrderId: orderID, Order: order}, nil
}

func (state *server) SetOrderStatus(ctx context.Context, req *orderApi.SetOrderStatusRequest) (*orderApi.SetOrderStatusReply, error) {
	fmt.Println("SetOrderStatus called")
	fmt.Println(req)
	orderID := req.GetOrderId()
	order, ok := state.orders[orderID]
	if !ok {
		return nil, fmt.Errorf("orderApi not found")
	}
	order.OrderStatus = req.GetStatus()
	return &orderApi.SetOrderStatusReply{OrderStatus: order.GetOrderStatus()}, nil
}

func (state *server) SetPaymentStatus(ctx context.Context, req *orderApi.SetPaymentStatusRequest) (*orderApi.SetPaymentStatusReply, error) {
	fmt.Println("SetPaymentStatus called")
	fmt.Println(req)
	orderID := req.GetOrderId()
	order, ok := state.orders[orderID]
	if !ok {
		return nil, fmt.Errorf("orderApi not found")
	}
	order.PaymentStatus = req.GetStatus()
	return &orderApi.SetPaymentStatusReply{PaymentStatus: order.GetPaymentStatus()}, nil
}

func (state *server) SetDeliveryStatus(ctx context.Context, req *orderApi.SetDeliveryStatusRequest) (*orderApi.SetDeliveryStatusReply, error) {
	fmt.Println("SetDeliveryStatus called")
	fmt.Println(req)
	orderID := req.GetOrderId()
	order, ok := state.orders[orderID]
	if !ok {
		return nil, fmt.Errorf("orderApi not found")
	}
	order.DeliveryStatus = req.GetStatus()
	return &orderApi.SetDeliveryStatusReply{DeliveryStatus: order.GetDeliveryStatus()}, nil
}

//TODO: implement interface for stockApi service

func main() {
	flagHost := flag.String("host", "127.0.0.1", "address of orderApi service")
	flagPort := flag.String("port", "50052", "port of orderApi service")
	flagRedis := flag.String("redis", "127.0.0.1:6379", "address and port of Redis server")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *flagHost, *flagPort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	orderApi.RegisterOrderServiceServer(s, &server{orders: make(map[uint32]*orderApi.Order)})
	fmt.Println("creating orderApi service finished")

	rdb := redis.NewClient(&redis.Options{
		Addr:     *flagRedis,
		Password: "",
	})

	go func() {
		fmt.Println("starting to update redis for orderApi service")
		for {
			rdb.Set(context.TODO(), "service:orderApi", address, 13*time.Second)
			time.Sleep(10 * time.Second)
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Helper
func generateUniqueOrderID(orders map[uint32]*orderApi.Order) uint32 {
	orderId := uint32(rand.Intn(1000))
	if len(orders) == 0 {
		return orderId
	}
	_, exists := orders[orderId]
	for exists {
		orderId = uint32(rand.Intn(1000))
		_, exists = orders[orderId]
	}
	return orderId
}
