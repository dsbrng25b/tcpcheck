package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	var timeout int
	var address string

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of: %s [options] host:port\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.IntVar(&timeout, "timeout", 5, "connection timeout in seconds")

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	address = flag.Arg(0)

	conn, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Second)
	if err != nil {
		fmt.Println("connection failed", err)
	} else {
		fmt.Println("connection ok")
		conn.Close()
	}
}
