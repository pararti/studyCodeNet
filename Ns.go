package main

import (
	"fmt"
	"os"
	"net"
)

func main(){
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Need a domain name!")
		return
	}

	domain := args[1]

	NSs, err := net.LookupMX(domain)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, NS := range NSs {
		fmt.Println(NS.Host)
	}
}

