package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/customerApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/orderApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/paymentApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/services"
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
	fmt.Println("CustomerID is: ", req.GetCustomerId())
	fmt.Println("Products are: ", req.GetProducts())
	err := state.nats.Publish("log.orderApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.orderApi: cannot publish event")
	}
	// Check if customer exists
	if req.GetCustomerId() != 0 {
		customer, err := checkCustomerID(state.redis, req.GetCustomerId())
		if err != nil {
			return nil, err
		}
		fmt.Println("Could get customer: ", customer.GetName())
	}

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
	// Publish message for stock service
	err = state.nats.Publish("sto.order", []byte(fmt.Sprintf("orderId %v", orderId)))
	if err != nil {
		return nil, err
	}
	state.orders[orderId] = order
	return &orderApi.NewOrderReply{OrderId: orderId, Order: order}, nil
}

func (state *server) GetOrder(ctx context.Context, req *orderApi.GetOrderRequest) (*orderApi.GetOrderReply, error) {
	fmt.Println("GetOrder called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println("OrderID is: ", req.GetOrderId())
	fmt.Println("CustomerID is: ", req.GetCustomerId())
	err := state.nats.Publish("log.orderApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.orderApi: cannot publish event")
	}
	// Check if customer exists
	if req.GetCustomerId() != 0 {
		fmt.Println("checking customerID: ", req.GetCustomerId())
		customer, err := checkCustomerID(state.redis, req.GetCustomerId())
		if err != nil {
			return nil, err
		}
		fmt.Println("Could get customer: ", customer.GetName())
	}

	orderID := req.GetOrderId()
	order, ok := state.orders[orderID]
	if !ok {
		return nil, fmt.Errorf("orderApi not found")
	}
	fmt.Println("Order with ID ", orderID, " has products: ", order.GetProducts())
	return &orderApi.GetOrderReply{OrderId: orderID, Order: order}, nil
}

func (state *server) SetOrderStatus(ctx context.Context, req *orderApi.SetOrderStatusRequest) (*orderApi.SetOrderStatusReply, error) {
	fmt.Println("SetOrderStatus called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println("OrderID is: ", req.GetOrderId())
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
	fmt.Println("OrderID is: ", req.GetOrderId())
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
	fmt.Println("OrderID is: ", req.GetOrderId())
	err := state.nats.Publish("log.orderApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.orderApi: cannot publish event")
	}
	orderID := req.GetOrderId()
	order, ok := state.orders[orderID]
	if !ok {
		return nil, fmt.Errorf("order not found: %v", orderID)
	}
	order.DeliveryStatus = req.GetStatus()
	return &orderApi.SetDeliveryStatusReply{DeliveryStatus: order.GetDeliveryStatus()}, nil
}

func (state *server) CancelOrder(ctx context.Context, req *orderApi.CancelOrderRequest) (*orderApi.CancelOrderReply, error) {
	fmt.Println("CancelOrder called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	err := state.nats.Publish("log.orderApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.orderApi: cannot publish event")
	}
	// Check if customer exists
	if req.GetCustomerId() != 0 {
		customer, err := checkCustomerID(state.redis, req.GetCustomerId())
		if err != nil {
			return nil, err
		}
		fmt.Println("Could get customer: ", customer.GetName())
	}

	orderID := req.GetOrderId()
	order, ok := state.orders[orderID]
	if !ok {
		return nil, fmt.Errorf("order not found: %v", orderID)
	}
	if order.GetDeliveryStatus() == types.DELIVERY_STATUS_NOT_SENT {
		// call stock to restock the products
		err := state.nats.Publish("sto.cancel", []byte(fmt.Sprintf("cancel %v", orderID)))
		if err != nil {
			return nil, fmt.Errorf("cannot publish event")
		}
		// call payment to refund the money
		if order.GetPaymentStatus() {
			refundOrder(state.redis, orderID, req.GetCustomerId())
		}
	}
	return &orderApi.CancelOrderReply{OrderCanceled: true}, nil
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

func refundOrder(rdb *redis.Client, orderID uint32, customerID uint32) bool {
	paymentAddress, err := rdb.Get(context.TODO(), "service:paymentApi").Result()
	if err != nil {
		log.Fatalf("error while trying to get the payment service address %v", err)
	}

	paymentConn, err := grpc.Dial(paymentAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to payment service: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while closing the connection to payment service %v", err)
		}
	}(paymentConn)

	paymentClient := services.NewPaymentServiceClient(paymentConn)

	res, err := paymentClient.RefundMyOrder(context.Background(), &paymentApi.RefundMyOrderRequest{OrderId: orderID, CustomerId: customerID})
	if err != nil {
		log.Fatalf("error while trying to refund the order %v", err)
	}
	fmt.Println("refund successful: ", res.GetRefundSuccess())
	return res.GetRefundSuccess()
}
