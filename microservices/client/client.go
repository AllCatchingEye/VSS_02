package client

import (
	"context"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	customerApi "gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/customer/api"
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

	fmt.Println("this is c: ", c)
	fmt.Println("this is ctx: ", ctx)
}
