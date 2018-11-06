package socket

import (
	"net"
	"fmt"
	_"os"
	"time"
)

var servce *net.TCPAddr


func Listen()  {
	listener, err := net.ListenTCP("tcp", servce)
	if err!=nil{
		fmt.Errorf("parse error")
	}
	for {
		conn, err := listener.Accept()
		fmt.Println("server accpet")

		if err != nil {
			fmt.Println("server error")
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // don't care about return value
		conn.Close()                // we're finished with this client
	}
}

func init()  {
	var err error
	raddr,err =net.ResolveTCPAddr("ip4","127.0.0.1:9088")
	if err!=nil{
		fmt.Errorf("parse error")
	}
}

