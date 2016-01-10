package main

import (
	"net/rpc"
	"time"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/gl"
)

const (
	server = "192.168.1.113:12345"
)

var (
	connected             = make(chan *rpc.Client)
	client    *rpc.Client = nil
)

func main() {
	go connectToServer()

	app.Main(func(a app.App) {
		var glctx gl.Context
		sz := size.Event{}
		for {
			select {
			case client = <-connected:
				a.Send(paint.Event{})

			case e := <-a.Events():
				switch e := a.Filter(e).(type) {
				case lifecycle.Event:
					glctx, _ = e.DrawContext.(gl.Context)
				case size.Event:
					sz = e
				case paint.Event:
					if glctx == nil {
						continue
					}
					onDraw(glctx, sz)
					a.Publish()
				}
			}
		}
	})
}

func connectToServer() {
	for {
		c, e := rpc.DialHTTP("tcp", server)
		if e == nil {
			connected <- c
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func onDraw(glctx gl.Context, sz size.Event) {
	if client == nil {
		glctx.ClearColor(1, 0, 0, 1)
	} else {
		glctx.ClearColor(0, 1, 0, 1)
	}
	glctx.Clear(gl.COLOR_BUFFER_BIT)
}
