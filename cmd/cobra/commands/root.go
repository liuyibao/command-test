package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:                        "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: "Hugo is a very fast static site generator",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")
	},
}

var name *string

func init()  {
	name = rootCmd.PersistentFlags().StringP("name", "n", "YOUR NAME", "Author name for copyright attribution")
	rootCmd.Flags()
}

func Execute()  {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(*name)
}