package main

import (
	console "github.com/AsynkronIT/goconsole"
	"github.com/nightwolfz/protoactor-go/actor"
	"github.com/nightwolfz/protoactor-go/examples/remotebenchmark/messages"
	"github.com/nightwolfz/protoactor-go/remote"
	"fmt"
)

func main() {
	remote.Start("127.0.0.1:8080")
	rootContext := actor.EmptyRootContext
	props := actor.
		PropsFromFunc(
			func(context actor.Context) {
				switch context.Message().(type) {
				case *messages.Ping:
					fmt.Println("Received ping from sender with address: " + context.Sender().Address)
					context.Respond(&messages.Pong{})
				}
			})

	rootContext.SpawnNamed(props, "remote")

	console.ReadLine()
}
