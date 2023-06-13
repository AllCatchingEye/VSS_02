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
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/stockApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"net"
	"reflect"
	"time"
)

type server struct {
	services.OrderServiceServer
	redis  *redis.Client
	nats   *nats.Conn
	orders map[uint32]*types.Order
}

func (state *server) NewOrder(ctx context.Context, req *orderApi.NewOrderRequest) (*orderApi.NewOrderReply, error) {
	fmt.Println("NewOrder called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req)
	err := state.nats.Publish("log.orderApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.orderApi: cannot publish event")
	}
	// Check if customer exists
	customer, err := checkCustomerID(state.redis, req.GetCustomerId())
	if err != nil {
		return nil, err
	}
	fmt.Println("Could get customer: ", customer.GetName())

	// Create new order
	productsStatus := make(map[uint32]bool)
	for _, product := range req.GetProducts() {
		productsStatus[product] = false
	}
	order := &types.Order{
		Customer:       req.GetCustomerId(),
		Products:       req.GetProducts(),
		ProductsStatus: productsStatus,
		OrderStatus:    false,
		PaymentStatus:  false,
		DeliveryStatus: types.DELIVERY_STATUS(0),
	}
	orderId := generateUniqueOrderID(state.orders)
	state.orders[orderId] = order
	// call stock to reserve products
	res, err := callStock(state.redis, orderId, order.GetProducts(), productsStatus)
	if err != nil {
		return nil, err
	}
	fmt.Println("Stock received request: ", res)
	state.orders[orderId] = order
	return &orderApi.NewOrderReply{OrderId: orderId, Order: order}, nil
}

func (state *server) GetOrder(ctx context.Context, req *orderApi.GetOrderRequest) (*orderApi.GetOrderReply, error) {
	fmt.Println("GetOrder called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req)
	err := state.nats.Publish("log.orderApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.orderApi: cannot publish event")
	}
	// Check if customer exists
	fmt.Println("checking customerID: ", req.GetCustomerId())
	customer, err := checkCustomerID(state.redis, req.GetCustomerId())
	if err != nil {
		return nil, err
	}
	fmt.Println("Could get customer: ", customer.GetName())

	orderID := req.GetOrderId()
	order, ok := state.orders[orderID]
	if !ok {
		return nil, fmt.Errorf("orderApi not found")
	}
	return &orderApi.GetOrderReply{OrderId: orderID, Order: order}, nil
}

func (state *server) SetOrderStatus(ctx context.Context, req *orderApi.SetOrderStatusRequest) (*orderApi.SetOrderStatusReply, error) {
	fmt.Println("SetOrderStatus called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req)
	err := state.nats.Publish("log.orderApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.orderApi: cannot publish event")
	}
	orderID := req.GetOrderId()
	order, ok := state.orders[orderID]
	if !ok {
		return nil, fmt.Errorf("orderApi not found")
	}
	order.OrderStatus = req.GetStatus()
	return &orderApi.SetOrderStatusReply{OrderStatus: order.GetOrderStatus()}, nil
}

func (state *server) SetPaymentStatus(ctx context.Context, req *orderApi.SetPaymentStatusRequest) (*orderApi.SetPaymentStatusReply, error) {
	fmt.Println("SetPaymentStatus called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req)
	err := state.nats.Publish("log.orderApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.orderApi: cannot publish event")
	}
	orderID := req.GetOrderId()
	order, ok := state.orders[orderID]
	if !ok {
		return nil, fmt.Errorf("order not found")
	}
	order.PaymentStatus = req.GetStatus()
	return &orderApi.SetPaymentStatusReply{PaymentStatus: order.GetPaymentStatus()}, nil
}

func (state *server) SetDeliveryStatus(ctx context.Context, req *orderApi.SetDeliveryStatusRequest) (*orderApi.SetDeliveryStatusReply, error) {
	fmt.Println("SetDeliveryStatus called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req)
	err := state.nats.Publish("log.orderApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.orderApi: cannot publish event")
	}
	orderID := req.GetOrderId()
	order, ok := state.orders[orderID]
	if !ok {
		return nil, fmt.Errorf("orderApi not found")
	}
	order.DeliveryStatus = req.GetStatus()
	return &orderApi.SetDeliveryStatusReply{DeliveryStatus: order.GetDeliveryStatus()}, nil
}

func main() {
	flagHost := flag.String("host", "127.0.0.1", "address of orderApi service")
	flagPort := flag.String("port", "50052", "port of orderApi service")
	flagRedis := flag.String("redis", "127.0.0.1:6379", "address and port of Redis server")
	flagNATS := flag.String("nats", "127.0.0.1:4222", "address and port of NATS server")
	flag.Parse()

	time.Sleep(5 * time.Second)

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
		fmt.Println("starting to update redis for orderApi service")
		for {
			rdb.Set(context.TODO(), "service:orderApi", address, 13*time.Second)
			time.Sleep(10 * time.Second)
		}
	}()

	nc, err := nats.Connect(*flagNATS)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	services.RegisterOrderServiceServer(s, &server{redis: rdb, nats: nc, orders: make(map[uint32]*types.Order)})
	fmt.Println("creating orderApi service finished")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Helper
func generateUniqueOrderID(orders map[uint32]*types.Order) uint32 {
	orderId := uint32(rand.Intn(1000))
	if len(orders) == 0 {
		return orderId
	}
	_, exists := orders[orderId]
	for exists {
		orderId = uint32(rand.Intn(1000))
		_, exists = orders[orderId]
	}
	return orderId
}

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
	if err != nil {
		return nil, fmt.Errorf("customer with ID %v does not exist: %v", customerID, err)
	}
	return res.GetCustomer(), nil
}

func callStock(client *redis.Client, orderId uint32, products map[uint32]uint32, productsStatus map[uint32]bool) (bool, error) {
	// Check if stock exists
	stockAddress, err := client.Get(context.Background(), "service:stockApi").Result()
	if err != nil {
		log.Fatalf("error while trying to get the stock service address %v", err)
	}
	fmt.Println("stockAddress successful.")
	stockConn, err := grpc.Dial(stockAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to stock service: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while closing the connection to stock service %v", err)
		}
	}(stockConn)
	fmt.Println("stockConn successful.")

	stockClient := services.NewStockServiceClient(stockConn)
	fmt.Println("stockClient successful.")
	// send order to stock service
	res, err := stockClient.OrderProducts(context.Background(), &stockApi.OrderProductsRequest{OrderId: orderId, Products: products, ProductsStatus: productsStatus})
	if err != nil {
		return false, fmt.Errorf("stock service could not process order %v: %v", orderId, err)
	}
	return res.GetRequestSuccessful(), nil
}
