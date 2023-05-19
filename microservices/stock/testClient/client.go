package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	stockApi "gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/stock/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	flagRedis := flag.String("redis", "127.0.0.1:6379", "address and port of Redis server")
	flag.Parse()

	rdb := redis.NewClient(&redis.Options{
		Addr:     *flagRedis,
		Password: "",
	})

	address, err := rdb.Get(context.TODO(), "service:stock").Result()
	if err != nil {
		log.Fatalf("error while trying to get the stock service address: %v", err)
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

	c := stockApi.NewStockServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// Adding Products
	product1 := &stockApi.Product{
		Name:        "Apple",
		Description: "A red and sweet apple",
		Price:       1.99,
		Supplier:    1,
	}
	product2 := &stockApi.Product{
		Name:        "Banana",
		Description: "A yellow and sweet banana",
		Price:       0.99,
		Supplier:    1,
	}
	product3 := &stockApi.Product{
		Name:        "Spoon",
		Description: "A silver spoon",
		Price:       5.99,
		Supplier:    2,
	}
	products := []*stockApi.Product{product1, product2, product3}
	r, err := c.AddProducts(ctx, &stockApi.AddProductsRequest{Products: products})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	log.Printf("Getting following products:")
	for _, product := range r.GetProductIds() {
		log.Printf("Product id: %d", product)
	}

	// Getting Customer
	productIds := []uint32{r.GetProductIds()[0], r.GetProductIds()[1]}
	r2, err := c.GetProducts(ctx, &stockApi.GetProductsRequest{ProductIds: productIds})
	if err != nil {
		log.Fatalf("could not get products: %v", err)
	}
	log.Printf("Greeting products")
	for _, product := range r2.GetProducts() {
		fmt.Printf("Product: %s\n", product.GetName())
	}

	// Remove Customer Apple
	r3, err := c.RemoveProduct(ctx, &stockApi.RemoveProductRequest{ProductId: productIds[0]})
	if err != nil {
		log.Fatalf("could not remove product: %v", err)
	}
	log.Printf("Removed: %s", r3.GetProduct().GetName())

}
