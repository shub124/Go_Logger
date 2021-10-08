package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	logserver "services/Libraries/Server"
	router "services/Router"
	"strings"
)

func main() {

	address := flag.String("address", ":80", "port on which it listens")
	brokers := flag.String("brokers", "", "comma separated broker list")
	flag.Parse()
	if *brokers == "" {
		os.Exit(1)
	}
	brokerlist := strings.Split(*brokers, ",")
	fmt.Println(brokerlist)
	logserver.Init(brokerlist)
	defer func() {
		logserver.Close()
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/log", handler)
	log.Fatal(logserver.Run(*address, mux))

}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if req.Method == "OPTIONS" {
		return
	}
	route, ok := req.URL.Query()["route"]
	if !ok || len(route) == 0 {
		fmt.Println("Bad request")
		return
	}
	controllerfunc, ok := router.Routemapping[route[0]]
	if ok {
		controllerfunc(w, req)
	} else {
		http.Error(w, "Requested URI doesn't exist", 404)
	}

}
