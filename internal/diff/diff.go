package diff

import (
	"fmt"
	"strings"

	"github.com/nikmd1306/cwai/internal/git"
)

func estimateTokens(text string) int {
	return len(text) / 4
}

func Truncate(maxTokens int) (string, error) {
	files, err := git.StagedNumStat()
	if err != nil {
		return "", err
	}
	if len(files) == 0 {
		return "", fmt.Errorf("no staged changes found")
	}

	var header strings.Builder
	header.WriteString("Files changed:\n")
	for _, f := range files {
		fmt.Fprintf(&header, "  %s (+%s/-%s)\n", f.File, f.Added, f.Removed)
	}
	header.WriteString("\n")

	headerStr := header.String()
	budgetTokens := maxTokens - estimateTokens(headerStr)
	if budgetTokens < 100 {
		return headerStr + "[Diff truncated. All file diffs omitted due to token limit.]", nil
	}

	var result strings.Builder
	result.WriteString(headerStr)

	included := 0
	for i, f := range files {
		fileDiff, err := git.StagedDiffForFile(f.File)
		if err != nil {
			continue
		}

		diffTokens := estimateTokens(fileDiff)
		if diffTokens > budgetTokens {
			remaining := len(files) - i
			fmt.Fprintf(&result, "[Diff truncated. %d more files not shown.]\n", remaining)
			break
		}

		result.WriteString(fileDiff)
		result.WriteString("\n")
		budgetTokens -= diffTokens
		included++
	}

	if included == 0 && len(files) > 0 {
		fmt.Fprintf(&result, "[Diff truncated. %d files not shown.]\n", len(files))
	}

	return result.String(), nil
}
