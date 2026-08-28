package web_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/douglas/markdown-server/web"
)

func TestEmbeddedFS(t *testing.T) {
	t.Run("Given embedded filesystem When templates are read Then base, document and 404 exist", func(t *testing.T) {
		base, err := web.EmbeddedFS.ReadFile("templates/base.html")
		require.NoError(t, err)
		assert.NotEmpty(t, base)
		assert.Contains(t, string(base), "Markdown Viewer")

		doc, err := web.EmbeddedFS.ReadFile("templates/document.html")
		require.NoError(t, err)
		assert.NotEmpty(t, doc)

		err404, err := web.EmbeddedFS.ReadFile("templates/404.html")
		require.NoError(t, err)
		assert.NotEmpty(t, err404)
	})

	t.Run("Given embedded filesystem When static assets are read Then CSS, JS and Image files exist", func(t *testing.T) {
		css, err := web.EmbeddedFS.ReadFile("static/css/style.css")
		require.NoError(t, err)
		assert.NotEmpty(t, css)

		chroma, err := web.EmbeddedFS.ReadFile("static/css/chroma.css")
		require.NoError(t, err)
		assert.NotEmpty(t, chroma)

		js, err := web.EmbeddedFS.ReadFile("static/js/app.js")
		require.NoError(t, err)
		assert.NotEmpty(t, js)

		mermaid, err := web.EmbeddedFS.ReadFile("static/js/mermaid.min.js")
		require.NoError(t, err)
		assert.NotEmpty(t, mermaid)

		qrcode, err := web.EmbeddedFS.ReadFile("static/js/qrcode.min.js")
		require.NoError(t, err)
		assert.NotEmpty(t, qrcode)

		icon, err := web.EmbeddedFS.ReadFile("static/img/icon.svg")
		require.NoError(t, err)
		assert.NotEmpty(t, icon)
	})
}
