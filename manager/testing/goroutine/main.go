package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int)
	go receive(ch)
	go get(ch)

	b := make(chan bool)
	<-b
	//time.Sleep(time.Second * 60)

	log.Println("main done...")
}

func receive(ch chan int) {
	log.Println("start receive...")
	i := 1
	for {
		log.Println("before set ch...")
		ch <-i
		log.Println("after receive...")
		time.Sleep(time.Second * 5)
		log.Println("after 5 receive...")
		i = i +10
	}
}

func get(ch chan int) {
	log.Println("start get...")
	for {
		select {
		case i := <- ch:
			log.Printf("=====================")
			log.Printf("get ch: %d", i)
		}
	}
	log.Println("get done...")
}
