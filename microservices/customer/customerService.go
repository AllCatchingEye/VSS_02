package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/customer/api"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
)

type server struct {
	customerApi.CustomerServiceServer
	customers map[uint32]*customerApi.Customer
}

func (state *server) AddCustomer(ctx context.Context, req *customerApi.AddCustomerRequest) (*customerApi.AddCustomerReply, error) {
	fmt.Println("AddCustomer called")
	fmt.Println(req.GetCustomer())
	newCustomer := req.GetCustomer()
	customerID := generateUniqueCustomerID(state.customers)
	state.customers[customerID] = newCustomer
	return &customerApi.AddCustomerReply{CustomerId: customerID}, nil
}

func (state *server) GetCustomer(ctx context.Context, req *customerApi.GetCustomerRequest) (*customerApi.GetCustomerReply, error) {
	fmt.Println("GetCustomer called")
	fmt.Println(req.GetCustomerId())
	customerId := req.GetCustomerId()
	customer, ok := state.customers[customerId]
	if !ok {
		return nil, fmt.Errorf("customer not found")
	}
	return &customerApi.GetCustomerReply{Customer: customer}, nil
}

func (state *server) RemoveCustomer(ctx context.Context, req *customerApi.RemoveCustomerRequest) (*customerApi.RemoveCustomerReply, error) {
	fmt.Println("RemoveCustomer called")
	fmt.Println(req.GetCustomerId())
	customerId := req.GetCustomerId()
	customer, ok := state.customers[customerId]
	if !ok {
		return nil, fmt.Errorf("customer not found")
	}
	delete(state.customers, customerId)
	return &customerApi.RemoveCustomerReply{Customer: customer}, nil
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

	customerApi.RegisterCustomerServiceServer(s, &server{customers: make(map[uint32]*customerApi.Customer)})
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

// Helper
func generateUniqueCustomerID(customers map[uint32]*customerApi.Customer) uint32 {
	customerID := uint32(rand.Intn(1000))
	if len(customers) == 0 {
		return customerID
	}
	_, exists := customers[customerID]
	for exists {
		customerID = uint32(rand.Intn(1000))
		_, exists = customers[customerID]
	}
	return customerID
}
