package cmd

import (
	"fmt"

	"github.com/nikmd1306/cwai/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage cwai configuration",
}

var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, value := args[0], args[1]

		if !isValidKey(key) {
			return fmt.Errorf("unknown config key: %s", key)
		}

		if err := config.Set(key, value); err != nil {
			return fmt.Errorf("failed to set config: %w", err)
		}

		fmt.Printf("%s = %s\n", key, value)
		return nil
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a configuration value",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		if !isValidKey(key) {
			return fmt.Errorf("unknown config key: %s", key)
		}

		value, err := config.Get(key)
		if err != nil {
			return fmt.Errorf("failed to get config: %w", err)
		}

		if value == "" {
			fmt.Printf("%s is not set\n", key)
		} else {
			fmt.Printf("%s = %s\n", key, value)
		}
		return nil
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configGetCmd)
}

var validKeys = map[string]bool{
	"CWAI_API_KEY":            true,
	"CWAI_API_URL":            true,
	"CWAI_MODEL":              true,
	"CWAI_LANGUAGE":           true,
	"CWAI_MAX_TOKENS_INPUT":   true,
	"CWAI_MAX_TOKENS_OUTPUT":  true,
	"CWAI_TEMPERATURE":        true,
	"CWAI_REASONING_EFFORT":   true,
	"CWAI_VERBOSITY":          true,
	"CWAI_STRUCTURED_OUTPUT":  true,
}

func isValidKey(key string) bool {
	return validKeys[key]
}
