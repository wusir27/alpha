package local

import (
	"errors"
	"io"
	"net"
)

func Socks5Auth(client net.Conn) (err error) {
	buf := make([]byte, 256)

	//read ver and nmethods
	n,err := io.ReadFull(client, buf[:2])
	if n != 2 {
		return errors.New("Reading header:" + err.Error())
	}

	ver,nMethods := int(buf[0]), int(buf[1])
	if ver != 5 {
		return errors.New("invalid socks version")
	}

	n,err = io.ReadFull(client, buf[:nMethods])
	if n != nMethods {
		return errors.New("reading methods:" + err.Error())
	}

	n,err = client.Write([]byte{0x05,0x00})
	if n != 2 || err != nil{
		return errors.New("write rsp err:" + err.Error())
	}

	return nil
}

func Socks5Connect(client net.Conn)(net.Conn, error)  {
	buf := make([]byte, 256)

	n,err := io.ReadFull(client, buf[:4])
	if n != 4 {
		return nil, errors.New("read header: "+ err.Error())
	}

	ver,cmd, _, atyp := buf[0], buf[1], buf[2], buf[3]
	if ver != 5 || cmd != 1{
		return nil, errors.New("invalid ver/cmd")
	}


	return nil,nil
}

func Socks5Forward(client net.Conn, target net.Conn)(err error)  {

	return nil
}