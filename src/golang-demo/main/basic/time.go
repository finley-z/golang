package basic

import (
	"time"
	"fmt"
)



//时间测试
func  TestTime()  {

	now:=time.Now()
	tim:=now.Format("2006-01-02 15:04:05")
	fmt.Println("标准当前时间:",tim)

	//将字符串装换成日期
	t,err:=time.Parse("2006-01-02 15:04:05","2018-11-19 15:02:33")
	fmt.Println("日期转换成时间:",t,"err:",err)

	//时间间隔
	duration:=time.Since(t)
	fmt.Println("duration:",duration,"hours:",duration.Hours(),"minute:",duration.Minutes())

	fmt.Println("now time:",time.Now())

	ti:=now.Add(10*time.Second)
	fmt.Println("after 10 s:",ti)


}

//日期测试
func TestDate()  {

}