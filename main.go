package main

import (
	"flag"
	"log"
	"os"
)

var (
	listen  = flag.String("listen", ":8002", "port to redirect")
	forward = flag.String("forward", "localhost:8001", "target address")
)

func main() {
	flag.Parse()

	log.Printf("Starting proxy from %s to %s with %v PID\n", *listen, *forward, os.Getpid())
	proxy := NewProxy(*listen, *forward)
	proxy.Run()
}
