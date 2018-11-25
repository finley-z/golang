package io

import "fmt"
import (
	"io/ioutil"
	"os"
)

/***
  文件IO
  1、文件创建、删除
  2、文件属性获取
  3、文件写入读取
 */

func ReadFiles(){
	dir_list, e := ioutil.ReadDir("/Users/finley/Documents/")
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	for i, v := range dir_list {
		fmt.Println(i, "=", v.Name())
		fmt.Println(v.Name(), "的权限是:", v.Mode())
		fmt.Println(v.Name(), "文件大小:", v.Size())
		fmt.Println(v.Name(), "创建时间", v.ModTime())
		fmt.Println(v.Name(), "系统信息", v.Sys())
		if v.IsDir() == true {
			fmt.Println(v.Name(), "是目录")
		}else{
			ioutil.ReadFile(v.Name())
			fmt.Println(v.Name(), "是文件")
		}
	}
}


func ReadFileInfo(){
	content, err := ioutil.ReadFile("/Users/finley/Documents/temp/baby/txt")
	if err != nil {
		fmt.Println("read error")
		os.Exit(1)
	}
	fmt.Println(string(content))
}

func WriteFile(){
	file,err:=ioutil.TempFile("H:/", "go.log")
	defer file.Close()
	if err != nil {
		fmt.Println("创建文件失败")
		return
	}
	file.WriteString("Hello word") //利用file指针的WriteString()详情见os.WriteString()
	filedata, _ := ioutil.ReadFile(file.Name())

	fmt.Println(string(filedata))
}