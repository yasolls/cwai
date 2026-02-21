package ai

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseCommitMessageJSON_Valid(t *testing.T) {
	raw := `{"changes_summary":"s","introduces_new_behavior":true,"fixes_broken_behavior":false,"restructures_only":false,"type_reasoning":"r","type":"feat","scope":"api","description":"add endpoint","bullet_points":["one","two"]}`
	resp, err := ParseCommitMessageJSON(raw)
	require.NoError(t, err)
	assert.Equal(t, "feat", resp.Type)
	assert.Equal(t, "api", resp.Scope)
	assert.Equal(t, "add endpoint", resp.Description)
	assert.True(t, resp.IntroducesNewBehavior)
	assert.Equal(t, []string{"one", "two"}, resp.BulletPoints)
}

func TestParseCommitMessageJSON_Invalid(t *testing.T) {
	_, err := ParseCommitMessageJSON("not json")
	assert.Error(t, err)
}

func TestParseCommitMessageJSON_Partial(t *testing.T) {
	raw := `{"type":"fix","scope":"db","description":"fix query"}`
	resp, err := ParseCommitMessageJSON(raw)
	require.NoError(t, err)
	assert.Equal(t, "fix", resp.Type)
	assert.Equal(t, "db", resp.Scope)
	assert.Empty(t, resp.BulletPoints)
}

func TestAssembleCommitMessage_HeaderOnly(t *testing.T) {
	resp := CommitMessageResponse{
		Type:        "fix",
		Scope:       "db",
		Description: "resolve deadlock",
	}
	assert.Equal(t, "fix(db): resolve deadlock", AssembleCommitMessage(resp))
}

func TestAssembleCommitMessage_LowercaseFirstChar(t *testing.T) {
	resp := CommitMessageResponse{
		Type:        "feat",
		Scope:       "api",
		Description: "Add new endpoint",
	}
	result := AssembleCommitMessage(resp)
	assert.True(t, strings.HasPrefix(result, "feat(api): add"))
}

func TestAssembleCommitMessage_TrailingDotRemoved(t *testing.T) {
	resp := CommitMessageResponse{
		Type:        "fix",
		Scope:       "ui",
		Description: "fix button color.",
	}
	assert.Equal(t, "fix(ui): fix button color", AssembleCommitMessage(resp))
}

func TestAssembleCommitMessage_TruncatesLongDescription(t *testing.T) {
	long := strings.Repeat("a", 100)
	resp := CommitMessageResponse{
		Type:        "feat",
		Scope:       "x",
		Description: long,
	}
	result := AssembleCommitMessage(resp)
	desc := strings.TrimPrefix(result, "feat(x): ")
	assert.LessOrEqual(t, len(desc), 72)
}

func TestAssembleCommitMessage_WithBulletPoints(t *testing.T) {
	resp := CommitMessageResponse{
		Type:         "refactor",
		Scope:        "auth",
		Description:  "restructure auth flow",
		BulletPoints: []string{"extract validator", "rename handler", "update tests"},
	}
	result := AssembleCommitMessage(resp)
	assert.Contains(t, result, "refactor(auth): restructure auth flow")
	assert.Contains(t, result, "\n\n- extract validator\n- rename handler\n- update tests")
}

func TestAssembleCommitMessage_SingleBulletPoint(t *testing.T) {
	resp := CommitMessageResponse{
		Type:         "fix",
		Scope:        "db",
		Description:  "fix query",
		BulletPoints: []string{"add index"},
	}
	result := AssembleCommitMessage(resp)
	assert.Contains(t, result, "\n\n- add index")
	assert.NotContains(t, result, "\n- add index\n")
}

func TestSupportsStructuredOutput_OverrideOn(t *testing.T) {
	assert.True(t, SupportsStructuredOutput("https://example.com", "on"))
}

func TestSupportsStructuredOutput_OverrideOff(t *testing.T) {
	assert.False(t, SupportsStructuredOutput("https://api.openai.com/v1", "off"))
}

func TestSupportsStructuredOutput_OpenAI(t *testing.T) {
	assert.True(t, SupportsStructuredOutput("https://api.openai.com/v1", ""))
}

func TestSupportsStructuredOutput_OpenRouter(t *testing.T) {
	assert.True(t, SupportsStructuredOutput("https://openrouter.ai/api/v1", ""))
}

func TestSupportsStructuredOutput_UnknownURL(t *testing.T) {
	assert.False(t, SupportsStructuredOutput("https://my-llm.example.com/v1", ""))
}
