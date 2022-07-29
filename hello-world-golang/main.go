package main

import (
    "fmt"
    "net/http"
	"rsc.io/quote"
)

func main() {
	quote.Hello()
	fmt.Println(quote.Opt())
	fmt.Println(quote.Glass())
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
	fmt.Println("Server listening on port 8080")
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("J'ai reçu une requête ouais ouais ヾ(•ω•`)o")
    fmt.Fprintf(w, "Hello, World! ヾ(•ω•`)o")
}