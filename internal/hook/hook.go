package hook

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nikmd1306/cwai/internal/git"
)

const marker = "# Installed by: cwai hook set"

const hookScript = `#!/bin/sh
# Installed by: cwai hook set
if command -v cwai >/dev/null 2>&1; then
    cwai --hook "$1" "$2" "$3"
fi
`

func hookPath() (string, error) {
	hooksDir, err := git.HooksPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(hooksDir, "prepare-commit-msg"), nil
}

func Set() error {
	path, err := hookPath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create hooks directory: %w", err)
	}

	if data, err := os.ReadFile(path); err == nil {
		if strings.Contains(string(data), marker) {
			return fmt.Errorf("cwai hook is already installed")
		}
		return fmt.Errorf("prepare-commit-msg hook already exists (not managed by cwai). Remove it first or use a different approach")
	}

	if err := os.WriteFile(path, []byte(hookScript), 0o755); err != nil {
		return fmt.Errorf("write hook: %w", err)
	}

	return nil
}

func Unset() error {
	path, err := hookPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("no prepare-commit-msg hook found")
		}
		return fmt.Errorf("read hook: %w", err)
	}

	if !strings.Contains(string(data), marker) {
		return fmt.Errorf("prepare-commit-msg hook exists but was not installed by cwai")
	}

	if err := os.Remove(path); err != nil {
		return fmt.Errorf("remove hook: %w", err)
	}

	return nil
}
