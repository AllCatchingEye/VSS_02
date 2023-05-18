//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/customer/api/api.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/shipment/api/api.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/order/api/api.proto
package generate
