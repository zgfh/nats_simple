package main

import (
	"fmt"
	"github.com/nats-io/nats"
	"time"
	"strconv"
)


func main()  {
	fmt.Println("-----------test nats------------------- ")
	go sub()
	go pub()
	select {
	}
}
func pub()  {
	nc, _ := nats.Connect("nats://root:dangerous@localhost:4222")

	for i:=1;i<10;i++ {
		fmt.Println("publish msg :"+strconv.Itoa(i))
		nc.Publish("foo", []byte("Hello World"+strconv.Itoa(i)))
		time.Sleep(3 * time.Second)
	}

}
func sub()  {
	nc, _ := nats.Connect("nats://root:dangerous@localhost:4222")
	nc.Subscribe("foo", func(m *nats.Msg) {
    		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

}
