package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
	"time"
)

func main() {
	flagRedis := flag.String("redis", "127.0.0.1:6379", "customerAddress and port of Redis server")
	flagNATS := flag.String("nats", "127.0.0.1:4222", "customerAddress and port of NATS server")
	flag.Parse()

	rdb := redis.NewClient(&redis.Options{
		Addr:     *flagRedis,
		Password: "",
	})

	// build connection to customer service
	customerAddress, err := rdb.Get(context.TODO(), "service:customerApi").Result()
	if err != nil {
		log.Fatalf("error while trying to get the result %v", err)
	}

	customerConn, err := grpc.Dial(customerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while closing the connection %v", err)
		}
	}(customerConn)

	//customerClient := services.NewCustomerServiceClient(customerConn)

	// set context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// subscribe to nats logging
	nc, err := nats.Connect(*flagNATS)
	if err != nil {
		log.Fatal("cannot connect to nats")
	}
	defer nc.Close()

	subscription, err := nc.Subscribe("log.*", func(msg *nats.Msg) {
		fmt.Printf("LOG: \tgot message from subject: %s\n\tdata: %s\n", msg.Subject, string(msg.Data))
	})
	if err != nil {
		log.Fatal("cannot subscribe")
	}
	defer func(subscription *nats.Subscription) {
		err := subscription.Unsubscribe()
		if err != nil {
			log.Fatal("cannot unsubscribe")
		}
	}(subscription) //nolint

	var wc sync.WaitGroup
	wc.Add(1)

	fmt.Println("Context: ", ctx)

	// wait for async
	wc.Wait()
}
