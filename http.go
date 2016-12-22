package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "God Jul!")
}

func to_roman(n int) string {
    if n == 2 {
        return "II"
    }
    return "I" // oopsie
}

type romanGenerator int

func (n romanGenerator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    number := r.URL.Query().Get("number")
    if len(number) == 0 {
        fmt.Fprintf(w, "Skriv inn ditt tall i URL'en")
    }
    i, err := strconv.Atoi(number)

    if err == nil {
        fmt.Fprintf(w, "Her er ditt lykketall: %s\n", to_roman(i))
    }
}

func main() {
    h := http.NewServeMux()

    h.Handle("/roman/", romanGenerator(1))
    h.HandleFunc("/", hello)

    err := http.ListenAndServe(":8000", h)
    panic(err)
}
