package prompt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildMessages_Structured(t *testing.T) {
	msgs := BuildMessages("English", "diff content", true)
	require.Len(t, msgs, 6)
	assert.Equal(t, "system", msgs[0].Role)
	assert.Equal(t, "user", msgs[1].Role)
	assert.Equal(t, "assistant", msgs[2].Role)
	assert.Equal(t, "user", msgs[3].Role)
	assert.Equal(t, "assistant", msgs[4].Role)
	assert.Equal(t, "user", msgs[5].Role)
	assert.Equal(t, "diff content", msgs[5].Content)
	assert.Contains(t, msgs[0].Content, "English")
}

func TestBuildMessages_Standard(t *testing.T) {
	msgs := BuildMessages("Russian", "my diff", false)
	require.Len(t, msgs, 4)
	assert.Equal(t, "system", msgs[0].Role)
	assert.Equal(t, "user", msgs[1].Role)
	assert.Equal(t, "assistant", msgs[2].Role)
	assert.Equal(t, "user", msgs[3].Role)
	assert.Equal(t, "my diff", msgs[3].Content)
	assert.Contains(t, msgs[0].Content, "Russian")
}

func TestBuildMessages_DiffInLastMessage(t *testing.T) {
	diff := "diff --git a/file.go b/file.go\n+new line"
	for _, structured := range []bool{true, false} {
		msgs := BuildMessages("English", diff, structured)
		last := msgs[len(msgs)-1]
		assert.Equal(t, "user", last.Role)
		assert.Equal(t, diff, last.Content)
	}
}
