package order

import (
	"context"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	customerApi "gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/customer/api"
	orderApi "gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/order/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	orderApi.OrderServiceServer
	customers map[uint32]*customerApi.Customer
}

func main() {
	flagHost := flag.String("host", "127.0.0.1", "address of order service")
	flagPort := flag.String("port", "50052", "port of order service")
	flagRedis := flag.String("redis", "127.0.0.1:6379", "address and port of Redis server")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *flagHost, *flagPort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	//orderApi.RegisterCustomerServiceServer(s, &server{customers: make(map[uint32]*customerApi.Customer)})
	fmt.Println("creating customer service finished")

	rdb := redis.NewClient(&redis.Options{
		Addr:     *flagRedis,
		Password: "",
	})

	go func() {
		fmt.Println("starting to update redis")
		for {
			rdb.Set(context.TODO(), "service:customer", address, 13*time.Second)
			time.Sleep(10 * time.Second)
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
