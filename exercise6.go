package main
 
import (
    "fmt"
    "net"
    "time"
    "strconv"
    "os/exec"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}



func Processpair(){
     channel1 := make(chan string, 1)
     var count = 0
     var status = 0
     go func(){
                   ServerAddr,err := net.ResolveUDPAddr("udp",":10001")
                    CheckError(err)
                    ServerConn, err := net.ListenUDP("udp", ServerAddr)
                    CheckError(err)
                    defer ServerConn.Close()
                    buf := make([]byte, 1024)
                    n,_,_ := ServerConn.ReadFromUDP(buf)
                    var message_from_client string = string(buf[0:n])
                    channel1 <- message_from_client
                }()


    select {
                    case <-channel1:
                           //Backup begin
                            ServerAddr,err := net.ResolveUDPAddr("udp",":10001")
                            CheckError(err)
                     
                             ServerConn, err := net.ListenUDP("udp", ServerAddr)
                             CheckError(err)
                             defer ServerConn.Close()
                     
                              buf := make([]byte, 1024)
                              var count int 

                             for {
                                    c1 := make(chan string, 1)
                                    go func(){
                                        n,_,_ := ServerConn.ReadFromUDP(buf)
                                        var message_from_client string = string(buf[0:n])
                                        c1 <- message_from_client
                                    }()
                                    select{
                                        case message := <-c1:
                                            count, _ = strconv.Atoi(message)
                                        case <-time.After(time.Millisecond*1200):
                                            count++
                                            fmt.Println("Hei")
                                            ServerConn.Close()
                                            cmd :=exec.Command("gnome-terminal", "-e", "./exercise6")
                                            cmd.Start()
                                            fmt.Println("Main timed out")
                                            status = 1
                                            break
                                    }
                                    if(status == 1){
                                        break
                                    }
                                 }
                                           // Main begin

                                         
                                            LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
                                            CheckError(err)
                                         
                                            Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
                                            CheckError(err)

                                         
                                            defer Conn.Close()

                                            i := count
                                            killtimer := 0

                                            for {
                                                

                                                msg := strconv.Itoa(i)
                                                i++
                                                buf := []byte(msg)
                                                _,err := Conn.Write(buf)
                                                fmt.Println(msg)
                                                if err != nil {
                                                    fmt.Println(msg, err)
                                                }
                                                
                                                  if(killtimer == 9){
                                                    break;
                                                  }else{
                                                    killtimer++
                                                  }
                                                time.Sleep(time.Second * 1)

                                            }

                                            //Main end
                          

                            
                        case <- time.After(time.Millisecond*2000):
                                            // Main begin
                                            ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
                                            CheckError(err)
                                         
                                            LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
                                            CheckError(err)
                                         
                                            Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
                                            CheckError(err)

                                         
                                            defer Conn.Close()

                                            i := count
                                            killtimer := 0
                                            for {
                                                

                                                msg := strconv.Itoa(i)
                                                i++
                                                buf := []byte(msg)
                                                _,err := Conn.Write(buf)
                                                //Klarer du å drepe her får vi feil
                                                fmt.Println(msg)
                                                if err != nil {
                                                    fmt.Println(msg, err)
                                                }
                                                
                                                  if(killtimer == 10){
                                                    break;
                                                  }else{
                                                    killtimer++
                                                  }
                                                time.Sleep(time.Second * 1)

                                            }

                                            //Main end
                       
            }
}






func main() {
   Processpair()
}
