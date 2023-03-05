package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func PingCommand() *cobra.Command {
	return &cobra.Command{
		Use: "ping",
		Run: ping,
	}
}

func ping(cmd *cobra.Command, args []string) {
	conn, err := conn()
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	for {

		message, err := protocol.NewMessage(1, []byte("ping"))

		if nil != err {
			fmt.Println("[new-message] error ", err.Error())
			continue
		}

		m, _ := protocol.Pack(message)

		conn.Write(m)

		time.Sleep(time.Second)
	}
}
