package management

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestPrompt(t *testing.T) {
	setupVCR(t)

	t.Cleanup(func() {
		err := m.Prompt.Update(&Prompt{
			UniversalLoginExperience: "classic",
			IdentifierFirst:          auth0.Bool(false),
		})
		require.NoError(t, err)
	})

	t.Run("Update to the new identifier first experience", func(t *testing.T) {
		err := m.Prompt.Update(&Prompt{
			UniversalLoginExperience: "new",
			IdentifierFirst:          auth0.Bool(true),
		})
		assert.NoError(t, err)

		ps, err := m.Prompt.Read()
		assert.NoError(t, err)
		assert.Equal(t, "new", ps.UniversalLoginExperience)
		assert.Equal(t, true, ps.GetIdentifierFirst())
	})

	t.Run("Update to the classic non identifier first experience", func(t *testing.T) {
		err := m.Prompt.Update(&Prompt{
			UniversalLoginExperience: "classic",
			IdentifierFirst:          auth0.Bool(false),
		})
		assert.NoError(t, err)

		ps, err := m.Prompt.Read()
		assert.NoError(t, err)
		assert.Equal(t, "classic", ps.UniversalLoginExperience)
		assert.Equal(t, false, ps.GetIdentifierFirst())
	})
}

func TestPromptCustomText(t *testing.T) {
	setupVCR(t)

	const prompt = "login"
	const lang = "en"

	t.Cleanup(func() {
		body := make(map[string]interface{})
		err := m.Prompt.SetCustomText(prompt, lang, body)
		require.NoError(t, err)
	})

	body := map[string]interface{}{
		"login": map[string]interface{}{
			"title": "Welcome",
		},
	}

	err := m.Prompt.SetCustomText(prompt, lang, body)
	assert.NoError(t, err)

	texts, err := m.Prompt.CustomText(prompt, lang)
	assert.NoError(t, err)
	assert.Equal(t, "Welcome", texts["login"].(map[string]interface{})["title"])
}
