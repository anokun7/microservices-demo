package main

import (
    "os"
    "fmt"
    "net/http"
    "github.com/garyburd/redigo/redis"
)

func handler(w http.ResponseWriter, r *http.Request) {
    host := os.Getenv("HOSTNAME")
    fmt.Fprintf(w, "<p>Hi there, from <b>%s</b>!", host)
    c, err := redis.Dial("tcp", "redis-db:6379")
    if err != nil {
      panic(err)
    }
    defer c.Close()
    c.Do("INCR", host)
    keys, _ := redis.Strings(c.Do("KEYS", "*"))
    fmt.Fprintf(w, "<hr/>")
    fmt.Fprintf(w, "<table style='border-collapse: collapse;'><tr><th style='border: 1px solid black;'>Container</th><th style='border: 1px solid black;'>#</th></tr>")
    for _, key := range keys {
      value, _ := redis.Int(c.Do("GET", key))
      fmt.Fprintf(w, "<tr><td style='border: 1px solid black;'>%s</td>",key)
      fmt.Fprintf(w, "<td style='border: 1px solid black;'>%d</td></tr>",value)
    }
    fmt.Fprintf(w, "</table>")
}

func main() {
    http.HandleFunc("/demo", handler)
    http.ListenAndServe(":8080", nil)
}
