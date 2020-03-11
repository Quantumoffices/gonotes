package main

import (
	"flag"
	"gonotes/rpcx"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:999", "server address")
)

func main() {
	flag.Parse()
	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(rpcx.Arith), "")
	s.Serve("tcp", *addr)
}
