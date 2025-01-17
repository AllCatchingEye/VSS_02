package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/services"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/supplierApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/types"
	"google.golang.org/grpc"
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
	services.SupplierServiceServer
	redis    *redis.Client
	nats     *nats.Conn
	supplier map[uint32]*types.Supplier
}

func (state *server) AddSupplier(ctx context.Context, req *supplierApi.AddSupplierRequest) (*supplierApi.AddSupplierReply, error) {
	fmt.Println("AddSupplier called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req.GetSupplier())
	err := state.nats.Publish("log.supplierApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.supplierApi: cannot publish event")
	}
	newSupplier := req.GetSupplier()
	supplierID := newSupplier.GetSupplierId()
	if supplierID == 0 {
		supplierID = generateUniqueSupplierID(state.supplier)
	}
	state.supplier[supplierID] = newSupplier
	return &supplierApi.AddSupplierReply{SupplierId: supplierID}, nil
}

func (state *server) GetSupplier(ctx context.Context, req *supplierApi.GetSupplierRequest) (*supplierApi.GetSupplierReply, error) {
	fmt.Println("GetSupplier called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req.GetSupplierId())
	err := state.nats.Publish("log.supplierApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.supplierApi: cannot publish event")
	}
	supplierId := req.GetSupplierId()
	supplier, ok := state.supplier[supplierId]
	if !ok {
		return nil, fmt.Errorf("supplier to get not found")
	}
	return &supplierApi.GetSupplierReply{Supplier: supplier}, nil
}

func (state *server) RemoveSupplier(ctx context.Context, req *supplierApi.RemoveSupplierRequest) (*supplierApi.RemoveSupplierReply, error) {
	fmt.Println("RemoveSupplier called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req.GetSupplierId())
	err := state.nats.Publish("log.supplierApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.supplierApi: cannot publish event")
	}
	supplierId := req.GetSupplierId()
	supplier, ok := state.supplier[supplierId]
	if !ok {
		return nil, fmt.Errorf("supplier to remove not found")
	}
	delete(state.supplier, supplierId)
	return &supplierApi.RemoveSupplierReply{Supplier: supplier}, nil
}

func (state *server) AddProducts(ctx context.Context, req *supplierApi.AddProductsRequest) (*supplierApi.AddProductsReply, error) {
	fmt.Println("AddProductsRequest called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println("SupplierId: ", req.GetSupplierId())
	err := state.nats.Publish("log.supplierApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.supplierApi: cannot publish event")
	}
	fmt.Println("Nats logs finished.")
	supplierId := req.GetSupplierId()
	supplier, ok := state.supplier[supplierId]
	if !ok {
		return nil, fmt.Errorf("supplier to add products to not found")
	}
	supplier.Products = append(supplier.Products, req.GetProducts()...)
	fmt.Println("Supplier: adding products succeeded.")
	return &supplierApi.AddProductsReply{Supplier: supplier}, nil
}

func (state *server) RemoveProducts(ctx context.Context, req *supplierApi.RemoveProductsRequest) (*supplierApi.RemoveProductsReply, error) {
	fmt.Println("RemoveProductsRequest called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req.GetSupplierId())
	supplierId := req.GetSupplierId()
	err := state.nats.Publish("log.supplierApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.supplierApi: cannot publish event")
	}
	supplier, ok := state.supplier[supplierId]
	if !ok {
		return nil, fmt.Errorf("supplier to remove products not found")
	}
	for _, productToRemove := range req.GetProducts() {
		for i, productOfSupplier := range supplier.Products {
			if productOfSupplier.GetProductId() == productToRemove {
				supplier.Products = append(supplier.Products[:i], supplier.Products[i+1:]...)
				break
			}
		}
	}
	return &supplierApi.RemoveProductsReply{Supplier: supplier}, nil
}

/**
 * OrderProduct deprecated
 */
func (state *server) OrderProduct(ctx context.Context, req *supplierApi.OrderProductRequest) (*supplierApi.OrderProductReply, error) {
	fmt.Println("OrderProduct called")
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("context deadline is ", deadline)
	}
	fmt.Println(req.GetSupplierId())
	err := state.nats.Publish("log.supplierApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.supplierApi: cannot publish event")
	}
	supplierId := req.GetSupplierId()
	supplier, ok := state.supplier[supplierId]
	if !ok {
		return nil, fmt.Errorf("supplier to order products from not found")
	}
	productToOrder := req.GetProductId()
	amount := req.GetAmount()
	// find product id in products of supplierApi, if not found return error
	found := false
	for _, productOfSupplier := range supplier.Products {
		if productOfSupplier.GetProductId() == productToOrder {
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("product not found")
	}
	time.Sleep(10 * time.Second)
	// TODO: send response to stock api (asyncon)
	return &supplierApi.OrderProductReply{ProductId: productToOrder, Amount: amount}, nil
}

func main() {
	flagHost := flag.String("host", "127.0.0.1", "address of supplierApi service")
	flagPort := flag.String("port", "50056", "port of supplierApi service")
	flagNATS := flag.String("nats", "127.0.0.1:4222", "address and port of NATS server")
	flagRedis := flag.String("redis", "127.0.0.1:6379", "address and port of Redis server")
	flag.Parse()

	time.Sleep(5 * time.Second)

	nc, err := nats.Connect(*flagNATS)
	if err != nil {
		log.Fatal("cannot connect to nats")
	}
	defer nc.Close()

	subscription, err := nc.Subscribe("supp.orderProduct", func(msg *nats.Msg) {
		fmt.Printf("LOG: \tgot message from subject: %s\n\tdata: %s\n", msg.Subject, string(msg.Data))
		message := strings.Split(string(msg.Data), " ")
		supplier := message[1]
		product := message[2]
		amount := message[3]
		fmt.Printf("supplier: %s, product: %s, amount: %s\n", supplier, product, amount)
		supplierID, _ := strconv.ParseUint(supplier, 10, 32)
		productID, _ := strconv.ParseUint(product, 10, 32)
		amountUint, _ := strconv.ParseUint(amount, 10, 32)
		time.Sleep(3 * time.Second)
		sendOrderedProductsToStockApi(nc, uint32(supplierID), uint32(productID), uint32(amountUint))
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
		fmt.Println("starting to update redis for supplierApi service")
		for {
			rdb.Set(context.TODO(), "service:supplierApi", address, 13*time.Second)
			time.Sleep(10 * time.Second)
		}
	}()

	services.RegisterSupplierServiceServer(s, &server{redis: rdb, nats: nc, supplier: make(map[uint32]*types.Supplier)})
	fmt.Println("creating supplierApi service finished")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	wc.Wait()
}

// Helper
func generateUniqueSupplierID(supplier map[uint32]*types.Supplier) uint32 {
	supplierId := uint32(rand.Intn(1000))
	if len(supplier) == 0 {
		return supplierId
	}
	_, exists := supplier[supplierId]
	for exists {
		supplierId = uint32(rand.Intn(1000))
		_, exists = supplier[supplierId]
	}
	return supplierId
}

func sendOrderedProductsToStockApi(nc *nats.Conn, supplierId uint32, product uint32, amount uint32) {
	err := nc.Publish("supp.deliverProduct", []byte(fmt.Sprintf("%d %d %d", supplierId, product, amount)))
	if err != nil {
		log.Fatal("cannot publish message to subject supp.deliverProduct")
	}
}
