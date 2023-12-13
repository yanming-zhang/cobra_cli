package cmd

import (
	"fmt"
	"os"
	"strings"

	"cobra_cli/pkg/gitClient"
	"cobra_cli/pkg/shell"
	"github.com/spf13/cobra"
)

var dependDesc = strings.Join([]string{
	"目前所需要安装的依赖包如下：",
	"1) git 		版本：^2",
	"2) go  		版本：^1.21",
	"3) node 		版本：^20",
	"4) pnpm 		版本：^8",
	"5) foundry 	版本：^0.2.0",
	"6) make 		版本：^4",
	"7) jq			版本：^1.6",
	"8) direnv		版本：^2",
}, "\n")

func init() {
	rootCmd.AddCommand(dependCmd)
}

var dependCmd = &cobra.Command{
	Use:   "depend",
	Short: "安装依赖包",
	Long:  dependDesc,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := gitClient.Clone("develop", "/data/optimism", "https://github.com/ethereum-optimism/optimism")
		if err != nil {
			fmt.Sprintf("error: %s", err)
			os.Exit(1)
		}
		_, err = gitClient.Clone("optimism", "/data/op-geth", "https://github.com/ethereum-optimism/op-geth")
		if err != nil {
			fmt.Sprintf("error: %s", err)
			os.Exit(1)
		}
		output, _ := shell.ExecuteShell([]string{"/data/optimism/packages/contracts-bedrock/scripts/getting-started/versions.sh"})
		if strings.Contains(output, "No version found") {
			_ = shell.ExecuteShellStdout([]string{"scripts/install_dependency_pkgs.sh"})
		}
	},
}
