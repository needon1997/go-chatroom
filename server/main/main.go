package main

import (
	"fmt"
	"net"

	"../controller"
	"../onlineMan"
	"../util"
)

var dispatcher = controller.NewDispatcher()

func main() {

	li, err := net.Listen("tcp", "0.0.0.0:8888")
	defer li.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	var transfer *util.Transfer = new(util.Transfer)
	transfer.Conn = conn
	for {
		msg, err := transfer.ReadMsg()
		if err != nil {
			fmt.Println(err)
			onlineMan.Manager.Offline(conn)
			return
		}
		fmt.Println(msg)
		result, err := dispatcher.HandleMsg(msg, conn)
		if err != nil {
			fmt.Println(err)
			// sendErr(conn, err)
			onlineMan.Manager.Offline(conn)
			return
		}
		fmt.Println(result)
		if result != nil {
			err = transfer.WriteResult(result)
			if err != nil {
				fmt.Println(err)
				// sendErr(conn, err)
				onlineMan.Manager.Offline(conn)
				return
			}
		}
	}
}

//need a decoder and encoder
//filter
//need a context tell the handler it should stop when is connection is dead
// dont waste time in doing IO
//need to organize the code
