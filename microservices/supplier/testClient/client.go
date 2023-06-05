package main

import (
	"context"
	"flag"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/services"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/supplierApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/types"
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

	address, err := rdb.Get(context.TODO(), "service:supplierApi").Result()
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

	c := services.NewSupplierServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*25)
	defer cancel()
	// Adding Supplier
	supplier := &types.Supplier{
		Name: "Obstlieferant",
		Address: &types.Address{
			Street:  "Obststraße",
			Zip:     "12345",
			City:    "Obsthausen",
			Country: "Obstland",
		},
		Products: []uint32{1, 2},
	}
	r, err := c.AddSupplier(ctx, &supplierApi.AddSupplierRequest{Supplier: supplier})
	if err != nil {
		log.Fatalf("could add supplierApi: %v", err)
	}
	log.Printf("Getting id: %d", r.GetSupplierId())

	// Getting Supplier
	r2, err := c.GetSupplier(ctx, &supplierApi.GetSupplierRequest{SupplierId: r.GetSupplierId()})
	if err != nil {
		log.Fatalf("could get supplierApi: %v", err)
	}
	log.Printf("Getting supplierApi: %v", r2.GetSupplier().GetName())

	// Add Products
	r3, err := c.AddProducts(ctx, &supplierApi.AddProductsRequest{SupplierId: r.GetSupplierId(), Products: []uint32{3, 4}})
	if err != nil {
		log.Fatalf("could add products: %v", err)
	}
	log.Printf("Removed: %#v", r3.GetSupplier().GetProducts())

	// Remove Products
	r4, err := c.RemoveProducts(ctx, &supplierApi.RemoveProductsRequest{SupplierId: r.GetSupplierId(), Products: []uint32{1}})
	if err != nil {
		log.Fatalf("could remove products: %v", err)
	}
	log.Printf("Removed products. These are remaining: %#v", r4.GetSupplier().GetProducts())

	// Remove Supplier
	r5, err := c.RemoveSupplier(ctx, &supplierApi.RemoveSupplierRequest{SupplierId: r.GetSupplierId()})
	if err != nil {
		log.Fatalf("could remove supplierApi: %v", err)
	}
	log.Printf("Removed supplierApi: %v", r5.GetSupplier())

	log.Printf("Done")
}
