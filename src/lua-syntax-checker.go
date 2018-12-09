package main

import (
	"os"
	"fmt"
	"strings"
	"path/filepath"
	"log"
	"io/ioutil"
	"github.com/azure/golua/lua"
)

func main() {
	path :=os.Args[1]
	output :=""
	if len(os.Args) == 3 {
		output =os.Args[2]
	}

	currPath :=""
	if path == "." {
		currPath = GetCurrentDirectory()
	}else{
		currPath = path
	}


	var opts = []lua.Option{lua.WithTrace(false), lua.WithVerbose(true)}
	state := lua.NewState(opts...)
	defer state.Close()
	str := ""
	xfiles, _ := GetAllFiles(currPath)
	for _, file := range xfiles {
		if err := state.LoadFile(file); err != nil {
			str +=fmt.Sprintf("[%s] Error -> %s",strings.Replace(file,currPath,"",-1),err.Error())
		}
	}

	if str != "" {
		if output == "" {
			fmt.Print(str)
		}else
		{
			WriteWithIoutil(output,str)
		}
	}
}


func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		}
	return strings.Replace(dir, "\\", "/", -1)
}


func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() {
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {

			ok := strings.HasSuffix(fi.Name(), ".lua")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

func WriteWithIoutil(name,content string) {
	data :=  []byte(content)
	if ioutil.WriteFile(name,data,0644) == nil {
		fmt.Println("写入文件成功:",content)
	}
}
