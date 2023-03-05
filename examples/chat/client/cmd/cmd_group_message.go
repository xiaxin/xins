package cmd

import (
	"fmt"
	"time"
	"xins/examples/chat/object"

	"github.com/spf13/cobra"
)

func GroupMessageCommand() *cobra.Command {
	return &cobra.Command{
		Use: "group_message",
		Run: groupMessage,
	}
}

func groupMessage(cmd *cobra.Command, args []string) {
	conn, err := conn()
	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	for {
		groupMessage := object.NewGroupMessage("1", "user-a", "abc")
		message, err := protocol.NewMessage(12, groupMessage)

		if nil != err {
			fmt.Println("[new-message] error ", err.Error())
			continue
		}

		m, _ := protocol.Pack(message)

		conn.Write(m)

		time.Sleep(time.Second)
	}
}
