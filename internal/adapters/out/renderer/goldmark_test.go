package renderer_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/douglas/markdown-server/internal/adapters/out/renderer"
)

func TestGoldmarkRenderer(t *testing.T) {
	r := renderer.NewGoldmarkRenderer()
	ctx := context.Background()

	t.Run("Given markdown with H1 title When RenderHTML is called Then title and HTML are returned", func(t *testing.T) {
		md := []byte("# Manual do Usuário\n\nEste é o manual de uso.")
		html, title, err := r.RenderHTML(ctx, md, "docs/manual.md")

		require.NoError(t, err)
		assert.Equal(t, "Manual do Usuário", title)
		assert.Contains(t, html, "Manual do Usuário")
		assert.Contains(t, html, "<p>Este é o manual de uso.</p>")
	})

	t.Run("Given markdown without H1 When RenderHTML is called Then title is empty", func(t *testing.T) {
		md := []byte("## Subtítulo sem H1\n\nTexto qualquer.")
		html, title, err := r.RenderHTML(ctx, md, "docs/sub.md")

		require.NoError(t, err)
		assert.Empty(t, title)
		assert.Contains(t, html, "Subtítulo sem H1")
	})

	t.Run("Given GFM tables and task lists When RenderHTML is called Then table tags are rendered", func(t *testing.T) {
		md := []byte(`
| Nome | Tipo |
| :--- | :--- |
| Go   | Backend |

- [x] Concluído
- [ ] Pendente
`)
		html, _, err := r.RenderHTML(ctx, md, "tables.md")

		require.NoError(t, err)
		assert.Contains(t, html, "<table>")
		assert.Contains(t, html, "Nome")
		assert.Contains(t, html, "Go")
		assert.Contains(t, html, "Backend")
		assert.Contains(t, html, `type="checkbox"`)
	})

	t.Run("Given a Go code block When RenderHTML is called Then Chroma syntax highlighting is applied", func(t *testing.T) {
		md := []byte("```go\npackage main\n\nfunc main() {}\n```")
		html, _, err := r.RenderHTML(ctx, md, "code.md")

		require.NoError(t, err)
		assert.Contains(t, html, `<div class="highlight go">`)
		assert.Contains(t, html, "package")
	})

	t.Run("Given a Mermaid code block When RenderHTML is called Then it is wrapped in div mermaid", func(t *testing.T) {
		md := []byte("```mermaid\nflowchart TD\n  A --> B\n```")
		html, _, err := r.RenderHTML(ctx, md, "diagram.md")

		require.NoError(t, err)
		assert.Contains(t, html, `<div class="mermaid">flowchart TD`)
		assert.Contains(t, html, `A --&gt; B</div>`)
	})

	t.Run("Given a MERMAID uppercase code block When RenderHTML is called Then it is also wrapped in div mermaid", func(t *testing.T) {
		md := []byte("```MERMAID\nsequenceDiagram\n  Alice->>Bob: Hello\n```")
		html, _, err := r.RenderHTML(ctx, md, "sequence.md")

		require.NoError(t, err)
		assert.Contains(t, html, `<div class="mermaid">sequenceDiagram`)
		assert.Contains(t, html, `Alice-&gt;&gt;Bob: Hello</div>`)
	})

	t.Run("Given relative markdown links When RenderHTML is called Then links are transformed to web routes", func(t *testing.T) {
		md := []byte(`
[Guia](./guide.md)
[Docs](docs/intro.md)
[README](../README.md)
[Externo](https://golang.org)
[Âncora](#secao)
`)
		html, _, err := r.RenderHTML(ctx, md, "articles/post.md")

		require.NoError(t, err)
		assert.Contains(t, html, `href="/articles/guide.md"`)
		assert.Contains(t, html, `href="/articles/docs/intro.md"`)
		assert.Contains(t, html, `href="/"`)
		assert.Contains(t, html, `href="https://golang.org"`)
		assert.Contains(t, html, `href="#secao"`)
	})

	t.Run("Given a generic code block without language When RenderHTML is called Then it renders standard pre code", func(t *testing.T) {
		md := []byte("```\nplain text content\n```")
		html, _, err := r.RenderHTML(ctx, md, "plain.md")

		require.NoError(t, err)
		assert.Contains(t, html, "<pre><code>")
		assert.Contains(t, html, "plain text content")
	})

	t.Run("Given a cancelled context When RenderHTML is called Then it returns context error", func(t *testing.T) {
		cancelCtx, cancel := context.WithCancel(context.Background())
		cancel()

		_, _, err := r.RenderHTML(cancelCtx, []byte("# Test"), "test.md")
		require.Error(t, err)
		assert.ErrorIs(t, err, context.Canceled)
	})
}

func TestTransformMarkdownLink(t *testing.T) {
	t.Run("Given external, anchor or special links When TransformMarkdownLink is called Then it returns unchanged", func(t *testing.T) {
		assert.Equal(t, "https://google.com", renderer.TransformMarkdownLink("https://google.com", "docs"))
		assert.Equal(t, "http://localhost", renderer.TransformMarkdownLink("http://localhost", "docs"))
		assert.Equal(t, "mailto:user@test.com", renderer.TransformMarkdownLink("mailto:user@test.com", "docs"))
		assert.Equal(t, "#secao-1", renderer.TransformMarkdownLink("#secao-1", "docs"))
		assert.Equal(t, "//cdn.example.com", renderer.TransformMarkdownLink("//cdn.example.com", "docs"))
		assert.Equal(t, "image.png", renderer.TransformMarkdownLink("image.png", "docs"))
	})

	t.Run("Given markdown relative paths When TransformMarkdownLink is called Then it resolves correctly", func(t *testing.T) {
		assert.Equal(t, "/docs/api.md", renderer.TransformMarkdownLink("./api.md", "docs"))
		assert.Equal(t, "/docs/api.md#v1", renderer.TransformMarkdownLink("./api.md#v1", "docs"))
		assert.Equal(t, "/", renderer.TransformMarkdownLink("README.md", ""))
		assert.Equal(t, "/", renderer.TransformMarkdownLink("/README.md", "docs"))
		assert.Equal(t, "/guide.md", renderer.TransformMarkdownLink("/guide.md", "docs"))
	})
}
