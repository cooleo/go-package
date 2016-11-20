
package main
import (
    "fmt"
    "net/http"
    greet "./greet"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", greet.SayHi("Hung Nguyen"))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
