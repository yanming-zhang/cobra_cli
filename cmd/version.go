package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"cobra_cli/pkg/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本信息",
	Long:  "显示版本信息",
	Run: func(cmd *cobra.Command, args []string) {
		v := version.Get()
		marshalled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
	},
}
