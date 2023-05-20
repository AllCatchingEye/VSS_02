//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/customer/api/api.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/shipment/api/api.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/stock/api/api.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/supplier/api/api.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/payment/api/api.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=.  --go-grpc_opt=paths=source_relative microservices/order/api/api.proto
package generate

// TODO: geht das Aufbauen der Verbindung auch einfacher?
// TODO: wie nutzen wir datentypen aus anderen services?
