package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"gonotes/rpcx"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:999", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &rpcx.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &rpcx.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(time.Second)
	}

}
