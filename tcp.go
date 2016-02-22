// 129.241.187.23
package main

import (
    "fmt"
    "net"
    "runtime"
   // "bufio"
)

const server = "129.241.187.23"
const tcpPortFixed = "34933"
const tcpPortDelim = "33546"

func connectTCPFixed() {
	conn, err := net.Dial("tcp", net.JoinHostPort(server, tcpPortFixed))
	if err != nil {
	fmt.Println("Klarte ikke å koble til TCP")
	}
	defer conn.Close()
	tmp := make([]byte, 1024)
	i, _ := conn.Read(tmp) 
	fmt.Print("Message from server: "+string(tmp[:i])) 
}

func connectTCPDelim() {
	conn, err := net.Dial("tcp", net.JoinHostPort(server, tcpPortDelim))
	if err != nil {
	fmt.Println("Klarte ikke å koble til TCP")
	}
	defer conn.Close()
	tmp := make([]byte, 1024)
	i, _ := conn.Read(tmp) 
	fmt.Print("Message from server: "+string(tmp[:i])) 
}


func connectTCPSendAndListen() {
	conn, err := net.Dial("tcp", net.JoinHostPort(server, tcpPortDelim))
	if err != nil {
	fmt.Println("Klarte ikke å koble til TCP")
	}
	defer conn.Close()

	text := "Heihei\x00"
	fmt.Fprintf(conn, text)
	tmp := make([]byte, 1024)
	i, _ := conn.Read(tmp) 
	fmt.Print("Message from server: "+string(tmp[:i])) 
}


func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    fmt.Println("=== Fixed TCP communication ===");
    connectTCPFixed()
    fmt.Println("\n")
    fmt.Println("=== Delim TCP communication ===");
    connectTCPDelim()
    fmt.Println("\n")
    fmt.Println("=== TCP communication Send and Listen===");
    connectTCPSendAndListen()
}
