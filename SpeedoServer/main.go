package main

import (
	"log"
	"net/http"
	"net/rpc"
)

func main() {
	rpc.Register(new(Speedo))
	rpc.HandleHTTP()
	if e := http.ListenAndServe(":12345", nil); e != nil {
		log.Fatal("listen error:", e)
	}
}

type Speedo struct{}
