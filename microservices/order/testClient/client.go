package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	orderApi "gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/order/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	flagRedis := flag.String("redis", "127.0.0.1:6379", "address and port of Redis server")
	flag.Parse()

	rdb := redis.NewClient(&redis.Options{
		Addr:     *flagRedis,
		Password: "",
	})

	address, err := rdb.Get(context.TODO(), "service:order").Result()
	if err != nil {
		log.Fatalf("error while trying to get the order service address: %v", err)
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while closing the connection %v", err)
		}
	}(conn)

	c := orderApi.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// New Order
	order := make(map[uint32]uint32)
	order[1] = 10
	order[2] = 5
	order[3] = 1

	r, err := c.NewOrder(ctx, &orderApi.NewOrderRequest{CustomerId: 1, Products: order})
	if err != nil {
		log.Fatalf("could not create order: %v", err)
	}
	log.Printf("order got payed, id: %#v", r.GetOrderId())

	// Set Order Status
	r2, err := c.SetOrderStatus(ctx, &orderApi.SetOrderStatusRequest{OrderId: r.GetOrderId(), Status: true})
	if err != nil {
		log.Fatalf("could not set order status: %v", err)
	}
	log.Printf("order status set to: %#v", r2.GetOrderStatus())

	// Set Payment Status
	r3, err := c.SetPaymentStatus(ctx, &orderApi.SetPaymentStatusRequest{OrderId: r.GetOrderId(), Status: true})
	if err != nil {
		log.Fatalf("could not set payment status: %v", err)
	}
	log.Printf("payment status set to: %#v", r3.GetPaymentStatus())

	// Set Delivery Status
	r4, err := c.SetDeliveryStatus(ctx, &orderApi.SetDeliveryStatusRequest{OrderId: r.GetOrderId(), Status: orderApi.DELIVERY_STATUS(1)})
	if err != nil {
		log.Fatalf("could not set delivery status: %v", err)
	}
	log.Printf("delivery status set to: %#v", r4.GetDeliveryStatus())

	// Get Order
	r5, err := c.GetOrder(ctx, &orderApi.GetOrderRequest{CustomerId: 1, OrderId: r.GetOrderId()})
	if err != nil {
		log.Fatalf("could not get order: %v", err)
	}
	log.Printf("order: %#v", r5.GetOrder())

	fmt.Println("Done")
}
