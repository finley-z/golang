package socket

import (
	"net"
	"fmt"
)

var stat bool=true
var raddr *net.TCPAddr


func GetConnet() (*net.TCPConn, error) {
	return net.DialTCP("TCP",nil,raddr)
}

func init()  {
	var err error
	raddr,err =net.ResolveTCPAddr("TCP","127.0.0.1:9088")
	if err!=nil{
		fmt.Errorf("parse error")
	}
}