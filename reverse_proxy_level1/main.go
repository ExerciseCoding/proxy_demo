package main

import (
	"net"
	"net/http"
	"time"
)

// 第一层代理地址
var addr = "127.0.0.1:2001"

func main(){

}

var transport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout: 30 * time.Second, // 连接超时
		KeepAlive: 30 * time.Second, // 长连接超时时间
	}).DialContext,
	MaxIdleConns: 100, // 最大空闲连接
	IdleConnTimeout: 90 * time.Second, // 空闲超时时间
	TLSHandshakeTimeout: 10 * time.Second, // tls握手超时时间
	ExpectContinueTimeout: 1 * time.Second, // 100-continue 超时时间
}


