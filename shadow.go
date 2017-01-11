package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	{
		// go1.7.4, go1.8rc1: declaration of "ctx" shadows declaration at shadow.go:9
		ctx, cancel := context.WithCancel(ctx)
		cancel()
	}
	fmt.Println(ctx)
}
