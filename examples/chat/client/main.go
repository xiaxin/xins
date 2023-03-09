package main

import (
	"fmt"
	"os"
	"xins/examples/chat/client/cmd"

	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use: "main",
	}

	rootCmd.AddCommand(cmd.LoginCommand())

	// TODO
	rootCmd.AddCommand(cmd.ReadCommand())
	rootCmd.AddCommand(cmd.PingCommand())
	rootCmd.AddCommand(cmd.GroupMessageCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
