package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
	"xins"
	"xins/examples/chat/server/object"
	protocol "xins/protocol/default"
)

var (
	handle = protocol.NewDefaultPacker()
	codc   = &xins.JsonCodec{}
)

func main() {
	//打开连接:
	conn, err := net.Dial("tcp", "localhost:9900")
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	// go attack(conn)
	// go speak(conn)
	go read(conn)
	// go toGroupMessage(conn)

	select {}

}

// func speak(conn net.Conn) {

// 	for {

// 		message := protocol.NewMessage(1, []byte("ping"))

// 		m, _ := handle.Pack(message)

// 		conn.Write(m)

// 		time.Sleep(time.Second)
// 	}
// }

func toGroupMessage(conn net.Conn) {

	for {
		groupMessage := object.NewGroupMessage("1", "user-a", "abc")
		bytes, err := json.Marshal(&groupMessage)

		if nil != err {
			fmt.Println(err.Error())
			continue
		}

		message := protocol.NewMessage(12, bytes)

		m, _ := handle.Pack(message)

		conn.Write(m)

		time.Sleep(time.Second)
	}

}

func read(conn net.Conn) {

	for {

		message, err := handle.Unpack(conn)

		if nil != err {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("[recv error] [read head] %s\n", err.Error())
			continue
		}

		fmt.Println(string(message.Data()))
	}
}
