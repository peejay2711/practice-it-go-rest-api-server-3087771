package backend

import (
	"fmt"
	"log"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world/n")

}

func Run(addr string) {
	http.HandleFunc("/", helloworld)

	fmt.Println("Server is on and its running", addr)

	log.Fatal(http.ListenAndServe(addr, nil))

}
