package basic

import (
	"os"
	"fmt"
	"io"
	"flag"
)

/**
 1、os包
 2、runtime包
 3、flag包
 */


 //os包练习测试

 func TestOs(){
 	//测试退出 0表示正常退出 ，非零表示错误。该语句直接终止执行，不会执行延迟函数
	//os.Exit(0)

	//获取环境参数
	gopath:=os.Getenv("GOPATH")
	fmt.Println("gopath:",gopath)

	 //获取uid
	 uid:=os.Geteuid()
	 fmt.Println("uid:",uid)

	 //获取pageSize
	 pageSize:=os.Getpagesize()
	 fmt.Println("pageSize:",pageSize)

	 //获取pid
	 pid:=os.Getpid()
	 fmt.Println("pid:",pid)

	 //获取ppid
	 ppid:=os.Getppid()
	 fmt.Println("ppid:",ppid)

	 //获取Hostname
	 hostname,err:=os.Hostname()
	 fmt.Println("hostname:",hostname,err)

	 //创建文件目录
	 var Access os.FileMode=1 << (32 - 1 - 1)
	 mkdirErr:=os.Mkdir("/Users/finley/Documents/temp/test",Access)
	 fmt.Println("Mkdir Err:",mkdirErr)
	 changeMode:=os.Chmod("/Users/finley/Documents/temp/test",os.ModePerm)
	 fmt.Println("changeMode Err:",changeMode)

	 //删除文件或文件夹
	 //RemoveErr:=os.Remove("/Users/finley/Documents/temp/test")
	 //fmt.Println("Remove Err:",RemoveErr)

	 //重命名文件或文件夹
	 //RnameErr:=os.Rename("/Users/finley/Documents/temp/test","/Users/finley/Documents/temp/baby")
	 //fmt.Println("RnameErr Err:",RnameErr)

	 //获取临时文件夹
	 tempdir:=os.TempDir()
	 fmt.Println("tempdir:",tempdir)

	 //创建文件
	 //tempTxt,crerr:=os.Create("/Users/finley/Documents/temp/baby/txt")
	 //tempTxt.WriteString("hello world")
	 //fmt.Println("crerr:",crerr)

	 //获取文件信息
	 fileInfo,serr:=os.Stat("/Users/finley/Documents/temp/baby/txt")
	 fmt.Println("fileinfo,size:",fileInfo.Size(),",modtime",fileInfo.ModTime())
	 fmt.Println("crerr:",serr)

	 //读取文件内容
	 //var buf []byte=make([]byte,1024)
	 var readBuf []byte=make([]byte,1024)

	 txt,operr:=os.Open("/Users/finley/Documents/temp/baby/txt")
	 //txt.Read(buf)
	 //str:=string(buf)
	 //fmt.Println("content:",str,"err:",operr)
	 io.ReadFull(txt,readBuf)
	 fmt.Println("content1:",string(readBuf),"err:",operr)

 }

func TestFlag()  {

	var location=flag.String("dir","/","logs dir")
	flag.Parse()
	fmt.Printf("the logs location is %v:",*location)
}