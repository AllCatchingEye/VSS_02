package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/services"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/stockApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/supplierApi"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
	"time"
)

func main() {
	flagRedis := flag.String("redis", "127.0.0.1:6379", "customerAddress and port of Redis server")
	flagNATS := flag.String("nats", "127.0.0.1:4222", "customerAddress and port of NATS server")
	flag.Parse()

	rdb := redis.NewClient(&redis.Options{
		Addr:     *flagRedis,
		Password: "",
	})

	// set context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// subscribe to nats logging
	nc, err := nats.Connect(*flagNATS)
	if err != nil {
		log.Fatal("cannot connect to nats")
	}
	defer nc.Close()

	subscription, err := nc.Subscribe("log.*", func(msg *nats.Msg) {
		fmt.Printf("LOG: \tgot message from subject: %s\n\tdata: %s\n", msg.Subject, string(msg.Data))
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

	// ********** Supplier Service **********
	// build connection to supplier service
	supplierAddress, err := rdb.Get(context.TODO(), "service:supplierApi").Result()
	if err != nil {
		log.Fatalf("error while trying to get the result %v", err)
	}

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

	supplierClient := services.NewSupplierServiceClient(supplierConn)

	suppliers := map[uint32]*types.Supplier{
		1: {
			SupplierId: 1,
			Name:       "Supplier 1",
			Address: &types.Address{
				Street:  "Supplierstraße 1",
				Zip:     "12345",
				City:    "Suppcity",
				Country: "Suppland",
			},
			Products: []*types.Product{},
		},
		2: {
			SupplierId: 2,
			Name:       "Supplier 2",
			Address: &types.Address{
				Street:  "Adilettenstraße 3",
				Zip:     "11111",
				City:    "Adistadt",
				Country: "Adidanien",
			},
			Products: []*types.Product{},
		},
		3: {
			SupplierId: 3,
			Name:       "Supplier 3",
			Address: &types.Address{
				Street:  "Obstplatz 2",
				Zip:     "23232",
				City:    "Obsthausen",
				Country: "Obstalien",
			},
			Products: []*types.Product{},
		},
	}

	supplierRes, err := supplierClient.AddSupplier(ctx, &supplierApi.AddSupplierRequest{Supplier: suppliers[1]})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	fmt.Printf("RES: \tGetting id: %v\n", supplierRes.GetSupplierId())
	supplierID := supplierRes.GetSupplierId()

	supplierRes, err = supplierClient.AddSupplier(ctx, &supplierApi.AddSupplierRequest{Supplier: suppliers[2]})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	fmt.Printf("RES: \tGetting id: %v\n", supplierRes.GetSupplierId())
	supplierID2 := supplierRes.GetSupplierId()

	supplierRes, err = supplierClient.AddSupplier(ctx, &supplierApi.AddSupplierRequest{Supplier: suppliers[3]})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	fmt.Printf("RES: \tGetting id: %v\n", supplierRes.GetSupplierId())
	supplierID3 := supplierRes.GetSupplierId()

	// ********** Stock Service **********
	// build connection to stock service
	stockAddress, err := rdb.Get(context.TODO(), "service:stockApi").Result()
	if err != nil {
		log.Fatalf("error while trying to get the result %v", err)
	}

	stockConn, err := grpc.Dial(stockAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while closing the connection %v", err)
		}
	}(stockConn)

	stockClient := services.NewStockServiceClient(stockConn)

	tecProducts := []*types.Product{
		{
			Name:        "Laptop",
			Description: "A laptop, very useful for working",
			Price:       1000,
			Supplier:    supplierID,
		},
		{
			Name:        "Headset",
			Description: "A headset, very useful for listening",
			Price:       150,
			Supplier:    supplierID,
		},
	}

	sportswearProducts := []*types.Product{
		{
			Name:        "Adiletten",
			Description: "Adiletten, very useful for walking",
			Price:       20,
			Supplier:    supplierID2,
		},
		{
			Name:        "Jogger",
			Description: "Jogger, very useful for running",
			Price:       50,
			Supplier:    supplierID2,
		},
		{
			Name:        "T-Shirt",
			Description: "T-Shirt, very useful for sweating",
			Price:       10,
			Supplier:    supplierID2,
		},
	}

	fruitProducts := []*types.Product{
		{
			ProductId:   111,
			Name:        "Apple",
			Description: "Apple, keeps the doctor away",
			Price:       1,
			Supplier:    supplierID3,
		},
		{
			ProductId:   222,
			Name:        "Banana",
			Description: "Banana, very useful for monkeys",
			Price:       2,
			Supplier:    supplierID3,
		},
		{
			ProductId:   333,
			Name:        "Orange",
			Description: "Orange, very orange",
			Price:       1,
			Supplier:    supplierID3,
		},
	}

	// add products to stock
	stockRes, err := stockClient.AddProducts(ctx, &stockApi.AddProductsRequest{Products: tecProducts})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	tecProductsIds := stockRes.GetProductIds()
	fmt.Printf("RES: \tGetting id: %v\n", tecProductsIds)

	stockRes, err = stockClient.AddProducts(ctx, &stockApi.AddProductsRequest{Products: sportswearProducts})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	sportswearProductsIds := stockRes.GetProductIds()
	fmt.Printf("RES: \tGetting id: %v\n", sportswearProductsIds)

	stockRes, err = stockClient.AddProducts(ctx, &stockApi.AddProductsRequest{Products: fruitProducts})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	fruitProductsIds := stockRes.GetProductIds()
	fmt.Printf("RES: \tGetting id: %v\n", fruitProductsIds)

}
