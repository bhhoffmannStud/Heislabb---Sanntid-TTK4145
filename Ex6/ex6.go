package main
 
import (
    "fmt"
    "net"
    "os"
    "strconv"
)
 
/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}
 
func main() {
    /* Lets prepare a address at any address at port 10001*/   
    ServerAddr,err := net.ResolveUDPAddr("udp",":10001")
    CheckError(err)
 
    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()
 
    buf := make([]byte, 1024)
 	var count int 

    for {

        n,_,err := ServerConn.ReadFromUDP(buf)
       var message_from_client string = string(buf[0:n])
 		
	    if message_from_client == "Hei jeg lever" {
	    	fmt.Println("You gave me something other than a number")
	    } else{
	    	count, _ = strconv.Atoi(message_from_client)
    		fmt.Printf("You gave me number %T, %v \n", count, count)
	    }
	        //if  //If it doesnt get a message in some time, realize that the other is down




        if err != nil {
            fmt.Println("Error: ",err)
        }
    }
}
