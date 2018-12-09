package main

import (
	"os"
	"github.com/azure/golua/lua"
)

var (
	trace bool = false
	debug bool = false
	tests bool = false
)

func main() {
	path :=os.Args[1]

	var opts = []lua.Option{lua.WithTrace(trace), lua.WithVerbose(debug)}
	state := lua.NewState(opts...)
	defer state.Close()

	if err := state.LoadFile(path); err != nil {
		panic(err)
	}
}
