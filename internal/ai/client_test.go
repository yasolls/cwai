package ai

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectModelFamily_Standard(t *testing.T) {
	for _, model := range []string{"gpt-4o", "gpt-4o-mini", "claude-3.5-sonnet", "deepseek-chat"} {
		assert.Equal(t, ModelFamilyStandard, DetectModelFamily(model), model)
	}
}

func TestDetectModelFamily_Reasoning(t *testing.T) {
	for _, model := range []string{
		"o1", "o1-preview", "o3", "o3-mini", "o4-mini",
		"deepseek-reasoner", "deepseek-r1",
		"qwq-32b",
		"gemini-2.5-pro", "gemini-2.5-flash", "gemini-3-flash",
		"glm-4.6", "glm-4.7",
		"kimi-k2-thinking", "qwen3-235b-thinking",
	} {
		assert.Equal(t, ModelFamilyReasoning, DetectModelFamily(model), model)
	}
}

func TestDetectModelFamily_GPT5(t *testing.T) {
	for _, model := range []string{"gpt-5", "gpt-5-0827"} {
		assert.Equal(t, ModelFamilyGPT5, DetectModelFamily(model), model)
	}
}

func TestNewClient_DefaultTokensOverrideForReasoning(t *testing.T) {
	c := NewClient(Params{
		Model:           "o3-mini",
		MaxTokensOutput: 500,
	})
	assert.Equal(t, DefaultReasoningMaxTokensOutput, c.params.MaxTokensOutput)
}

func TestNewClient_DefaultTokensOverrideForGPT5(t *testing.T) {
	c := NewClient(Params{
		Model:           "gpt-5",
		MaxTokensOutput: 500,
	})
	assert.Equal(t, DefaultReasoningMaxTokensOutput, c.params.MaxTokensOutput)
}

func TestNewClient_NoOverrideWhenExplicitlySet(t *testing.T) {
	c := NewClient(Params{
		Model:              "o3-mini",
		MaxTokensOutput:    200,
		HasMaxTokensOutput: true,
	})
	assert.Equal(t, 200, c.params.MaxTokensOutput)
}

func TestNewClient_NoOverrideForStandardModel(t *testing.T) {
	c := NewClient(Params{
		Model:           "gpt-4o",
		MaxTokensOutput: 500,
	})
	assert.Equal(t, 500, c.params.MaxTokensOutput)
}

func TestTokenParamName_OpenRouter(t *testing.T) {
	assert.Equal(t, "max_tokens", tokenParamName("https://openrouter.ai/api/v1"))
}

func TestTokenParamName_DeepSeek(t *testing.T) {
	assert.Equal(t, "max_tokens", tokenParamName("https://api.deepseek.com/v1"))
}

func TestTokenParamName_OpenAI(t *testing.T) {
	assert.Equal(t, "max_completion_tokens", tokenParamName("https://api.openai.com/v1"))
}

func TestTokenParamName_Unknown(t *testing.T) {
	assert.Equal(t, "max_completion_tokens", tokenParamName("https://my-llm.example.com"))
}
