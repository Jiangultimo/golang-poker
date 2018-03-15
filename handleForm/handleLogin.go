package main

import (
  "fmt"
  "html/template"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/login", login)
  err := http.ListenAndServe(":9090", nil)
  if err != nil {
    log.Fatal("ListenAndServer: ", err)
  }
}

func login(res http.ResponseWriter, req *http.Request) {
  fmt.Println("request method: ", req.Method)
  req.ParseForm()
  if req.Method == "GET" {
    t, _ := template.ParseFiles("form.html")
    log.Println(t.Execute(res, nil))
  } else {
    fmt.Println("username: ", req.Form["username"])
    fmt.Println("password: ", req.Form["password"])
    http.HandleFunc("/success", success)
  }
}

func success(res http.ResponseWriter, req *http.Request) {
  if req.Method == "GET" {
    t, _ := template.ParseFiles("success.html")
    log.Println(t.Execute(res, nil))
  }
}
