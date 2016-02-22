package main
 
import (
    "fmt"
    "net"
    "time"
    "strconv"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}


/* 
func processpair(bool i, lastcount){
    if(i==0){
        backupmode
        for{
            if(clientwentdown){
                spawn processpair(0)
                break //becomes client
            }
        }
    } 
    clientmode
}
*/




func main() {
   // spawn processpair(1)
   // spawn processpair(0)




    ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
    CheckError(err)
 
    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    CheckError(err)
 
    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)
 
    defer Conn.Close()
    i := 0
    for {
        
        for xx := 0; xx < 3; xx++{
            msg := "Hei jeg lever"
            buf := []byte(msg)
            _,err := Conn.Write(buf)

             if err != nil {
            fmt.Println(msg, err)
             }
            time.Sleep(time.Second * 1)

        }

        msg := strconv.Itoa(i)
        i++
        time.Sleep(time.Second*2)
        buf := []byte(msg)
        _,err := Conn.Write(buf)
        fmt.Println(msg)
        if err != nil {
            fmt.Println(msg, err)
        }
        time.Sleep(time.Second * 1)
    }
}