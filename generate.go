//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/api/generalApi/generalApi.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/api/customerApi/customerApi.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/api/orderApi/orderApi.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/api/paymentApi/paymentApi.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/api/shipmentApi/shipmentApi.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/api/stockApi/stockApi.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/api/supplierApi/supplierApi.proto
package generate

// TODO: geht das Aufbauen der Verbindung auch einfacher?
// TODO: wie nutzen wir datentypen aus anderen services?
