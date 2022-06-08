package local

import (
	"fmt"
	"net"
	"sync"
)

func LocalBootstrape() {

	server,err := net.Listen("tcp",":1080")
	if err != nil{
		fmt.Printf("Listen failed:%v\n", err)
		return
	}
	for{
		client,err := server.Accept()
		if err != nil{
			fmt.Printf("Accept failed:%v\n",err)
			continue
		}
		go process(client)
	}
}

func Shutdown(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Alpha client stopped.\n")
	//shutdwon resource
}

func process(client net.Conn)  {
	if err := Socks5Auth(client); err != nil{
		fmt.Printf("Socks5 auth error:%v\n",err)
		client.Close()
		return
	}

	target, err := Socks5Connect(client)
	if err != nil{
		fmt.Printf("connect error:%v\n",err)
		return
	}

	Socks5Forward(client,target)
}
