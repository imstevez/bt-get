package main

import (
	"bt-get/clean"
	"context"
	"fmt"
	"os"
	"time"
)

func init() {
	time.Local = time.UTC
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	clean.PushCancel(cancel)
	err := cmds.RunContext(ctx, os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
