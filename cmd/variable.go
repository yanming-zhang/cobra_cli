package cmd

import (
	"fmt"
	"os"
	"strings"

	"cobra_cli/pkg/shell"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

var varDesc = strings.Join([]string{
	"主要环境变量如下：",
	"1) L1_RPC_URL",
	"2) L1_RPC_KIND",
	"3) GS_ADMIN_ADDRESS",
	"4) GS_ADMIN_PRIVATE_KEY",
	"5) GS_BATCHER_ADDRESS",
	"6) GS_BATCHER_PRIVATE_KEY",
	"7) GS_PROPOSER_ADDRESS",
	"8) GS_PROPOSER_PRIVATE_KEY",
	"9) GS_SEQUENCER_ADDRESS",
	"10) GS_SEQUENCER_PRIVATE_KEY",
}, "\n")

func init() {
	rootCmd.AddCommand(envCmd)
}

var envCmd = &cobra.Command{
	Use:   "env_vars",
	Short: "设置环境变量",
	Long:  varDesc,
	Run: func(cmd *cobra.Command, args []string) {
		ini.PrettyFormat = false

		output, _ := shell.ExecuteShell([]string{"scripts/env_variables.sh"})
		cfg, err := ini.Load([]byte(output))
		if err != nil {
			fmt.Printf("Fail to read values: %v", err)
			os.Exit(1)
		}
		fmt.Println(output)

		c, err := ini.Load("/data/optimism/.envrc")
		if err != nil {
			fmt.Printf("Fail to read file: %v", err)
			os.Exit(1)
		}
		adminAddress := cfg.Section("").Key("export GS_ADMIN_ADDRESS").String()
		c.Section("").Key("export GS_ADMIN_ADDRESS").SetValue(adminAddress)
		adminPrivate := cfg.Section("").Key("export GS_ADMIN_PRIVATE_KEY").String()
		c.Section("").Key("export GS_ADMIN_PRIVATE_KEY").SetValue(adminPrivate)

		batcherAddress := cfg.Section("").Key("export GS_BATCHER_ADDRESS").String()
		c.Section("").Key("export GS_BATCHER_ADDRESS").SetValue(batcherAddress)
		batcherPrivate := cfg.Section("").Key("export GS_BATCHER_PRIVATE_KEY").String()
		c.Section("").Key("export GS_BATCHER_PRIVATE_KEY").SetValue(batcherPrivate)

		proposerAddress := cfg.Section("").Key("export GS_PROPOSER_ADDRESS").String()
		c.Section("").Key("export GS_PROPOSER_ADDRESS").SetValue(proposerAddress)
		proposerPrivate := cfg.Section("").Key("export GS_PROPOSER_PRIVATE_KEY").String()
		c.Section("").Key("export GS_PROPOSER_PRIVATE_KEY").SetValue(proposerPrivate)

		sequencerAddress := cfg.Section("").Key("export GS_SEQUENCER_ADDRESS").String()
		c.Section("").Key("export GS_SEQUENCER_ADDRESS").SetValue(sequencerAddress)
		sequencerPrivate := cfg.Section("").Key("export GS_SEQUENCER_PRIVATE_KEY").String()
		c.Section("").Key("export GS_SEQUENCER_PRIVATE_KEY").SetValue(sequencerPrivate)

		lc, err := ini.Load("./config/cfg.ini")
		if err != nil {
			fmt.Printf("Fail to read file: %v", err)
			os.Exit(1)
		}
		url := lc.Section("L1_RPC").Key("URL").String()
		c.Section("").Key("export L1_RPC_URL").SetValue(url)
		kind := lc.Section("L1_RPC").Key("KIND").String()
		c.Section("").Key("export L1_RPC_KIND").SetValue(kind)

		c.SaveTo("/data/optimism/.envrc")
		_ = shell.ExecuteShellStdout([]string{"scripts/effect_env_variables.sh"})
	},
}
