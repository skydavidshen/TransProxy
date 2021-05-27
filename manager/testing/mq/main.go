package main

import (
	"TransProxy/utils"
	"fmt"
	"log"
	"time"
)

type Item struct {
	Value string
	Num   int
	Mod   int
}

var c chan Item
var m map[int]int

func main() {
	m = make(map[int]int)

	go insert()
	time.Sleep(time.Second)
	go get()

	go func() {
		for {
			fmt.Println(m)
			println()
			time.Sleep(time.Second * 3)
		}
	}()
	b := make(chan Item)
	<- b
}

func insert() {
	log.Println("insert ...")
	c = make(chan Item)
	for i := 0; i < 1000; i = i + 1 {
		log.Println("insert for loop...")
		val := utils.GetRandomString(8)
		item := Item{
			Value: val,
			Mod:   i % 10,
			Num:   i,
		}
		c <- item
		time.Sleep(time.Second)
	}
	close(c)
}

func get() {
	log.Println("get ...")
	for i := range c {
		if _, ok := m[i.Mod]; !ok {
			m[i.Mod] = 0
		} else {
			m[i.Mod]++
		}
	}
}