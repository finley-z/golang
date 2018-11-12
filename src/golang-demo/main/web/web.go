package web

import (
	"flag"
	"net/http"
	"log"
	_"html/template"
	_"os"
	"fmt"
	_"io/ioutil"
	"html/template"
	"io/ioutil"
)

var addr = flag.String("127.0.0.1", ":1718", "http service address") // Q=17, R=18

func Index(){
	flag.Parse()
	//http.Handle("/", http.HandlerFunc(QR))
	http.HandleFunc("/", toIndex)
	http.HandleFunc("/edit", register)
	http.HandleFunc("/save", save)
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
	buffer, err:=ioutil.ReadFile("D:/picture/20180418140901191964467WECHATLOCATION.jpg")
	if(err!=nil){
		fmt.Errorf("open file failed")
	}
	//http.Error(w,"server error",500)
	w.Write(buffer)
}

func register(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("src/golang-demo/main/template/edit.html")
		if err != nil {
			fmt.Fprintf(w, "parse template error: %s", err.Error())
			return
		}
		t.Execute(w, nil)
	} else {
		username := r.Form["username"]
		password := r.Form["password"]
		fmt.Fprintf(w, "username = %s, password = %s", username, password)
	}

}


func save(w http.ResponseWriter, r *http.Request)  {
      	r.ParseForm()
	    body := r.Form["body"]
		fmt.Fprintf(w, "body = %s", body)
}