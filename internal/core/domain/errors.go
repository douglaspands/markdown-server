package domain

import "errors"

var (
	// ErrFileNotFound indica que o arquivo solicitado não foi localizado no diretório raiz.
	ErrFileNotFound = errors.New("arquivo não encontrado")

	// ErrPathEscape indica uma tentativa de Directory Traversal ou acesso fora da raiz permitida.
	ErrPathEscape = errors.New("acesso negado: caminho fora do diretório raiz permitido")

	// ErrInvalidMarkdown indica que o conteúdo não pôde ser processado como Markdown.
	ErrInvalidMarkdown = errors.New("conteúdo markdown inválido")

	// ErrDirectoryScan indica uma falha durante a varredura recursiva do diretório raiz.
	ErrDirectoryScan = errors.New("falha ao escanear diretório")
)
