package web

import (
	"flag"
	"net/http"
	"log"
	_"html/template"
	_"os"
	"fmt"
	_"io/ioutil"
)

var addr = flag.String("127.0.0.1", ":1718", "http service address") // Q=17, R=18

func Index(){
	flag.Parse()
	//http.Handle("/", http.HandlerFunc(QR))
	http.HandleFunc("/", toIndex)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func toIndex(w http.ResponseWriter, r *http.Request) {
	//t, _ := template.ParseFiles("template/view.html")
	//t.Execute(w, nil)
	//f,err :=os.Open("D")
	size :=r.FormValue("size")
	fmt.Println("size=",size)
	//buffer, err:=ioutil.ReadFile("D:/picture/20180418140901191964467WECHATLOCATION.jpg")
	//if(err!=nil){
	//	fmt.Errorf("open file failed")
	//}
	http.Error(w,"server error",500)
	//w.Write(buffer)
}

