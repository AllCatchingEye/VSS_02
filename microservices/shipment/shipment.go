package shipment

import "gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/shipment/api"

type server struct {
	api.UnimplementedShipmentServiceServer
}

func (*server) ShipmentOrder() {

}

func main() {

}
