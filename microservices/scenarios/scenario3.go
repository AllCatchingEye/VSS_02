package scenarios

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
)

func Scene3(rdb *redis.Client, nc *nats.Conn) bool {
	// TODO implement
	fmt.Println("having nats ", nc.IsConnected())
	fmt.Println("having redis ", rdb)
	return true
}
