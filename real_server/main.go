// Package real_server 后端真实服务
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type RealServer struct {
	Addr string
}

func (r *RealServer) Run(){
	log.Println("Starting httpserver at" + r.Addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/",r.HelloHandler)
	mux.HandleFunc("/base/error",r.ErrorHandler)
	mux.HandleFunc("/test_http_string/test_http_string/aaa",r.TimeoutHandler)
	server := &http.Server{
		Addr: r.Addr,
		WriteTimeout: time.Second * 3,
		Handler: mux,
	}
	go func(){
		log.Fatal(server.ListenAndServe())
	}()
}

func (r *RealServer) HelloHandler(w http.ResponseWriter, req *http.Request){
	// 例如: 127.0.0.1:8080/abc?sdsdsa=11
	// r.Addr = 127.0.0.1:8080
	// req.URL.Path = /abc
	// fmt.Println(req.Host)
	upath := fmt.Sprintf("http://%s%s\n", r.Addr,req.URL.Path)
	readIP := fmt.Sprintf("RemoteAddr=%s,X-Forwarded-For=%v,X-Real-Ip=%v\n",req.RemoteAddr,req.Header.Get("X-Forwarded-For"),req.Header.Get("X-Real-Ip"))
	header := fmt.Sprintf("headers=%v\n",req.Header)
	io.WriteString(w,upath)
	io.WriteString(w,readIP)
	io.WriteString(w,header)
}

func (r *RealServer) ErrorHandler(w http.ResponseWriter,req *http.Request){
	upath := "error handler"
	w.WriteHeader(500)
	io.WriteString(w,upath)
}

func (r *RealServer) TimeoutHandler(w http.ResponseWriter, req *http.Request){
	time.Sleep(6 * time.Second)
	upath := "timeout handler"
	w.WriteHeader(200)
	io.WriteString(w,upath)
}

func main(){
	rs1 := &RealServer{Addr: "127.0.0.1:2003"}
	rs1.Run()
	rs2 := &RealServer{Addr: "127.0.0.1:2004"}
	rs2.Run()


	//监听关闭服务
	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<- quit
}