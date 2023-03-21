package cmd

import (
	"errors"
	"fmt"
	"io"
	"xins/core"

	"github.com/spf13/cobra"
)

func ReadCommand() *cobra.Command {
	return &cobra.Command{
		Use: "read",
		Run: read,
	}
}

func read(cmd *cobra.Command, args []string) {
	conn, err := conn()
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	for {

		message, err := protocol.Unpack(conn)

		if nil != err {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("[recv error] [read head] %s\n", err.Error())
			continue
		}

		fmt.Println(string(message.(*core.Message).Data()))
	}
}
