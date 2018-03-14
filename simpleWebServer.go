package main
import (
  "fmt"
  "net/http"
  "strings"
  "log"
)

func main() {
  http.HandleFunc("/", sayhelloName)
  err := http.ListenAndServe(":9090", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
  r.ParseForm() //解析参数，默认是不会解析
  fmt.Println(r.Form)
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])
  for k, v := range r.Form {
    fmt.Println("key: ", k)
    fmt.Println("val: ", strings.Join(v, ""))
  }
  fmt.Fprintf(w, "hello golang")
}
