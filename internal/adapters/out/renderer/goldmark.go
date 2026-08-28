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

// NewGoldmarkRenderer cria uma nova instância configurada com GFM, tabelas, autolinks e cabeçalhos automáticos.
func NewGoldmarkRenderer() *GoldmarkRenderer {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
			extension.Strikethrough,
			extension.Linkify,
			extension.TaskList,
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

// RenderHTML processa o Markdown, executa o AST, reescreve links relativos, destaca sintaxe e formata blocos Mermaid.
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

var _ ports.MarkdownRendererPort = (*GoldmarkRenderer)(nil)
