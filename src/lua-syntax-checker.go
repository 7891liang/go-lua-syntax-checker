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

var (
	trace bool = false
	debug bool = true
	tests bool = false
)

func main() {
	path :=os.Args[1]
	currPath :=""
	if path == "." {
		currPath = GetCurrentDirectory()
	}else{
		currPath = path
	}


	var opts = []lua.Option{lua.WithTrace(trace), lua.WithVerbose(debug)}
	state := lua.NewState(opts...)
	defer state.Close()
	str := ""
	xfiles, _ := GetAllFiles(currPath)
	for _, file := range xfiles {
		if err := state.LoadFile(file); err != nil {
			str +=fmt.Sprintf("[%s] Error -> %s",file,err.Error())
		}
	}

	if str != "" {
		fmt.Print(str)
	}
}


func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
		}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}



//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".lua")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}