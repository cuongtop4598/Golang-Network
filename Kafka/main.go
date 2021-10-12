package main

import (
	"context"
)

//https://www.sohamkamani.com/golang/working-with-kafka/
func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	go produce(ctx)
	consume(ctx)
}
