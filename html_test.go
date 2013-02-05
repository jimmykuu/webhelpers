package webhelpers

import (
	"testing"
)

func TestStripTags(t *testing.T) {
	html := "Text <em>emphasis</em> <script>Javascript</script>."
	text := "Text emphasis Javascript."
	actual := StripTags(html)
	if actual != text {
		t.Errorf("StripTags(`%s`) want `%s` not `%s`", html, text, actual)
	}
}

func TestRemoveFormatting(t *testing.T) {
	html := `<p>How to use the <a href="/cmd/go/">go command</a> to fetch, build, and install packages, commands, and run tests.</p>`
	text := "How to use the go command to fetch, build, and install packages, commands, and run tests."
	actual := RemoveFormatting(html)
	if actual != text {
		t.Errorf("RemoveFormatting(`%s`) want `%s` not `%s`", html, text, actual)
	}
}
