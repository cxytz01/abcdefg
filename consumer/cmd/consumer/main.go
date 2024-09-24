package main

import (
	"context"
	"fmt"
	"os"

	"consumer/internal"

	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "consumer",
		Short: "consumer",
		Run: func(cmd *cobra.Command, args []string) {
			if err := entry(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
	}

	flags := cmd.PersistentFlags()
	flags.StringVar(&internal.Flags.Config, "conf", "", "path of the configuration file")

	return cmd
}

func entry() error {
	internal.NewServer(context.Background()).Start()
	return nil
}

func main() {
	if err := NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
