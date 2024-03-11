package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

type databases map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db databases) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db databases) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db databases) add(w http.ResponseWriter, req *http.Request) {
	items := req.URL.Query()
	if items == nil {
		return
	}
	for key, value := range items {
		// key为参数名，values为参数值的切片
		v, err := strconv.ParseFloat(value[0], 32)
		if err != nil {
			fmt.Fprintf(w, "%s\n", err)
			return
		}
		db[key] = dollars(v)
	}
}

func (db databases) del(w http.ResponseWriter, req *http.Request) {
	items := req.URL.Query().Get("item")
	delete(db, items)
}

func main() {
	db := databases{"shoes": 50, "socks": 5}
	// mux := http.NewServeMux()
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/del", db.del)
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
