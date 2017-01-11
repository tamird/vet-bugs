package main

import (
	"context"
	"strings"
)

func main() {
	ctx := context.Background()
	{
		// go1.7.4, go1.8rc1: declaration of "ctx" shadows declaration at shadow.go:9
		ctx, cancel := context.WithCancel(ctx)
		_ = ctx
		cancel()
	}

	var strs []string
	for _, str := range strs {
		_ = str

		// no warning
		str := "foo"
		_ = str
	}
	for _, str := range strs {
		// go1.8rc1: declaration of "str" shadows declaration at shadow.go:25
		str := strings.Replace(str, "1", "2", -1)
		_ = str
	}
}
