package customer

import (
	"context"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/customer/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	api.CustomerServiceServer
}

func (*server) AddCustomer(ctx context.Context, req *api.AddCustomerRequest) (*api.AddCustomerReply, error) {
	return &api.AddCustomerReply{Customer: req.GetCustomer()}, nil
}

func main() {
	flagHost := flag.String("host", "127.0.0.1", "address of customer service")
	flagPort := flag.String("port", "50051", "port of customer service")
	flagRedis := flag.String("redis", "127.0.0.1:6379", "address and port of Redis server")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *flagHost, *flagPort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	api.RegisterCustomerServiceServer(s, &server{})

	rdb := redis.NewClient(&redis.Options{
		Addr:     *flagRedis,
		Password: "",
	})

	go func() {
		for {
			rdb.Set(context.TODO(), "service:customer", address, 13*time.Second)
			time.Sleep(10 * time.Second)
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
