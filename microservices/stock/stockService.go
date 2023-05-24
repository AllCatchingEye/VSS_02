package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/stockApi"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"reflect"
	"time"
)

type server struct {
	stockApi.StockServiceServer
	redis    *redis.Client
	nats     *nats.Conn
	products map[uint32]*stockApi.Product
}

func (state *server) AddProducts(ctx context.Context, req *stockApi.AddProductsRequest) (*stockApi.AddProductsReply, error) {
	fmt.Println("AddProduct called")
	fmt.Println(req.GetProducts())
	err := state.nats.Publish("log.stockApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.stockApi: cannot publish event")
	}
	newProducts := req.GetProducts()
	var productIDs []uint32
	for _, product := range newProducts {
		productID := generateUniqueProductID(state.products)
		productIDs = append(productIDs, productID)
		state.products[productID] = product
	}
	return &stockApi.AddProductsReply{ProductIds: productIDs}, nil
}

func (state *server) GetProducts(ctx context.Context, req *stockApi.GetProductsRequest) (*stockApi.GetProductsReply, error) {
	fmt.Println("GetProduct called")
	fmt.Println(req.GetProductIds())
	err := state.nats.Publish("log.stockApi", []byte(fmt.Sprintf("got message %v", reflect.TypeOf(req))))
	if err != nil {
		log.Print("log.stockApi: cannot publish event")
	}
	productIDs := req.GetProductIds()
	var products []*stockApi.Product
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

func main() {
	flagHost := flag.String("host", "127.0.0.1", "address of customerApi service")
	flagPort := flag.String("port", "50055", "port of customerApi service")
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
			rdb.Set(context.TODO(), "service:stockApi", address, 13*time.Second)
			time.Sleep(10 * time.Second)
		}
	}()

	nc, err := nats.Connect(*flagNATS)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	stockApi.RegisterStockServiceServer(s, &server{redis: rdb, nats: nc, products: make(map[uint32]*stockApi.Product)})
	fmt.Println("creating stockApi service finished")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Helper
func generateUniqueProductID(products map[uint32]*stockApi.Product) uint32 {
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
