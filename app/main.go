package main

import (
    "fmt"
    "net/http"
)

func response(w http.ResponseWriter, req *http.Request) {
    var sum = 0
    for i := 0; i < 1000000000; i++ {
        sum += i
    }
    fmt.Fprintf(w, "%-d\n", sum)
}

func main() {
    http.HandleFunc("/", response)

    http.ListenAndServe("0.0.0.0:8080", nil)
}
