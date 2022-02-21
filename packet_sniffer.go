package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil{
		fmt.Println(err)
		os.Exit(1) // Exiting the program with a status code of 1
	}
	fmt.Printf("--------- Network Interfaces for user: %c ---------- \n", hostname)
	PrintInterfaces()
	os.Exit(1)
}


func PrintInterfaces(){
 
    interfaces, err := net.Interfaces()
 
    if err == nil {
        for _, i := range interfaces {
            if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				fmt.Println(i.Name)
				fmt.Println("-------------------------------------------------")
           }
        }
     }
 
}
