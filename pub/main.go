package main

import (
	"time"

	nats "github.com/nats-io/go-nats"
)

func main() {
	nc, _ := nats.Connect("0.0.0.0:4222")

	nc.Publish("test", []byte("message"))

	time.Sleep(2 * time.Second)

	nc.Publish("test", []byte("message2"))

	time.Sleep(1 * time.Second)

	nc.Publish("chn", []byte("ok"))

	nc.Drain()
	nc.Close()
}
