package basic

import (
	"encoding/json"
	"fmt"
)


//将对象转换成json字符串
func ParseToJSON(info *UserInfo)  {
	buf,err:=json.Marshal(info)
	fmt.Println("json:",string(buf),"err",err)
}


//将json字符串转换成对象
func ConvertToObject(jsonStr string)  {
	var info UserInfo
	err:=json.Unmarshal([]byte(jsonStr),&info)
	fmt.Println("info:",info.Name,"err",err)
}