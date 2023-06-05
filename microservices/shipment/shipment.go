package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/customerApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/orderApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/services"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/shipmentApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"reflect"
	"time"
)

// TODO: Ist es erlaubt hier die anderen Typen zu importieren?
type server struct {
	services.ShipmentServiceServer
	redis      *redis.Client
	nats       *nats.Conn
	sentOrders []uint32
}

func (state *server) ShipmentOrder(ctx context.Context, req *shipmentApi.ShipMyOrderRequest) (*shipmentApi.ShipMyOrderReply, error) {
	fmt.Println("ShipmentOrder called")
	fmt.Println(req.GetOrderId())
	err := state.nats.Publish("log.shipmentApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("shipmentApi: cannot publish event")
	}
	address := state.getCustomerAddress(ctx, req)

	orderId := state.getCustomersOrder(ctx, req)

	state.sentOrders = append(state.sentOrders, orderId)

	return &shipmentApi.ShipMyOrderReply{OrderId: orderId, Address: ConvertToShipmentAddress(address)}, nil
}

func (state *server) getCustomerAddress(ctx context.Context, req *shipmentApi.ShipMyOrderRequest) *types.Address {
	address, err := state.redis.Get(context.TODO(), "service:customerApi").Result()
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

	c := services.NewCustomerServiceClient(conn)

	res, err := c.GetCustomer(ctx, &customerApi.GetCustomerRequest{CustomerId: req.CustomerId})
	if err != nil {
		log.Fatalf("could not find customerId %d: %v", req.CustomerId, err)
	}

	return res.Customer.Address
}

func (state *server) getCustomersOrder(ctx context.Context, req *shipmentApi.ShipMyOrderRequest) uint32 {
	address, err := state.redis.Get(context.TODO(), "service:orderApi").Result()
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

	c := services.NewOrderServiceClient(conn)

	res, err := c.GetOrder(ctx, &orderApi.GetOrderRequest{CustomerId: req.CustomerId, OrderId: req.OrderId})
	if err != nil {
		log.Fatalf("could not find customerId %d: %v", req.CustomerId, err)
	}

	return res.OrderId
}

func main() {
	flagHost := flag.String("host", "127.0.0.1", "address of shipmentApi service")
	flagPort := flag.String("port", "50054", "port of shipmentApi service")
	flagRedis := flag.String("redis", "127.0.0.1:6379", "address and port of Redis server")
	flagNATS := flag.String("nats", "127.0.0.1:4222", "address and port of NATS server")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *flagHost, *flagPort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	rdb := redis.NewClient(&redis.Options{
		Addr:     *flagRedis,
		Password: "",
	})

	go func() {
		fmt.Println("starting to update redis")
		for {
			rdb.Set(context.TODO(), "service:shipmentApi", address, 13*time.Second)
			time.Sleep(10 * time.Second)
		}
	}()

	nc, err := nats.Connect(*flagNATS)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	services.RegisterShipmentServiceServer(s, &server{redis: rdb, nats: nc, sentOrders: []uint32{}})
	fmt.Println("creating shipmentApi service finished")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// ConvertToShipmentAddress Helper
func ConvertToShipmentAddress(address *types.Address) *types.Address {
	return &types.Address{
		Street:  address.GetStreet(),
		Zip:     address.GetZip(),
		City:    address.GetCity(),
		Country: address.GetCountry(),
	}
}
