// reverse_proxy_base 是一个简单的反向代理demo
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
)
var(
	proxy_addr = "http://127.0.0.1:2003"
	port = "2002"
)
func handler(w http.ResponseWriter, r *http.Request){
	// 1.解析代理地址，并更改请求体的协议和主机
	proxy, err := url.Parse(proxy_addr)
	r.URL.Scheme = proxy.Scheme // Scheme代表请求的协议 例如: http,https等
	r.URL.Host = proxy.Host
	fmt.Println(r.Host)
	// 2.请求下游
	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(r)
	if err != nil{
		log.Print(err,"---")
		return
	}

	// 3.把下游请求内容返回给上游
	for k, val := range resp.Header{
		for _,v := range val{
			w.Header().Add(k,v)
		}
	}

	defer resp.Body.Close()
	bufio.NewReader(resp.Body).WriteTo(w) //利用缓存io将下游请求body写进响应体中
}

func main(){
	http.HandleFunc("/",handler)
	log.Println("Start serving on port: " + port)
	err := http.ListenAndServe(":"+port,nil)
	if err != nil{
		log.Fatal(err)
	}
}