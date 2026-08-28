package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/douglas/markdown-server/internal/core/domain"
)

// PageData estrutura os dados passados para renderização dos templates HTML.
type PageData struct {
	Document *domain.MarkdownDocument
	FileTree *domain.FileNode
	Error    string
	LocalURL string
	LANURL   string
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	status, err := s.healthService.CheckHealth(r.Context())
	if err != nil {
		http.Error(w, `{"status":"ERROR"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(status)
}

func (s *Server) handleDocument(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqPath := r.URL.Path

	fileTree, _ := s.mdService.GetFileTree(ctx)
	localURL := s.LocalURL()
	lanURL := s.LANURL()

	doc, err := s.mdService.GetDocument(ctx, reqPath)
	if err != nil {
		if errors.Is(err, domain.ErrPathEscape) {
			http.Error(w, "403 Forbidden: Acesso Negado", http.StatusForbidden)
			return
		}

		// Se o arquivo não existir, renderiza a página amigável 404
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		_ = s.err404Tmpl.Execute(w, PageData{
			FileTree: fileTree,
			Error:    "Documento não encontrado",
			LocalURL: localURL,
			LANURL:   lanURL,
		})
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_ = s.tmpl.Execute(w, PageData{
		Document: doc,
		FileTree: fileTree,
		LocalURL: localURL,
		LANURL:   lanURL,
	})
}
