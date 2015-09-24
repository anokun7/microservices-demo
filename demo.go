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
    fmt.Fprintf(w, "<div>")
    for _, key := range keys {
      value, _ := redis.Int(c.Do("GET", key))
      fmt.Fprintf(w, "<span style=\"width: 8em; padding: .2em; border: 1px dotted\">%s</span>",key)
      fmt.Fprintf(w, "<span style=\"width: 2em; padding: .2em; border: 1px dotted\">%d</span>",value)
    }
    fmt.Fprintf(w, "</div>")
}

func main() {
    http.HandleFunc("/demo", handler)
    http.ListenAndServe(":8080", nil)
}
