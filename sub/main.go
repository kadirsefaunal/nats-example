package main

import (
	"fmt"

	nats "github.com/nats-io/go-nats"
)

func main() {
	nc, _ := nats.Connect("0.0.0.0:4222")

	nc.Subscribe("test", func(m *nats.Msg) {
		fmt.Printf("%s\n", string(m.Data))
	})

	ch := make(chan *nats.Msg, 64)
	sub, _ := nc.ChanSubscribe("chn", ch)
	msg := <-ch

	fmt.Printf("%s\n", string(msg.Data))

	sub.Unsubscribe()
	sub.Drain()

	nc.Drain()
	nc.Close()
}
