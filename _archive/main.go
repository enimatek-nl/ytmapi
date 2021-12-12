package main

import (
	"flag"
	"log"
	"runtime"
	"ytmapi/api"
)

func main() {
	var port string
	var driver string
	var verbose bool
	flag.StringVar(&port, "p", "8989", "Specify the listening port.")
	flag.StringVar(&driver, "g", "geckodriver", "Specify path to geckodriver.")
	flag.BoolVar(&verbose, "v", false, "Enable verbose logging.")
	flag.Parse()

	log.Printf("OS: %s, Architecture: %s", runtime.GOOS, runtime.GOARCH)

	server := api.NewAPI(port, driver, verbose)
	log.Printf("API listening on port: %s", port)
	server.Start()
}
