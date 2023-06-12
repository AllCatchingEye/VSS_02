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
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/shipmentApi"
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

	nc, err := nats.Connect(*flagNATS)
	if err != nil {
		log.Fatal("cannot connect to nats")
	}
	defer nc.Close()

	fillStore(rdb, nc)

	// build connection to customer service
	customerAddress, err := rdb.Get(context.TODO(), "service:customerApi").Result()
	if err != nil {
		log.Fatalf("error while trying to get the customer service address %v", err)
	}

	customerConn, err := grpc.Dial(customerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while closing the connection %v", err)
		}
	}(customerConn)

	customerClient := services.NewCustomerServiceClient(customerConn)

	// build connection to order service
	orderAddress, err := rdb.Get(context.TODO(), "service:orderApi").Result()
	if err != nil {
		log.Fatalf("error while trying to get the newOrder service address %v", err)
	}

	orderConn, err := grpc.Dial(orderAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to newOrder service: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while closing the connection to newOrder service %v", err)
		}
	}(orderConn)

	orderClient := services.NewOrderServiceClient(orderConn)

	// build connection to payment service
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

	// build connection to payment service
	shipmentAddress, err := rdb.Get(context.TODO(), "service:shipmentApi").Result()
	if err != nil {
		log.Fatalf("error while trying to get the shipment service address %v", err)
	}

	shipmentConn, err := grpc.Dial(shipmentAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to shipment service: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("error while closing the connection to shipment service %v", err)
		}
	}(shipmentConn)

	shipmentClient := services.NewShipmentServiceClient(shipmentConn)

	// set context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	log.Printf("Context: %v", ctx)

	// Add Customer
	customer := &types.Customer{
		Name: "Max Mustermann",
		Address: &types.Address{
			Street:  "Mustermannstraße 42",
			Zip:     "80335",
			City:    "Munich",
			Country: "Germany",
		},
	}
	rCustomer, err := customerClient.AddCustomer(ctx, &customerApi.AddCustomerRequest{Customer: customer})
	if err != nil {
		log.Fatalf("could not add customer %v: %v", customer.GetName(), err)
	}
	log.Printf("Getting id: %v", rCustomer.GetCustomerId())
	customerID := rCustomer.GetCustomerId()

	// New Order
	newOrder := make(map[uint32]uint32)
	newOrder[111] = 5
	newOrder[222] = 10
	rOrder, err := orderClient.NewOrder(ctx, &orderApi.NewOrderRequest{CustomerId: customerID, Products: newOrder})
	if err != nil {
		log.Fatalf("could not create newOrder %v: %v", newOrder, err)
	}
	log.Printf("Getting id: %v", rOrder.GetOrderId())
	order := rOrder.GetOrder()
	log.Printf("Order: %v", order)
	orderID := rOrder.GetOrderId()
	log.Printf("OrderID: %v", orderID)

	log.Printf("Orderstatus: %v", order.GetOrderStatus())
	log.Printf("Paymentstatus: %v", order.GetPaymentStatus())
	log.Printf("Deliverystatus: %v", order.GetDeliveryStatus())

	// Pay Order
	rPayment, err := paymentClient.PayMyOrder(ctx, &paymentApi.PayMyOrderRequest{
		CustomerId: customerID,
		OrderId:    orderID,
	})
	if err != nil {
		log.Fatalf("could not pay order %v: %v", orderID, err)
	}
	log.Printf("Payed order with ID: %v", rPayment.GetOrderId())

	// Check if order is payed
	rPayment2, err := paymentClient.IsOrderPayed(ctx, &paymentApi.IsOrderPayedRequest{CustomerId: customerID, OrderId: orderID})
	if err != nil {
		log.Fatalf("could not check if order is payed %v: %v", orderID, err)
	}
	log.Printf("Order %v is payed: %v", orderID, rPayment2.GetIsPayed())

	// Get Order
	rOrder2, err := orderClient.GetOrder(ctx, &orderApi.GetOrderRequest{CustomerId: customerID, OrderId: orderID})
	if err != nil {
		log.Fatalf("could not get order %v: %v", orderID, err)
	}
	log.Printf("Getting order: %v", rOrder2.GetOrder())
	order = rOrder2.GetOrder()
	log.Printf("Order: %v", order)
	orderID = rOrder2.GetOrderId()
	log.Printf("OrderID: %v", orderID)

	log.Printf("Orderstatus: %v", order.GetOrderStatus())
	log.Printf("Paymentstatus: %v", order.GetPaymentStatus())
	log.Printf("Deliverystatus: %v", order.GetDeliveryStatus())

	// Set Delivery Status
	rShipment, err := shipmentClient.ShipMyOrder(ctx, &shipmentApi.ShipMyOrderRequest{CustomerId: customerID, OrderId: orderID})
	if err != nil {
		log.Fatalf("could not ship order %v: %v", orderID, err)
	}
	log.Printf("Shipped order with ID: %v", rShipment.GetOrderId())

	// Check if order is shipped
	rShipment2, err := shipmentClient.IsOrderShipped(ctx, &shipmentApi.IsOrderShippedRequest{CustomerId: customerID, OrderId: orderID})
	if err != nil {
		log.Fatalf("could not check if order %v is shipped: %v", orderID, err)
	}
	log.Printf("Order with ID: %v is shipped: %v", orderID, rShipment2.GetIsShipped())

	// Get Order
	rOrder3, err := orderClient.GetOrder(ctx, &orderApi.GetOrderRequest{CustomerId: customerID, OrderId: orderID})
	if err != nil {
		log.Fatalf("could not get order %v: %v", orderID, err)
	}
	log.Printf("Getting order: %v", rOrder2.GetOrder())
	order = rOrder3.GetOrder()
	log.Printf("Order: %v", order)
	orderID = rOrder3.GetOrderId()
	log.Printf("OrderID: %v", orderID)

	log.Printf("Orderstatus: %v", order.GetOrderStatus())
	log.Printf("Paymentstatus: %v", order.GetPaymentStatus())
	log.Printf("Deliverystatus: %v", order.GetDeliveryStatus())

	fmt.Println("Done")
}

func fillStore(rdb *redis.Client, nc *nats.Conn) {
	// set context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// subscribe to nats logging

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
