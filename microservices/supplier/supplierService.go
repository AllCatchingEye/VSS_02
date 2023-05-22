package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/api/supplierApi"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
)

type server struct {
	supplierApi.SupplierServiceServer
	supplier map[uint32]*supplierApi.Supplier
}

func (state *server) AddSupplier(ctx context.Context, req *supplierApi.AddSupplierRequest) (*supplierApi.AddSupplierReply, error) {
	fmt.Println("AddSupplier called")
	fmt.Println(req.GetSupplier())
	newSupplier := req.GetSupplier()
	supplierID := generateUniqueSupplierID(state.supplier)
	state.supplier[supplierID] = newSupplier
	return &supplierApi.AddSupplierReply{SupplierId: supplierID}, nil
}

func (state *server) GetSupplier(ctx context.Context, req *supplierApi.GetSupplierRequest) (*supplierApi.GetSupplierReply, error) {
	fmt.Println("GetSupplier called")
	fmt.Println(req.GetSupplierId())
	supplierId := req.GetSupplierId()
	supplier, ok := state.supplier[supplierId]
	if !ok {
		return nil, fmt.Errorf("supplierApi not found")
	}
	return &supplierApi.GetSupplierReply{Supplier: supplier}, nil
}

func (state *server) RemoveSupplier(ctx context.Context, req *supplierApi.RemoveSupplierRequest) (*supplierApi.RemoveSupplierReply, error) {
	fmt.Println("RemoveSupplier called")
	fmt.Println(req.GetSupplierId())
	supplierId := req.GetSupplierId()
	supplier, ok := state.supplier[supplierId]
	if !ok {
		return nil, fmt.Errorf("supplierApi not found")
	}
	delete(state.supplier, supplierId)
	return &supplierApi.RemoveSupplierReply{Supplier: supplier}, nil
}

func (state *server) AddProducts(ctx context.Context, req *supplierApi.AddProductsRequest) (*supplierApi.AddProductsReply, error) {
	fmt.Println("AddProductsRequest called")
	fmt.Println(req.GetSupplierId())
	supplierId := req.GetSupplierId()
	supplier, ok := state.supplier[supplierId]
	if !ok {
		return nil, fmt.Errorf("supplierApi not found")
	}
	for _, productToAdd := range req.GetProducts() {
		supplier.Products = append(supplier.Products, productToAdd)
	}
	//supplierApi.Products = append(supplierApi.Products, req.GetProducts()...)
	return &supplierApi.AddProductsReply{Supplier: supplier}, nil
}

func (state *server) RemoveProducts(ctx context.Context, req *supplierApi.RemoveProductsRequest) (*supplierApi.RemoveProductsReply, error) {
	fmt.Println("RemoveProductsRequest called")
	fmt.Println(req.GetSupplierId())
	supplierId := req.GetSupplierId()
	supplier, ok := state.supplier[supplierId]
	if !ok {
		return nil, fmt.Errorf("supplierApi not found")
	}
	for _, productToRemove := range req.GetProducts() {
		for i, productOfSupplier := range supplier.Products {
			if productOfSupplier == productToRemove {
				supplier.Products = append(supplier.Products[:i], supplier.Products[i+1:]...)
				break
			}
		}
	}
	return &supplierApi.RemoveProductsReply{Supplier: supplier}, nil
}

func (state *server) OrderProduct(ctx context.Context, req *supplierApi.OrderProductRequest) (*supplierApi.OrderProductReply, error) {
	fmt.Println("OrderProduct called")
	fmt.Println(req.GetSupplierId())
	supplierId := req.GetSupplierId()
	supplier, ok := state.supplier[supplierId]
	if !ok {
		return nil, fmt.Errorf("supplierApi not found")
	}
	productToOrder := req.GetProductId()
	amount := req.GetAmount()
	// find product id in products of supplierApi, if not found return error
	found := false
	for _, productOfSupplier := range supplier.Products {
		if productOfSupplier == productToOrder {
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("product not found")
	}
	time.Sleep(10 * time.Second)
	return &supplierApi.OrderProductReply{ProductId: productToOrder, Amount: amount}, nil
}

func main() {
	flagHost := flag.String("host", "127.0.0.1", "address of supplierApi service")
	flagPort := flag.String("port", "50056", "port of supplierApi service")
	flagRedis := flag.String("redis", "127.0.0.1:6379", "address and port of Redis server")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *flagHost, *flagPort)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	supplierApi.RegisterSupplierServiceServer(s, &server{supplier: make(map[uint32]*supplierApi.Supplier)})
	fmt.Println("creating supplierApi service finished")

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

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Helper
func generateUniqueSupplierID(supplier map[uint32]*supplierApi.Supplier) uint32 {
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