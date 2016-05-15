package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: "/dev/tty.usbmodem1421", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Write([]byte{byte(50)})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wrote a, Waiting for a while")
	time.Sleep(10 * time.Second)

	_, err = s.Write([]byte{byte(120)})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wrote Z. Waiting for a while")
	time.Sleep(3 * time.Second)
}
