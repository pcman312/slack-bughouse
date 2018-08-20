package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/JTurpin/slack-bughouse/handlers"
)

func main() {
	fmt.Printf("Starting slack-bughouse...\n")

	fmt.Printf("===== ENVIRONMENT: =========\n")
	for _, e := range os.Environ() {
		fmt.Printf("%s\n", e)
	}
	fmt.Printf("===== DONE ENVIRONMENT =====\n")

	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Printf("IP: %s\n", ip.String())
		}
	}

	http.HandleFunc("/", handlers.HandleRoot)
	http.HandleFunc("/teams", handlers.HandleTeams)
	fmt.Printf("Listening...\n")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
