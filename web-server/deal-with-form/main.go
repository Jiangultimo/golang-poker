package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析url传递的参数，对于POST则解析响应包的主题（request body）
	// 如果不调用ParseForm()，下面无法获取表单的数据
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Scheme)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Hing!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("U are here")
		} else {
			fmt.Println("token is not correct")
		}
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // 进行转义
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) // 输出到客户端
	}
	//r.ParseForm()
	//fmt.Println("Method:", r.Method) // 获取请求方法
	//if r.Method == "GET" {
	//	t, _ := template.ParseFiles("login.gtpl")
	//	log.Println(t.Execute(w, nil))
	//} else {
	//	fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // 进行转义
	//	fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
	//	template.HTMLEscape(w, []byte(r.Form.Get("username"))) // 输出到客户端
	//}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}
}
