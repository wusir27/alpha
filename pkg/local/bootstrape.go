package local

import (
	"fmt"
	"net"
)

func LocalBootstrape() {

	server,err := net.Listen("tcp",":1090")
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

func process(client net.Conn)  {
	if err := Socks5Auth(client); err != nil{
		fmt.Printf("Socks5 auth error:%v",err)
		client.Close()
		return
	}

	target, err := Socks5Connect(client)
	if err != nil{
		fmt.Printf("connect error:%v",err)
		return
	}

	Socks5Forward(client,target)
}
