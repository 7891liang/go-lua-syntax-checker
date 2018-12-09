package main

import (
	"github.com/yuin/gopher-lua"
	"os"
)

func main() {
	path :=os.Args[1]

	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile(path); err != nil {
		panic(err)
	}
}
