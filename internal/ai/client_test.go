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
	for _, model := range []string{"o1", "o1-preview", "o3", "o3-mini", "o4-mini"} {
		assert.Equal(t, ModelFamilyReasoning, DetectModelFamily(model), model)
	}
}

func TestDetectModelFamily_GPT5(t *testing.T) {
	for _, model := range []string{"gpt-5", "gpt-5-0827"} {
		assert.Equal(t, ModelFamilyGPT5, DetectModelFamily(model), model)
	}
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
