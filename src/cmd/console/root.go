package console

import (
	"fmt"
	"github.com/alimarzban99/go-blog-api/internal/commands"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-blog-api",
	Short: "Go Blog API CLI",
	Long:  `A command line interface for managing the go blog api.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(commands.ServeCmd, commands.MigrateCmd)
}
