package cmd

import (
	"log"
	"strings"

	"cobra_cli/pkg/shell"
	"github.com/spf13/cobra"
)

var buildDesc = strings.Join([]string{
	"build的软件类型如下：",
	"1) optimism",
	"2) op-geth",
}, "\n")

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build程序包",
	Long:  buildDesc,
	Run: func(cmd *cobra.Command, args []string) {
		switch bt {
		case "optimism":
			_ = shell.ExecuteShellStdout([]string{"scripts/build_pkgs.sh", "optimism"})
		case "op-geth":
			_ = shell.ExecuteShellStdout([]string{"scripts/build_pkgs.sh", "op-geth"})
		default:
			log.Fatalf("暂不支持该build类型")
		}
	},
}

var bt string

func init() {
	buildCmd.Flags().StringVarP(&bt, "build_type", "t", "", "build类型")
}
