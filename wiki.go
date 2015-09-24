package main

import (
    "os"
    "fmt"
    "net/http"
    "github.com/garyburd/redigo/redis"
)

func handler(w http.ResponseWriter, r *http.Request) {
    host := os.Getenv("HOSTNAME")
    fmt.Fprintf(w, "Hi there, from %s!", host)
    c, err := redis.Dial("tcp", "db:6379")
    if err != nil {
      panic(err)
    }
    defer c.Close()
    c.Do("SET", host, 1)
    stats, err := redis.String(c.Do("GET", host))
    if err != nil {
      fmt.Fprintf(w, "Key not found")
    }
    fmt.Fprintf(w, "%s: %s", host, stats)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
