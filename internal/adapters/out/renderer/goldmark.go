package renderer

import (
	"bytes"
	"context"
	"fmt"
	"html"
	"path"
	"strings"

	"github.com/alecthomas/chroma/v2"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"

	"github.com/douglas/markdown-server/internal/core/ports"
)

// GoldmarkRenderer implementa ports.MarkdownRendererPort usando Goldmark, Chroma e suporte a Mermaid.
type GoldmarkRenderer struct {
	markdown goldmark.Markdown
}

// NewGoldmarkRenderer cria uma nova instância configurada com GFM, footnotes, listas de definição, tipografia e cabeçalhos automáticos.
func NewGoldmarkRenderer() *GoldmarkRenderer {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
			extension.Strikethrough,
			extension.Linkify,
			extension.TaskList,
			extension.Footnote,
			extension.DefinitionList,
			extension.Typographer,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			goldmarkhtml.WithUnsafe(),
		),
	)

	return &GoldmarkRenderer{
		markdown: md,
	}
}

// RenderHTML processa o Markdown, executa o AST, reescreve links relativos, destaca sintaxe, formata Mermaid e GitHub Alerts.
func (r *GoldmarkRenderer) RenderHTML(ctx context.Context, source []byte, currentPath string) (string, string, error) {
	if ctx != nil && ctx.Err() != nil {
		return "", "", ctx.Err()
	}

	reader := text.NewReader(source)
	doc := r.markdown.Parser().Parse(reader)

	title := ""
	currentDir := path.Dir(currentPath)
	if currentDir == "." || currentDir == "/" {
		currentDir = ""
	}

	// Caminha pelo AST para extrair o título H1 e transformar links .md
	_ = ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		// Extrai o primeiro H1 como título do documento
		if heading, ok := node.(*ast.Heading); ok && heading.Level == 1 && title == "" {
			var headingText bytes.Buffer
			for child := heading.FirstChild(); child != nil; child = child.NextSibling() {
				if textNode, ok := child.(*ast.Text); ok {
					headingText.Write(textNode.Segment.Value(source))
				}
			}
			title = strings.TrimSpace(headingText.String())
		}

		// Transforma links relativos .md para rotas web
		if link, ok := node.(*ast.Link); ok {
			dest := string(link.Destination)
			transformedDest := TransformMarkdownLink(dest, currentDir)
			link.Destination = []byte(transformedDest)
		}

		return ast.WalkContinue, nil
	})

	var rawHTMLBuf bytes.Buffer
	if err := r.markdown.Renderer().Render(&rawHTMLBuf, source, doc); err != nil {
		return "", "", err
	}

	// Pós-processa blocos de código com Chroma e Mermaid
	processedHTML := postProcessCodeBlocks(rawHTMLBuf.String())

	// Pós-processa GitHub Alerts / Admonitions
	processedHTML = postProcessGitHubAlerts(processedHTML)

	return processedHTML, title, nil
}

// TransformMarkdownLink reescreve URLs relativas que apontam para arquivos .md mantendo a hierarquia web.
func TransformMarkdownLink(dest, currentDir string) string {
	// Links externos, âncoras puras ou esquemas especiais não são transformados
	if strings.HasPrefix(dest, "http://") ||
		strings.HasPrefix(dest, "https://") ||
		strings.HasPrefix(dest, "mailto:") ||
		strings.HasPrefix(dest, "#") ||
		strings.HasPrefix(dest, "//") {
		return dest
	}

	// Separa âncora interna (#secao) da URL se houver
	parts := strings.SplitN(dest, "#", 2)
	filePath := parts[0]
	anchor := ""
	if len(parts) > 1 {
		anchor = "#" + parts[1]
	}

	// Se o link termina com .md ou .markdown, transforma em rota web
	lowerPath := strings.ToLower(filePath)
	if strings.HasSuffix(lowerPath, ".md") || strings.HasSuffix(lowerPath, ".markdown") {
		var resolvedPath string
		if strings.HasPrefix(filePath, "/") {
			resolvedPath = path.Clean(filePath)
		} else {
			if currentDir != "" {
				resolvedPath = path.Clean(path.Join("/", currentDir, filePath))
			} else {
				resolvedPath = path.Clean(path.Join("/", filePath))
			}
		}

		// Rota raiz se for /README.md
		if strings.EqualFold(resolvedPath, "/README.md") || strings.EqualFold(resolvedPath, "/readme.md") {
			return "/" + anchor
		}

		return resolvedPath + anchor
	}

	return dest
}

// postProcessCodeBlocks detecta blocos <pre><code class="language-..."> e aplica Chroma ou container Mermaid.
func postProcessCodeBlocks(htmlContent string) string {
	const preOpen = "<pre><code"
	const preClose = "</code></pre>"

	var result strings.Builder
	idx := 0

	for {
		start := strings.Index(htmlContent[idx:], preOpen)
		if start == -1 {
			result.WriteString(htmlContent[idx:])
			break
		}
		start += idx
		result.WriteString(htmlContent[idx:start])

		end := strings.Index(htmlContent[start:], preClose)
		if end == -1 {
			result.WriteString(htmlContent[start:])
			break
		}
		end += start + len(preClose)

		blockContent := htmlContent[start:end]
		idx = end

		renderedBlock := processSingleCodeBlock(blockContent)
		result.WriteString(renderedBlock)
	}

	return result.String()
}

func processSingleCodeBlock(block string) string {
	// Extrai a classe de linguagem, ex: class="language-go"
	lang := ""
	classIdx := strings.Index(block, `class="language-`)
	if classIdx != -1 {
		classSub := block[classIdx+len(`class="language-`):]
		endQuote := strings.Index(classSub, `"`)
		if endQuote != -1 {
			lang = classSub[:endQuote]
		}
	}

	// Extrai o código bruto entre <code> e </code>
	codeTagIdx := strings.Index(block, "<code")
	if codeTagIdx == -1 {
		return block
	}
	codeStart := strings.Index(block[codeTagIdx:], ">")
	if codeStart == -1 {
		return block
	}
	codeStart = codeTagIdx + codeStart + 1
	codeEnd := strings.LastIndex(block, "</code>")
	if codeEnd == -1 || codeEnd <= codeStart {
		return block
	}

	rawCode := block[codeStart:codeEnd]
	// Desfaz escape básico de HTML gerado pelo parser para obter o código original
	rawCode = html.UnescapeString(rawCode)

	// Se for diagrama Mermaid, encapsula em <div class="mermaid">
	if strings.EqualFold(lang, "mermaid") {
		return fmt.Sprintf(`<div class="mermaid">%s</div>`, html.EscapeString(strings.TrimSpace(rawCode)))
	}

	// Se houver linguagem especificada, usa Chroma para destaque de sintaxe
	if lang != "" {
		lexer := lexers.Get(lang)
		if lexer == nil {
			lexer = lexers.Fallback
		}
		lexer = chroma.Coalesce(lexer)

		style := styles.Get("github-dark")
		if style == nil {
			style = styles.Fallback
		}

		formatter := chromahtml.New(
			chromahtml.WithClasses(true),
			chromahtml.TabWidth(4),
		)

		iterator, err := lexer.Tokenise(nil, rawCode)
		if err == nil {
			var buf bytes.Buffer
			if err := formatter.Format(&buf, style, iterator); err == nil {
				return fmt.Sprintf(`<div class="highlight %s">%s</div>`, lang, buf.String())
			}
		}
	}

	// Fallback padrão
	return block
}

// postProcessGitHubAlerts localiza blocos blockquote que iniciam com [!NOTE], [!TIP], etc., e os converte em alerts estilizados.
func postProcessGitHubAlerts(htmlContent string) string {
	const bqOpen = "<blockquote>"
	const bqClose = "</blockquote>"

	var result strings.Builder
	idx := 0

	for {
		start := strings.Index(htmlContent[idx:], bqOpen)
		if start == -1 {
			result.WriteString(htmlContent[idx:])
			break
		}
		start += idx
		result.WriteString(htmlContent[idx:start])

		end := strings.Index(htmlContent[start:], bqClose)
		if end == -1 {
			result.WriteString(htmlContent[start:])
			break
		}
		end += start + len(bqClose)

		blockContent := htmlContent[start:end]
		idx = end

		renderedBlock := processSingleAlertBlockquote(blockContent)
		result.WriteString(renderedBlock)
	}

	return result.String()
}

func processSingleAlertBlockquote(block string) string {
	inner := strings.TrimPrefix(block, "<blockquote>")
	inner = strings.TrimSuffix(inner, "</blockquote>")
	trimmedInner := strings.TrimSpace(inner)

	if !strings.HasPrefix(trimmedInner, "<p>") {
		return block
	}

	afterP := strings.TrimPrefix(trimmedInner, "<p>")
	trimmedAfterP := strings.TrimSpace(afterP)

	alertTypes := []struct {
		tag   string
		name  string
		class string
		icon  string
	}{
		{
			tag:   "[!NOTE]",
			name:  "Note",
			class: "markdown-alert-note",
			icon:  `<svg class="octicon octicon-info" viewBox="0 0 16 16" width="16" height="16" aria-hidden="true"><path d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8Zm8-6.5a6.5 6.5 0 1 0 0 13 6.5 6.5 0 0 0 0-13ZM6.5 7.75A.75.75 0 0 1 7.25 7h1a.75.75 0 0 1 .75.75v2.75h.25a.75.75 0 0 1 0 1.5h-2a.75.75 0 0 1 0-1.5h.25v-2h-.25a.75.75 0 0 1-.75-.75ZM8 6a1 1 0 1 1 0-2 1 1 0 0 1 0 2Z"></path></svg>`,
		},
		{
			tag:   "[!TIP]",
			name:  "Tip",
			class: "markdown-alert-tip",
			icon:  `<svg class="octicon octicon-light-bulb" viewBox="0 0 16 16" width="16" height="16" aria-hidden="true"><path d="M8 1.5c-2.363 0-4 1.69-4 3.75 0 .984.424 1.625.984 2.304l.214.253c.223.264.47.556.673.848.284.411.537.896.621 1.49a.75.75 0 0 1-1.484.211c-.04-.282-.163-.547-.37-.847a8.456 8.456 0 0 0-.542-.68c-.09-.107-.18-.214-.27-.323C3.12 7.64 2.5 6.64 2.5 5.25 2.5 2.35 4.863 0 8 0s5.5 2.35 5.5 5.25c0 1.39-.62 2.39-1.326 3.235-.09.11-.18.216-.27.324-.177.213-.357.43-.542.68-.207.3-.33.565-.37.847a.75.75 0 0 1-1.485-.212c.084-.593.337-1.078.621-1.489.203-.292.45-.584.673-.848.075-.088.147-.173.213-.253.561-.679.985-1.32.985-2.304 0-2.06-1.637-3.75-4-3.75ZM5.75 12h4.5a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1 0-1.5Zm1 3h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1 0-1.5Z"></path></svg>`,
		},
		{
			tag:   "[!IMPORTANT]",
			name:  "Important",
			class: "markdown-alert-important",
			icon:  `<svg class="octicon octicon-report" viewBox="0 0 16 16" width="16" height="16" aria-hidden="true"><path d="M0 1.75C0 .784.784 0 1.75 0h12.5C15.216 0 16 .784 16 1.75v9.5A1.75 1.75 0 0 1 14.25 13H9.06l-2.573 2.573A1.458 1.458 0 0 1 4 14.543V13H1.75A1.75 1.75 0 0 1 0 11.25Zm1.75-.25a.25.25 0 0 0-.25.25v9.5c0 .138.112.25.25.25h2.5a.75.75 0 0 1 .75.75v2.19l2.72-2.72a.749.749 0 0 1 .53-.22h6a.25.25 0 0 0 .25-.25v-9.5a.25.25 0 0 0-.25-.25Zm7 2a.75.75 0 0 1 .75.75v3.5a.75.75 0 0 1-1.5 0v-3.5A.75.75 0 0 1 8.75 3.5ZM8 10a1 1 0 1 1 0-2 1 1 0 0 1 0 2Z"></path></svg>`,
		},
		{
			tag:   "[!WARNING]",
			name:  "Warning",
			class: "markdown-alert-warning",
			icon:  `<svg class="octicon octicon-alert" viewBox="0 0 16 16" width="16" height="16" aria-hidden="true"><path d="M6.457 1.047c.659-1.234 2.427-1.234 3.086 0l6.082 11.39A1.75 1.75 0 0 1 14.082 15H1.918A1.75 1.75 0 0 1 .377 12.437Zm1.629 1.549a.25.25 0 0 0-.44 0L1.565 13.985a.25.25 0 0 0 .22.365h12.164a.25.25 0 0 0 .22-.365ZM8 5c.552 0 1 .448 1 1v2c0 .552-.448 1-1 1s-1-.448-1-1V6c0-.552.448-1 1-1Zm0 7a1 1 0 1 0 0-2 1 1 0 0 0 0 2Z"></path></svg>`,
		},
		{
			tag:   "[!CAUTION]",
			name:  "Caution",
			class: "markdown-alert-caution",
			icon:  `<svg class="octicon octicon-stop" viewBox="0 0 16 16" width="16" height="16" aria-hidden="true"><path d="M4.47.22A.749.749 0 0 1 5 0h6c.199 0 .389.079.53.22l4.25 4.25c.141.14.22.331.22.53v6a.749.749 0 0 1-.22.53l-4.25 4.25A.749.749 0 0 1 11 16H5a.749.749 0 0 1-.53-.22L.22 11.53A.749.749 0 0 1 0 11V5c0-.199.079-.389.22-.53Zm.84 1.28L1.5 5.31v5.38l3.81 3.81h5.38l3.81-3.81V5.31L10.69 1.5ZM8 4a.75.75 0 0 1 .75.75v3.5a.75.75 0 0 1-1.5 0v-3.5A.75.75 0 0 1 8 4Zm0 8a1 1 0 1 1 0-2 1 1 0 0 1 0 2Z"></path></svg>`,
		},
	}

	for _, alert := range alertTypes {
		if strings.HasPrefix(strings.ToUpper(trimmedAfterP), alert.tag) {
			restOfP := trimmedAfterP[len(alert.tag):]
			restOfP = strings.TrimPrefix(restOfP, "<br />")
			restOfP = strings.TrimPrefix(restOfP, "<br>")
			restOfP = strings.TrimPrefix(restOfP, "<br/>")
			restOfP = strings.TrimSpace(restOfP)

			var bodyBuilder strings.Builder
			if restOfP != "" && restOfP != "</p>" {
				if !strings.HasSuffix(restOfP, "</p>") {
					bodyBuilder.WriteString("<p>")
					bodyBuilder.WriteString(restOfP)
					bodyBuilder.WriteString("</p>")
				} else {
					bodyBuilder.WriteString("<p>")
					bodyBuilder.WriteString(restOfP)
				}
			}

			firstPEnd := strings.Index(trimmedInner, "</p>")
			if firstPEnd != -1 && firstPEnd+4 < len(trimmedInner) {
				remainingContent := strings.TrimSpace(trimmedInner[firstPEnd+4:])
				if remainingContent != "" {
					if bodyBuilder.Len() > 0 {
						bodyBuilder.WriteString("\n")
					}
					bodyBuilder.WriteString(remainingContent)
				}
			}

			body := bodyBuilder.String()
			if body != "" {
				return fmt.Sprintf(`<div class="markdown-alert %s"><p class="markdown-alert-title">%s%s</p>%s</div>`,
					alert.class, alert.icon, alert.name, body)
			}
			return fmt.Sprintf(`<div class="markdown-alert %s"><p class="markdown-alert-title">%s%s</p></div>`,
				alert.class, alert.icon, alert.name)
		}
	}

	return block
}

var _ ports.MarkdownRendererPort = (*GoldmarkRenderer)(nil)
