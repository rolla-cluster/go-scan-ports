// This is an example project to look into building a port scanner in Go
package main

import (
	"fmt"
	"net"
	"sort"
  "os"
)

var url string // global for url to scan

func init() {
  // General info
  VERSION := "1.0.1"
  AUTHOR := "Rolla-cluster"
  fmt.Println("Initializing port scanner version: ", VERSION)
  fmt.Printf("--- %s ---\n", AUTHOR)  
  // Check for 2 args (prog inclusive)
  if len(os.Args) != 2 {
    panic("Missing argument: url")
  }

}

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", url, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
  
  url = os.Args[1]
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
