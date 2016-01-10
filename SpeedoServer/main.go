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

func (speedo *Speedo) Accelerate(arg float32, _ *int) error {
	log.Printf("Speedo.Accelerate %+v", arg)
	return nil
}

func (speedo *Speedo) Turn(arg float32, _ *int) error {
	log.Printf("Speedo.Turn %+v", arg)
	return nil
}
