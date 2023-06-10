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

func (state *server) ShipMyOrder(ctx context.Context, req *shipmentApi.ShipMyOrderRequest) (*shipmentApi.ShipMyOrderReply, error) {
	fmt.Println("ShipmentOrder called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req.GetOrderId())
	err := state.nats.Publish("log.shipmentApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("shipmentApi: cannot publish event")
	}
	// Check if customer exists
	fmt.Println("checking customerID: ", req.GetCustomerId())
	customer, err := checkCustomerID(state.redis, req.GetCustomerId())
	if err != nil {
		return nil, err
	}
	fmt.Println("Could get customer: ", customer.GetName())
	// Check if order exists
	fmt.Println("checking orderID: ", req.GetOrderId())
	orderID, err := checkOrderID(state.redis, req.GetOrderId(), req.GetCustomerId())
	if err != nil {
		return nil, err
	}
	fmt.Println("Could get order: ", orderID)
	state.sentOrders = append(state.sentOrders, orderID)

	// set delivery status
	_, err = setDeliveryStatus(state.redis, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	fmt.Println("Delivery status set successfully.")

	return &shipmentApi.ShipMyOrderReply{OrderId: orderID, Address: customer.GetAddress()}, nil
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

// Helper
func checkCustomerID(redis *redis.Client, customerID uint32) (*types.Customer, error) {
	// Check if customer exists
	customerAddress, err := redis.Get(context.TODO(), "service:customerApi").Result()
	if err != nil {
		log.Fatalf("error while trying to get the customer service address %v", err)
	}
	fmt.Println("customerAddress successful.")
	customerConn, err := grpc.Dial(customerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to customer service: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while closing the connection to customer service %v", err)
		}
	}(customerConn)
	fmt.Println("customerConn successful.")

	customerClient := services.NewCustomerServiceClient(customerConn)
	fmt.Println("customerClient successful.")
	res, err := customerClient.GetCustomer(context.Background(), &customerApi.GetCustomerRequest{CustomerId: customerID})
	fmt.Println("checking customerID finished.")
	if err != nil {
		return nil, fmt.Errorf("customer with ID %v does not exist: %v", customerID, err)
	}
	return res.GetCustomer(), nil
}

func checkOrderID(redis *redis.Client, orderID uint32, customerID uint32) (uint32, error) {
	// Check if order exists
	orderAddress, err := redis.Get(context.TODO(), "service:orderApi").Result()
	if err != nil {
		log.Fatalf("error while trying to get the order service address %v", err)
	}
	fmt.Println("orderAddress successful.")
	orderConn, err := grpc.Dial(orderAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to order service: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while closing the connection to order service %v", err)
		}
	}(orderConn)
	fmt.Println("orderConn successful.")

	orderClient := services.NewOrderServiceClient(orderConn)
	fmt.Println("orderClient successful.")
	res, err := orderClient.GetOrder(context.Background(), &orderApi.GetOrderRequest{CustomerId: customerID, OrderId: orderID})
	if err != nil {
		return 0, fmt.Errorf("order with ID %v does not exist: %v", orderID, err)
	}
	return res.GetOrderId(), nil
}

func setDeliveryStatus(redis *redis.Client, orderID uint32) (types.DELIVERY_STATUS, error) {
	// Check if order exists
	orderAddress, err := redis.Get(context.TODO(), "service:orderApi").Result()
	if err != nil {
		log.Fatalf("error while trying to get the order service address %v", err)
	}
	fmt.Println("orderAddress successful.")
	orderConn, err := grpc.Dial(orderAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to order service: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while closing the connection to order service %v", err)
		}
	}(orderConn)
	fmt.Println("orderConn successful.")

	orderClient := services.NewOrderServiceClient(orderConn)
	fmt.Println("orderClient successful.")
	res, err := orderClient.SetDeliveryStatus(context.Background(), &orderApi.SetDeliveryStatusRequest{OrderId: orderID, Status: types.DELIVERY_STATUS_UNDER_WAY})
	if err != nil {
		return 0, fmt.Errorf("could not set payment status of order with ID %v: %v", orderID, err)
	}
	return res.GetDeliveryStatus(), nil
}
