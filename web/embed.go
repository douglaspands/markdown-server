package web

import "embed"

// EmbeddedFS contém todos os assets estáticos e templates HTML empacotados diretamente no binário.
//
//go:embed templates/* static/*
var EmbeddedFS embed.FS
