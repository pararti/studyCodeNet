package main

import (
	"fmt"
	"net"
)

func main(){
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	for _, i := range interfaces {
		fmt.Printf("Interface: %v\n", i)

		byName, err := net.InterfaceByName(i.Name)
		if err != nil{
			fmt.Println(err)
		}

		addr, err := byName.Addrs()
		if err != nil{
			fmt.Println(err)
			continue
		}
		for k, v := range addr {
			fmt.Printf("Interface addres: #%v, %s\n", k, v.String())
		}
		fmt.Println()
	}
}
