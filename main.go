package main
import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)
var (
	online = make(chan bool,1)
	bind = false
	local = flag.String("l", "0.0.0.0:222", "0.0.0.0:222 [本地网卡监听端口]")
	dst = flag.String("r", "8.8.8.8:1234", "8.8.8.8:1234 [NC、CobaltStrike、Metasploit等服务监听端口]")
	A net.Conn
	B net.Conn
)
func init() {
	flag.BoolVar(&bind, "bind", true, "local to local [正向监听:127.0.0.1:10=>192.168.1.100:10,边界:192.168.1.100:10=>(边界)=>8.8.8.8:1234]")
}
func main() {
	flag.Parse()
	if len(os.Args)<=2 {
		flag.Usage()
	}else {
		if bind {
			go bind_listen_A()
			bind_listen_B()
		}else {
			listen()
		}
	}

}


func listen()  {
	log.Println("Start listen",*local)
	for {
		ln,err := net.Listen("tcp",*local)
		if err != nil {
			fmt.Println("tcp_listen:",err)
			return
		}
		defer ln.Close()
		for{
			tcp_Conn,err:=ln.Accept()
			if err!=nil{
				fmt.Println("Accept:",err)
				return
			}
			go tcp_handle(tcp_Conn)
		}
	}
}
func tcp_handle(tcpConn net.Conn){
	remote_tcp,err:=net.Dial("tcp",*dst)
	if err!=nil{
		fmt.Println(err)
		return
	}
	log.Println(dst,"=>",*local)
	go io.Copy(remote_tcp,tcpConn)
	log.Println(local,"=>",*dst)
	go io.Copy(tcpConn,remote_tcp)

}

func bind_listen_A()  {
	log.Println("Start listen",*local)
	for {
		ln,err := net.Listen("tcp",*local)
		if err != nil {
			fmt.Println("tcp_listen:",err)
			return
		}
		defer ln.Close()
		for{
			A,err=ln.Accept()
			if err!=nil{
				fmt.Println("Accept:",err)
				return
			}
			go func() {
				if <-online {
					go bind_tcp_handle(A,B)
				}
			}()
		}
	}
}

func bind_listen_B()  {
	log.Println("Start listen",*dst)
	for {
		ln,err := net.Listen("tcp",*dst)
		if err != nil {
			fmt.Println("tcp_listen:",dst)
			return
		}
		defer ln.Close()
		for{
			B,err=ln.Accept()
			if err!=nil{
				fmt.Println("Accept:",err)
				return
			}
			go func() {
				online <- true
				go bind_tcp_handle(B,A)
			}()

		}
	}
}
func bind_tcp_handle(local net.Conn,dst net.Conn){
	log.Println(dst.LocalAddr(),"=>",local.LocalAddr())
	io.Copy(local,dst)
	log.Println(local.LocalAddr(),"=>",dst.LocalAddr())
	io.Copy(dst,local)
}
