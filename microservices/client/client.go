package main

import (
	"gitlab.lrz.de/vss/semester/ob-23ss/blatt-2/blatt2-grp06/microservices/scenarios"
	"log"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)
	// run scenario1
	scene1 := scenarios.Scene1()
	if scene1 {
		log.Println("Scenario 1 successful")
	} else {
		log.Println("Scenario 1 failed")
	}
	// run scenario2
	scene2 := scenarios.Scene2()
	if scene2 {
		log.Println("Scenario 2 successful")
	} else {
		log.Println("Scenario 2 failed")
	}
	// run scenario3
	scene3 := scenarios.Scene3()
	if scene3 {
		log.Println("Scenario 3 successful")
	} else {
		log.Println("Scenario 3 failed")
	}
	// run scenario4
	scene4 := scenarios.Scene4()
	if scene4 {
		log.Println("Scenario 4 successful")
	} else {
		log.Println("Scenario 4 failed")
	}
	// run scenario5
	scene5 := scenarios.Scene5()
	if scene5 {
		log.Println("Scenario 5 successful")
	} else {
		log.Println("Scenario 5 failed")
	}
}
