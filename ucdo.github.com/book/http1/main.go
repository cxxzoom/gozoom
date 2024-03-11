package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type databases map[string]dollars

func (db databases) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprintf(w, "no such page: %s\n", req.URL)
		msg := fmt.Sprintf("no such page: %s \n", req.URL)
		http.Error(w, msg, http.StatusNotFound)
	}

}

func main() {
	db := databases{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:7777", db))
}
