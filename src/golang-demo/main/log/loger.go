package log

import (
	//"flag"
	"fmt"
	"log"
	"os"
	"time"
)

//var logF = flag.String("log", "test.log", "Log file name") //将运行时参数 地址 绑定到logF 运行时没带参数默认logF为test.log


func TestLog(logF *string) {
	//flag.Parse() //解析参数付给logF
	outfile, err := os.OpenFile(*logF, os.O_CREATE|os.O_RDWR|os.O_APPEND, 777) //打开文件，若果文件不存在就创建一个同名文件并打开
	if err != nil {
		fmt.Println(*outfile, "open failed")
		os.Exit(1)
	}
	log.SetOutput(outfile) //设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate | log.Ltime |log.Lmicroseconds| log.Llongfile) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名
	log.SetPrefix("[DEBUG]  ")
	//write log
	log.Printf("test out:%v \n", "test log") //向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了
}

func  TestLogger()  {
	//创建输出日志文件
	logFile, err := os.Create("./" + time.Now().Format("20060102") + ".txt");
	if err != nil {
		fmt.Println(err);
	}

	loger := log.New(logFile, "[OUTPUT]", log.Ldate|log.Ltime|log.Lshortfile);

	fmt.Println(loger.Flags());
	//SetFlags设置输出选项
	loger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile);
	//等价于print();os.Exit(1);
	//loger.Fatal("我是错误");

	//等价于print();panic();
	//	loger.Panic("我是错误");
	loger.Println("INFO 输出")
}