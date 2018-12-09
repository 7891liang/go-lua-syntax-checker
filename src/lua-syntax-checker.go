package main

import (
	"fmt"
	"go-lua-syntax-checker/github.com/golua/lua"
	"os"
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



	error :=state.LoadFile(path)
	if error != nil {
		fmt.Errorf("Lua Error %s",error)
	}
}
