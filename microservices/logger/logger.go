package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	flagNATS := flag.String("nats", "127.0.0.1:4222", "address and port of NATS server")
	flag.Parse()

	time.Sleep(5 * time.Second)

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
	wc.Wait()
}
