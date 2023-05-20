package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/customer/api"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/order/api"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/shipment/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"time"
)

// TODO: Ist es erlaubt hier die anderen Typen zu importieren?
type server struct {
	shipmentApi.ShipmentServiceServer
	sentOrders []uint32
}

func (state *server) ShipmentOrder(ctx context.Context, req *shipmentApi.ShipMyOrderRequest) (*shipmentApi.ShipMyOrderReply, error) {
	address := getCustomerAddress(ctx, req)

	orderId := getCustomersOrder(ctx, req)

	state.sentOrders = append(state.sentOrders, orderId)

	return &shipmentApi.ShipMyOrderReply{
		OrderId: orderId,
		Address: ConvertToShipmentAddress(address),
	}, nil
}

func getCustomerAddress(ctx context.Context, req *shipmentApi.ShipMyOrderRequest) *customerApi.Address {
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

	res, err := c.GetCustomer(ctx, &customerApi.GetCustomerRequest{CustomerId: req.CustomerId})
	if err != nil {
		log.Fatalf("could not find customerId %d: %v", req.CustomerId, err)
	}

	return res.Customer.Address
}

func getCustomersOrder(ctx context.Context, req *shipmentApi.ShipMyOrderRequest) uint32 {
	flagRedis := flag.String("redis", "127.0.0.1:6379", "address and port of Redis server")
	flag.Parse()

	rdb := redis.NewClient(&redis.Options{
		Addr:     *flagRedis,
		Password: "",
	})

	address, err := rdb.Get(context.TODO(), "service:order").Result()
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

	c := orderApi.NewOrderServiceClient(conn)

	res, err := c.GetOrder(ctx, &orderApi.GetOrderRequest{CustomerId: req.CustomerId, OrderId: req.OrderId})
	if err != nil {
		log.Fatalf("could not find customerId %d: %v", req.CustomerId, err)
	}

	return res.OrderId
}

func main() {
	flagHost := flag.String("host", "127.0.0.1", "address of shipment service")
	flagPort := flag.String("port", "50054", "port of shipment service")
	flagRedis := flag.String("redis", "127.0.0.1:6379", "address and port of Redis server")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *flagHost, *flagPort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	shipmentApi.RegisterShipmentServiceServer(s, &server{})

	rdb := redis.NewClient(&redis.Options{
		Addr:     *flagRedis,
		Password: "",
	})

	go func() {
		for {
			rdb.Set(context.TODO(), "service:shipment", address, 13*time.Second)
			time.Sleep(10 * time.Second)
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// ConvertToShipmentAddress Helper
func ConvertToShipmentAddress(address *customerApi.Address) *shipmentApi.Address {
	return &shipmentApi.Address{
		Street:  address.GetStreet(),
		Zip:     address.GetZip(),
		City:    address.GetCity(),
		Country: address.GetCountry(),
	}
}
