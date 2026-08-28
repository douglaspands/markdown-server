package domain

// MarkdownDocument representa um documento Markdown lido, processado e estruturado para visualização.
type MarkdownDocument struct {
	// RelativePath é o caminho relativo do arquivo a partir da raiz (ex: "docs/architecture.md").
	RelativePath string `json:"relative_path"`

	// Title é o título do documento (extraído do primeiro cabeçalho H1 ou do nome do arquivo).
	Title string `json:"title"`

	// RawContent é o conteúdo original em bytes do Markdown.
	RawContent []byte `json:"-"`

	// HTMLContent é o HTML resultante do processamento do Markdown, com highlight e links reescritos.
	HTMLContent string `json:"html_content"`

	// SizeBytes é o tamanho do arquivo em bytes.
	SizeBytes int64 `json:"size_bytes"`
}
