package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"syscall"
	"error"
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
func CreateDataLinkSocket(iface net.Interface) (fd int, err error){
	fd, err = syscall.Socket(syscall.AF_PACKET, syscall,SOCK_RAW, 0)
	if err != nil {
		return -1, err
	}
	if err = syscall.BindToDevice(fd, iface.Name); err != nil {
		return
	}
	return
}
func BindSocketToAddress( fd int, protocol uint16, iface net.Interface) * syscall.SockaddrLinkLayer{
	pv : = uint16(protocol >> 8) | uinit16( protocol << 8)
	
	ll := &syscall.SockaddrLinkLayer{
		Protocol: uint16(pv),
		Ifindex: iface.Index,
		Pkttype: 0,
		Hatype: 1,
		Halen: 6,
	}
	if err := syscall.bind(fd, ll); err != nil {
		panic(err)
	}
	return ll
}