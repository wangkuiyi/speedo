package main

import (
	"net/rpc"
	"time"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
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
		var width, height float32
		ops := make(map[touch.Sequence]touch.Event)
		for {
			select {
			case client = <-connected:
				a.Send(paint.Event{})

			case e := <-a.Events():
				switch e := a.Filter(e).(type) {
				case lifecycle.Event:
					glctx, _ = e.DrawContext.(gl.Context)
				case size.Event:
					width, height = float32(e.WidthPx), float32(e.HeightPx)
				case touch.Event:
					if client != nil {
						switch e.Type {
						case touch.TypeBegin:
							ops[e.Sequence] = e
						case touch.TypeMove, touch.TypeEnd:
							op, arg := OpArg(ops, e, width, height)
							var dumb int
							if e := client.Call("Speedo."+op, arg, &dumb); e != nil {
								client.Close()
								client = nil
							}
						}
					}
				case paint.Event:
					if glctx == nil {
						continue
					}
					onDraw(glctx)
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

func onDraw(glctx gl.Context) {
	if client == nil {
		glctx.ClearColor(1, 0, 0, 1)
	} else {
		glctx.ClearColor(0, 1, 0, 1)
	}
	glctx.Clear(gl.COLOR_BUFFER_BIT)
}

func OpArg(ops map[touch.Sequence]touch.Event, e touch.Event, width, height float32) (op string, arg float32) {
	begin := ops[e.Sequence]

	if begin.Y > height/2 {
		op = "Accelerate"
		arg = (begin.X - e.X) / width
	} else {
		op = "Turn"
		arg = (begin.Y - e.Y) / height / 2.0
	}
	return
}

func maxMin(x, y int) (float32, float32) {
	if x >= y {
		return float32(x), float32(y)
	} else {
		return float32(y), float32(x)
	}
}
