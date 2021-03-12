// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2021/3/12 - 10:02 下午 - UTC/GMT+08:00

package main

import (
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/csrf",csrfHandler)
	log.Println("CSRF Server Running....")
	http.ListenAndServe(":2021", nil)
}

func csrfHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "csrf.html")
}