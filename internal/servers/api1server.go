package servers

import (
	"fmt"
	"net"

	procs "github.com/AlexDespod/shippingapi/internal/processes"
)

type API1server struct {
	host  string
	handl func(net.Conn) //handl its a custom function for handling connection
}

func (s *API1server) Init(host string) {
	s.host = host
	s.handl = defaulHandlerAPI1
}

func (s *API1server) SetHandler(f func(net.Conn)) {
	s.handl = f
}

func (s *API1server) Run() error {
	ln, err := net.Listen("tcp", s.host)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Printf("%T %+v", err, err)
			return err
		}

		go s.handl(conn)
	}

}

func defaulHandlerAPI1(conn net.Conn) {
	defer conn.Close()

	Proc := new(procs.API1Processing)

	Proc.SetConn(&conn)

	messageRecv, err := Proc.GetMessage()

	if err != nil {
		fmt.Printf("%T %+v \n", err, err)
		return
	}

	fmt.Printf("%v \n", *messageRecv)

	Proc.DoSomeProcces()

	err = Proc.SendJson()

	if err != nil {
		fmt.Printf("%T %+v \n", err, err)
		return
	}

	fmt.Printf("message has sent\n")
}
