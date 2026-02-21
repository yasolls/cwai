package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nikmd1306/cwai/internal/config"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Interactive configuration setup",
	RunE:  runSetup,
}

func runSetup(cmd *cobra.Command, args []string) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("cwai setup")
	fmt.Println("----------")
	fmt.Println()

	apiKey := promptInput(reader, "API Key (CWAI_API_KEY)", "")
	if apiKey == "" {
		return fmt.Errorf("API key is required")
	}
	if err := config.Set("CWAI_API_KEY", apiKey); err != nil {
		return err
	}

	apiURL := promptInput(reader, "API URL (CWAI_API_URL)", config.DefaultAPIURL)
	if err := config.Set("CWAI_API_URL", apiURL); err != nil {
		return err
	}

	model := promptInput(reader, "Model (CWAI_MODEL)", config.DefaultModel)
	if err := config.Set("CWAI_MODEL", model); err != nil {
		return err
	}

	language := promptInput(reader, "Language (CWAI_LANGUAGE)", config.DefaultLanguage)
	if err := config.Set("CWAI_LANGUAGE", language); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Configuration saved to ~/.cwai")
	fmt.Println("You can now use 'cwai' to generate commit messages.")
	return nil
}

func promptInput(reader *bufio.Reader, label, defaultVal string) string {
	if defaultVal != "" {
		fmt.Printf("%s [%s]: ", label, defaultVal)
	} else {
		fmt.Printf("%s: ", label)
	}

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		return defaultVal
	}
	return input
}
