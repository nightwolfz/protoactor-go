package opentracing

import (
	"github.com/nightwolfz/protoactor-go/actor"
	"github.com/nightwolfz/protoactor-go/actor/middleware/propagator"
)

func TracingMiddleware() actor.SpawnMiddleware {
	return propagator.New().
		WithItselfForwarded().
		WithSpawnMiddleware(SpawnMiddleware()).
		WithSenderMiddleware(SenderMiddleware()).
		WithReceiverMiddleware(ReceiverMiddleware()).
		SpawnMiddleware
}
