package main

import (
	"github.com/flyteorg/flyte/cmd/single"
	"github.com/golang/glog"
)

func main() {
	glog.V(2).Info("Starting Flyte")
	err := entrypoints.Execute()
	if err != nil {
		panic(err)
	}
}
