package cmd

import (
	"fmt"
	"time"
	"xins/examples/chat/object"

	"github.com/spf13/cobra"
)

func LoginCommand() *cobra.Command {
	return &cobra.Command{
		Use: "login",
		Run: login,
	}
}

func login(cmd *cobra.Command, args []string) {
	conn, err := conn()
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	for {

		login := object.NewLogin("token-id")

		message, err := protocol.NewMessage(1000, login)

		if nil != err {
			fmt.Println("[new-message] error ", err.Error())
			continue
		}

		m, _ := protocol.Pack(message)

		conn.Write(m)

		time.Sleep(5 * time.Second)
	}
}
