package main

import (
	"context"
	"flag"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/customer/api"
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

	address, err := rdb.Get(context.TODO(), "service:customer").Result()
	if err != nil {
		log.Fatalf("error while trying to get the result %v", err)
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

	c := customerApi.NewCustomerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// Adding Customer
	customer := &customerApi.Customer{
		Name: "Max Mustermann",
		Address: &customerApi.Address{
			Street:  "Mustermannstraße 42",
			Zip:     "80335",
			City:    "Munich",
			Country: "Germany",
		},
	}
	r, err := c.AddCustomer(ctx, &customerApi.AddCustomerRequest{Customer: customer})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	log.Printf("Getting id: %d", r.GetCustomerId())

	// Getting Customer
	r2, err := c.GetCustomer(ctx, &customerApi.GetCustomerRequest{CustomerId: r.GetCustomerId()})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r2.GetCustomer().GetName())

	// Remove Customer
	r3, err := c.RemoveCustomer(ctx, &customerApi.RemoveCustomerRequest{CustomerId: r.GetCustomerId()})
	if err != nil {
		log.Fatalf("could not remove: %v", err)
	}
	log.Printf("Removed: %s", r3.GetCustomer().GetName())

}
