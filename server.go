package main

import (
	"flag"
	"fmt"
	"log"
	"osiris/internal/handlers"
	"net/http"
	"os"
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: helloserver [options]\n")
    flag.PrintDefaults()
    os.Exit(2)
}

var (
    addr = flag.String("addr", "localhost:8080", "address to serve")
)

func main() {
    // Parse flags.
    flag.Usage = usage
    flag.Parse()

    // Register handlers
    http.HandleFunc("/", handlers.GreetHandler)
    http.HandleFunc("/version", handlers.VersionHandler)

    log.Printf("serving http://%s\n", *addr)
    log.Fatal(http.ListenAndServe(*addr, nil))
}
