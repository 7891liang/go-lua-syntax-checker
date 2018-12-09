package main

import (
	"os"
	"github.com/azure/golua/lua"
	"fmt"
)

var (
	trace bool = false
	debug bool = true
	tests bool = false
)

func main() {
	path :=os.Args[1]

	var opts = []lua.Option{lua.WithTrace(trace), lua.WithVerbose(debug)}
	state := lua.NewState(opts...)
	defer state.Close()

	str := ""
	if err := state.LoadFile(path); err != nil {
		str +=fmt.Sprintf("[%s] Error [%s] \n",path,err.Error())
	}
	fmt.Print(str)

	if str != "" {
		fmt.Print("has error")
	}
}
