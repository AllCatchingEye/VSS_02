package main

import (
	"context"
	"flag"
	"github.com/redis/go-redis/v9"
	paymentApi "gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/payment/api"
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

	address, err := rdb.Get(context.TODO(), "service:payment").Result()
	if err != nil {
		log.Fatalf("error while trying to get the payment service address: %v", err)
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

	c := paymentApi.NewPaymentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// Adding Products
	r, err := c.PayMyOrder(ctx, &paymentApi.PayMyOrderRequest{CustomerId: 1, OrderId: 1})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	log.Printf("order got payed, id: %d", r.GetOrderId())

	// check payment status
	r2, err := c.IsOrderPayed(ctx, &paymentApi.IsOrderPayedRequest{OrderId: 1})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	log.Printf("order is payed: %t", r2.GetIsPayed())
}
