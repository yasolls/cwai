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
	fmt.Println("cwai generates commit messages from your staged changes using an AI model.")
	fmt.Println("You'll need an API key from your AI provider to get started.")
	fmt.Println()

	fmt.Println("Step 1/4: API Key")
	fmt.Println()
	fmt.Println("  Where to get one:")
	fmt.Println("    OpenAI:     https://platform.openai.com/api-keys")
	fmt.Println("    OpenRouter: https://openrouter.ai/keys")
	fmt.Println("    DeepSeek:   https://platform.deepseek.com/api_keys")
	fmt.Println()
	apiKey := promptInput(reader, "  API Key (CWAI_API_KEY)", "")
	if apiKey == "" {
		return fmt.Errorf("API key is required")
	}
	if len(apiKey) < 10 {
		fmt.Println()
		fmt.Println("  Warning: API key seems unusually short. Continuing anyway.")
	}
	if err := config.Set("CWAI_API_KEY", apiKey); err != nil {
		return err
	}
	fmt.Println()

	fmt.Println("Step 2/4: API URL")
	fmt.Println()
	fmt.Println("  The base URL of your AI provider's API.")
	fmt.Println("  Change this only if you use a non-OpenAI provider or a local model server.")
	fmt.Println()
	apiURL := promptInput(reader, "  API URL (CWAI_API_URL)", config.DefaultAPIURL)
	if err := config.Set("CWAI_API_URL", apiURL); err != nil {
		return err
	}
	fmt.Println()

	fmt.Println("Step 3/4: Model")
	fmt.Println()
	fmt.Println("  The AI model to use for generating commit messages.")
	fmt.Println("  Common options: gpt-5-mini, deepseek-chat, gemini-2.5-flash")
	fmt.Println()
	model := promptInput(reader, "  Model (CWAI_MODEL)", config.DefaultModel)
	if err := config.Set("CWAI_MODEL", model); err != nil {
		return err
	}
	fmt.Println()

	fmt.Println("Step 4/4: Language")
	fmt.Println()
	fmt.Println("  Language for commit messages (ISO 639-1 code: en, de, fr, es, etc.)")
	fmt.Println()
	language := promptInput(reader, "  Language (CWAI_LANGUAGE)", config.DefaultLanguage)
	if err := config.Set("CWAI_LANGUAGE", language); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Configuration saved to ~/.cwai")
	fmt.Println("You can now run 'cwai' in any git repo to generate commit messages.")
	fmt.Println()
	fmt.Println("Tip: For advanced settings (token limits, temperature, structured output),")
	fmt.Println("     use 'cwai config set KEY VALUE'. Run 'cwai config --help' for details.")
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
