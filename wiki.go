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
    c, err := redis.Dial("tcp", "db:6379")
    if err != nil {
      panic(err)
    }
    defer c.Close()
    n, err := c.Do("EXISTS", host)
    if n == 0 {
      c.Do("SET", host, 1)
    } else {
      incr, _ := redis.Int(c.Do("GET", host))
      c.Do("SET", host, incr+1)
    }
    stats, err := redis.String(c.Do("GET", host))
    if err != nil {
      fmt.Fprintf(w, "Key not found")
    }
    fmt.Fprintf(w, "<hr/>")
    fmt.Fprintf(w, "<table>")
    fmt.Fprintf(w, "<tr><td>%s</td><td>%s</td></tr>", host, stats)
    fmt.Fprintf(w, "</table>")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
