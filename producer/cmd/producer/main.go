package main

import (
	"context"
	"producer/internal/busi"

	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// @title abcdefg producer api backend
// @version 1.0
// @description abcdefg producer api backend
// @termsOfService http://swagger.io/terms/

// @contact.name abcdefg
// @contact.email aaaaa

// @host 127.0.0.1:7005
// @BasePath /

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "producer",
		Short: "producer",
		Run: func(cmd *cobra.Command, args []string) {
			if err := entry(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
	}

	cmd.PersistentFlags().StringVar(&busi.Flags.Config, "conf", "", "path of the configuration file")

	return cmd
}

func entry() error {
	busi.NewServer(context.Background()).Start()
	return nil
}

func main() {
	if err := NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
