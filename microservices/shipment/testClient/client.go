package main

import (
	"context"
	"flag"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/services"
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

	address, err := rdb.Get(context.TODO(), "service:shipmentApi").Result()
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

	c := services.NewShipmentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Ship Order
	//TODO: Hier müssen noch die Funktionen getestet werden
	log.Printf("shipment client: %v\n", c)
	log.Printf("context: %v\n", ctx)
	log.Println("Done")
}
