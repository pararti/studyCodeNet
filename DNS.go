package main

import (
	"fmt"
	"os"
	"net"
)

func lookIP(addr string) ([]string, error) {
	hosts, err := net.LookupAddr(addr)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func lookHost(hostname string) ([]string, error) {
	IPs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}
	return IPs, nil
}

func main(){
	args := os.Args

	if len(args) == 1 {
		fmt.Println("Pls provide me an arguments!")
		return
	}

	input := args[1]
	IPaddr := net.ParseIP(input)

	if IPaddr == nil {
		IPs, err := lookHost(input)
		if err == nil {
			for _, singleIP := range IPs {
				fmt.Println(singleIP)

			}
		}
	} else {
		hosts, err := lookIP(input)
		if err == nil {
			for _, hostname := range hosts {
				fmt.Println(hostname)
			}
		}
	}
}

