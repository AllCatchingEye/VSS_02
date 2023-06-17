package main

import (
	"flag"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/scenarios"
	"log"
	"os"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)

	flagRedis := flag.String("redis", "127.0.0.1:6379", "customerAddress and port of Redis server")
	flagNATS := flag.String("nats", "127.0.0.1:4222", "customerAddress and port of NATS server")
	flag.Parse()

	time.Sleep(5 * time.Second)

	rdb := redis.NewClient(&redis.Options{
		Addr:     *flagRedis,
		Password: "",
	})

	nc, err := nats.Connect(*flagNATS)
	if err != nil {
		log.Fatal("cannot connect to nats")
	}
	defer nc.Close()

	// run scenario1
	scene1 := scenarios.Scene1(rdb, nc)
	if scene1 {
		log.Println("Scenario 1 successful")
	} else {
		log.Println("Scenario 1 failed")
	}
	// run scenario2
	scene2 := scenarios.Scene2(rdb, nc)
	if scene2 {
		log.Println("Scenario 2 successful")
	} else {
		log.Println("Scenario 2 failed")
	}
	// run scenario3
	scene3 := scenarios.Scene3(rdb, nc)
	if scene3 {
		log.Println("Scenario 3 successful")
	} else {
		log.Println("Scenario 3 failed")
	}
	// run scenario4
	scene4 := scenarios.Scene4(rdb, nc)
	if scene4 {
		log.Println("Scenario 4 successful")
	} else {
		log.Println("Scenario 4 failed")
	}
	// run scenario5
	scene5 := scenarios.Scene5(rdb, nc)
	if scene5 {
		log.Println("Scenario 5 successful")
	} else {
		log.Println("Scenario 5 failed")
	}

	os.Exit(0)
}
