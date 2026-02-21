package cmd

import (
	"fmt"

	"github.com/nikmd1306/cwai/internal/hook"
	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Manage git prepare-commit-msg hook",
}

var hookSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Install cwai as prepare-commit-msg hook",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := hook.Set(); err != nil {
			return err
		}
		fmt.Println("cwai hook installed successfully.")
		fmt.Println("Now 'git commit' will auto-generate commit messages.")
		return nil
	},
}

var hookUnsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Remove cwai prepare-commit-msg hook",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := hook.Unset(); err != nil {
			return err
		}
		fmt.Println("cwai hook removed successfully.")
		return nil
	},
}

func init() {
	hookCmd.AddCommand(hookSetCmd)
	hookCmd.AddCommand(hookUnsetCmd)
}
