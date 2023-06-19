package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/scenarios"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("starting client")
	time.Sleep(10 * time.Second)

	scenario := os.Getenv("SCENARIO")

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

	switch scenario {
	case "1":
		// run scenario1
		fmt.Println("##########################################")
		fmt.Println("########## Running Scenario 1 ############")
		fmt.Println("##########################################")
		scene1 := scenarios.Scene1(rdb, nc)
		if scene1 {
			log.Println("Scenario 1 successful")
		} else {
			log.Println("Scenario 1 failed")
		}
	case "2":
		// run scenario2
		fmt.Println("##########################################")
		fmt.Println("########## Running Scenario 2 ############")
		fmt.Println("##########################################")
		scene2 := scenarios.Scene2(rdb, nc)
		if scene2 {
			log.Println("Scenario 2 successful")
		} else {
			log.Println("Scenario 2 failed")
		}
	case "3":
		// run scenario3
		fmt.Println("##########################################")
		fmt.Println("########## Running Scenario 3 ############")
		fmt.Println("##########################################")
		scene3 := scenarios.Scene3(rdb, nc)
		if scene3 {
			log.Println("Scenario 3 successful")
		} else {
			log.Println("Scenario 3 failed")
		}
	case "4":
		// run scenario4
		fmt.Println("##########################################")
		fmt.Println("########## Running Scenario 4 ############")
		fmt.Println("##########################################")
		scene4 := scenarios.Scene4(rdb, nc)
		if scene4 {
			log.Println("Scenario 4 successful")
		} else {
			log.Println("Scenario 4 failed")
		}
	case "5":
		// run scenario5
		fmt.Println("##########################################")
		fmt.Println("########## Running Scenario 5 ############")
		fmt.Println("##########################################")
		scene5 := scenarios.Scene5(rdb, nc)
		if scene5 {
			log.Println("Scenario 5 successful")
		} else {
			log.Println("Scenario 5 failed")
		}
	}

	os.Exit(0)
}
