package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/orderApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/services"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/stockApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/supplierApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"net"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

type server struct {
	services.StockServiceServer
	redis    *redis.Client
	nats     *nats.Conn
	products map[uint32]*types.Product
	orders   map[uint32]map[uint32]uint32
}

func (state *server) AddProducts(ctx context.Context, req *stockApi.AddProductsRequest) (*stockApi.AddProductsReply, error) {
	fmt.Println("AddProduct called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req.GetProducts())
	err := state.nats.Publish("log.stockApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.stockApi: cannot publish event")
	}
	fmt.Println("Nats logs finished.")
	newProducts := req.GetProducts()
	var productIDs []uint32
	for _, product := range newProducts {
		productID := product.GetProductId()
		if productID == 0 {
			productID = generateUniqueProductID(state.products)
		}
		productIDs = append(productIDs, productID)
		state.products[productID] = product
	}
	fmt.Println("Stock: adding products succeeded.")
	// Add products to supplier
	supplierAddress, err := state.redis.Get(context.TODO(), "service:supplierApi").Result()
	if err != nil {
		log.Fatalf("error while trying to get the result %v", err)
	}
	fmt.Println("supplierAddress successful.")
	supplierConn, err := grpc.Dial(supplierAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while closing the connection %v", err)
		}
	}(supplierConn)
	fmt.Println("supplierConn successful.")

	supplierClient := services.NewSupplierServiceClient(supplierConn)
	fmt.Println("supplierClient successful.")
	res, err := supplierClient.AddProducts(context.Background(), &supplierApi.AddProductsRequest{SupplierId: newProducts[0].Supplier, Products: newProducts})
	if err != nil {
		return nil, err
	}
	fmt.Println("Added products to supplier: ", res.GetSupplier().GetName(), "(ID: ", res.GetSupplier().GetSupplierId(), ")")
	return &stockApi.AddProductsReply{ProductIds: productIDs}, nil
}

func (state *server) GetProducts(ctx context.Context, req *stockApi.GetProductsRequest) (*stockApi.GetProductsReply, error) {
	fmt.Println("GetProduct called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req.GetProductIds())
	err := state.nats.Publish("log.stockApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.stockApi: cannot publish event")
	}
	productIDs := req.GetProductIds()
	var products []*types.Product
	for _, productID := range productIDs {
		product, ok := state.products[productID]
		if !ok {
			return nil, fmt.Errorf("product not found")
		}
		products = append(products, product)
	}
	return &stockApi.GetProductsReply{Products: products}, nil
}

func (state *server) RemoveProduct(ctx context.Context, req *stockApi.RemoveProductRequest) (*stockApi.RemoveProductReply, error) {
	fmt.Println("RemoveProduct called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req.GetProductId())
	err := state.nats.Publish("log.stockApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.stockApi: cannot publish event")
	}
	productID := req.GetProductId()
	product, ok := state.products[productID]
	if !ok {
		return nil, fmt.Errorf("product not found")
	}
	delete(state.products, productID)
	return &stockApi.RemoveProductReply{Product: product}, nil
}

func (state *server) OrderProducts(ctx context.Context, req *stockApi.OrderProductsRequest) (*stockApi.OrderProductsReply, error) {
	fmt.Println("OrderProduct called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	err := state.nats.Publish("log.stockApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.stockApi: cannot publish event")
	}
	orderId := req.GetOrderId()
	fmt.Println("Received order: ", orderId)
	orderProducts := req.GetProducts()
	fmt.Println("Order products: ", orderProducts)
	productsCount := len(orderProducts)
	state.orders[orderId] = orderProducts

	for product, amount := range orderProducts {
		fmt.Println("Handle product ", product, " with amount ", amount, " in order ", orderId, ".")
		if state.products[product] != nil && state.products[product].GetAmount() >= amount {
			// reserve products in stock (dekrement number)
			fmt.Println("Reserve product ", product)
			state.orders[orderId][product] = 0
			state.products[product].Amount -= amount
			productsCount--
		} else if state.products[product] != nil && state.products[product].GetAmount() < amount {
			// Order from supplier
			fmt.Println("Order product ", product, " from supplier ", state.products[product].GetSupplier())
			err = state.nats.Publish("supp.orderProduct", []byte(fmt.Sprintf("order %v %v %v", state.products[product].GetSupplier(), product, amount-state.products[product].GetAmount()+2*state.products[product].GetAmount())))
			if err != nil {
				log.Print("supp.orderProduct: cannot publish event")
			}
		} else {
			fmt.Println("Product ", product, " not found")
			delete(orderProducts, product)
			delete(state.orders[orderId], product)
		}
	}
	if productsCount == 0 {
		// set order status
		fmt.Println("Order ", orderId, " is complete")
		setOrderStatus(state.redis, orderId)
	}

	if len(orderProducts) == 0 {
		fmt.Println("no products available")
	}
	return &stockApi.OrderProductsReply{Received: true}, nil
}

func main() {
	flagHost := flag.String("host", "127.0.0.1", "address of customerApi service")
	flagPort := flag.String("port", "50055", "port of customerApi service")
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
		fmt.Println("starting to update redis")
		for {
			rdb.Set(context.TODO(), "service:stockApi", address, 13*time.Second)
			time.Sleep(10 * time.Second)
		}
	}()

	nc, err := nats.Connect(*flagNATS)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	server := &server{redis: rdb, nats: nc, products: make(map[uint32]*types.Product), orders: make(map[uint32]map[uint32]uint32)}

	subscription, err := nc.Subscribe("supp.deliverProduct", func(msg *nats.Msg) {
		fmt.Printf("LOG: \tgot message from subject: %s\n\tdata: %s\n", msg.Subject, string(msg.Data))
		message := strings.Split(string(msg.Data), " ")
		supplier := message[0]
		product := message[1]
		amount := message[2]
		fmt.Printf("supplier: %s, product: %s, amount: %s\n", supplier, product, amount)
		productID, _ := strconv.ParseUint(product, 10, 32)
		amountUint, _ := strconv.ParseUint(amount, 10, 32)
		server.products[uint32(productID)].Amount += uint32(amountUint)
		// für alle offenen Bestellungen prüfen, ob jetzt genug Produkte vorhanden sind
		for orderID, orderProducts := range server.orders {
			fmt.Println("Check order ", orderID)
			productsCount := len(orderProducts)
			for product, amount := range orderProducts {
				fmt.Println("Handle product ", product, " with amount ", amount, " in order ", orderID, ".")
				if server.products[product] != nil && server.products[product].GetAmount() >= amount {
					// reserve products in stock (dekrement number)
					fmt.Println("Reserve product ", product)
					server.orders[orderID][product] = 0
					server.products[product].Amount -= amount
					productsCount--
				}
			}
			if productsCount == 0 {
				// set order status
				fmt.Println("Order ", orderID, " is complete")
				setOrderStatus(server.redis, orderID)
			}
		}
	})
	if err != nil {
		log.Fatal("cannot subscribe")
	}
	defer func(subscription *nats.Subscription) {
		err := subscription.Unsubscribe()
		if err != nil {
			log.Fatal("cannot unsubscribe")
		}
	}(subscription) //nolint

	var wc sync.WaitGroup
	wc.Add(1)

	subscription2, err := nc.Subscribe("sto.order", func(msg *nats.Msg) {
		fmt.Printf("LOG: \tgot message from subject: %s\n\tdata: %s\n", msg.Subject, string(msg.Data))
		message := strings.Split(string(msg.Data), " ")
		orderIDString := message[1]
		orderID, err := strconv.ParseUint(orderIDString, 10, 32)
		if err != nil {
			log.Fatal("cannot parse string to uint")
		}

		// Get Order
		orderRes, _ := getOrder(rdb, uint32(orderID))
		fmt.Println("Getting Order successful")
		// Reserve Products
		products := orderRes.GetProducts()
		fmt.Println("STOCK products: ", products)
		_, err = server.OrderProducts(context.Background(), &stockApi.OrderProductsRequest{OrderId: uint32(orderID), Products: products})
		if err != nil {
			log.Fatal("cannot reserve products")
		}
	})
	if err != nil {
		log.Fatal("cannot subscribe")
	}
	defer func(subscription2 *nats.Subscription) {
		err := subscription2.Unsubscribe()
		if err != nil {
			log.Fatal("cannot unsubscribe")
		}
	}(subscription2) //nolint

	wc.Add(2)

	subscription3, err := nc.Subscribe("sto.cancel", func(msg *nats.Msg) {
		fmt.Printf("LOG: \tgot message from subject: %s\n\tdata: %s\n", msg.Subject, string(msg.Data))
		message := strings.Split(string(msg.Data), " ")
		orderIDString := message[1]
		orderID, err := strconv.ParseUint(orderIDString, 10, 32)
		if err != nil {
			log.Fatal("cannot parse string to uint")
		}

		// Get Order
		orderRes, _ := getOrder(rdb, uint32(orderID))
		fmt.Println("Getting Order successful")

		// Reserve Products
		products := orderRes.GetProducts()
		fmt.Println("order products: ", products)
		// Restock products that was reserved for order
		for product, amount := range products {
			if server.orders[uint32(orderID)][product] == 0 {
				server.products[product].Amount += amount
			}
		}
	})
	if err != nil {
		log.Fatal("cannot subscribe")
	}
	defer func(subscription3 *nats.Subscription) {
		err := subscription3.Unsubscribe()
		if err != nil {
			log.Fatal("cannot unsubscribe")
		}
	}(subscription3) //nolint

	wc.Add(3)

	services.RegisterStockServiceServer(s, server)
	fmt.Println("creating stockApi service finished")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	wc.Wait()
}

// Helper
func generateUniqueProductID(products map[uint32]*types.Product) uint32 {
	productId := uint32(rand.Intn(1000))
	if len(products) == 0 {
		return productId
	}
	_, exists := products[productId]
	for exists {
		productId = uint32(rand.Intn(1000))
		_, exists = products[productId]
	}
	return productId
}

func getOrder(rdb *redis.Client, orderID uint32) (*types.Order, error) {
	// get orderRes
	orderAddress, err := rdb.Get(context.Background(), "service:orderApi").Result()
	if err != nil {
		log.Fatalf("error while trying to get the orderRes service address %v", err)
	}
	fmt.Println("orderAddress successful.")
	orderConn, err := grpc.Dial(orderAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to orderRes service: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while closing the connection to orderRes service %v", err)
		}
	}(orderConn)
	fmt.Println("orderConn successful.")

	orderClient := services.NewOrderServiceClient(orderConn)
	fmt.Println("orderClient successful.")
	orderRes, err := orderClient.GetOrder(context.Background(), &orderApi.GetOrderRequest{CustomerId: 0, OrderId: orderID})
	if err != nil {
		return nil, fmt.Errorf("error while trying to get orderRes %v", err)
	}
	return orderRes.GetOrder(), nil
}

func setOrderStatus(rdb *redis.Client, orderId uint32) bool {
	// set order status
	orderAddress, err := rdb.Get(context.Background(), "service:orderApi").Result()
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
	_, err = orderClient.SetOrderStatus(context.Background(), &orderApi.SetOrderStatusRequest{OrderId: orderId, Status: true})
	return err == nil
}
