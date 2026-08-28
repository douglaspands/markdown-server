package http

import (
	"context"
	"fmt"
	"html/template"
	"io/fs"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/douglas/markdown-server/internal/core/domain"
	"github.com/douglas/markdown-server/internal/core/ports"
	"github.com/douglas/markdown-server/web"
)

// Server gerencia o servidor HTTP local servindo o Markdown e a API REST.
type Server struct {
	config        domain.ServerConfig
	mdService     ports.MarkdownServicePort
	healthService ports.HealthServicePort
	httpServer    *http.Server
	listener      net.Listener
	tmpl          *template.Template
	err404Tmpl    *template.Template
	router        *http.ServeMux
	lanIP         string
	mu            sync.RWMutex
}

// NewServer inicializa o servidor HTTP, compila templates embutidos, aloca a porta TCP e configura rotas.
func NewServer(cfg domain.ServerConfig, mdService ports.MarkdownServicePort, healthService ports.HealthServicePort) (*Server, error) {
	funcMap := template.FuncMap{
		"unescapeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	tmpl, err := template.New("base.html").Funcs(funcMap).ParseFS(web.EmbeddedFS, "templates/base.html", "templates/document.html")
	if err != nil {
		return nil, fmt.Errorf("falha ao compilar templates principais: %w", err)
	}

	err404Tmpl, err := template.New("base.html").Funcs(funcMap).ParseFS(web.EmbeddedFS, "templates/base.html", "templates/404.html")
	if err != nil {
		return nil, fmt.Errorf("falha ao compilar template 404: %w", err)
	}

	var bindHost string
	var lanIP string
	if cfg.ExposeLAN {
		bindHost = "0.0.0.0"
		lanIP = DetectLANIP()
	} else {
		bindHost = "127.0.0.1"
		lanIP = ""
	}

	addr := fmt.Sprintf("%s:%d", bindHost, cfg.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		// Se a porta estiver ocupada, tenta alocar em qualquer porta disponível na mesma interface
		listener, err = net.Listen("tcp", fmt.Sprintf("%s:0", bindHost))
		if err != nil {
			return nil, fmt.Errorf("não foi possível abrir porta TCP: %w", err)
		}
	}

	allocatedPort := listener.Addr().(*net.TCPAddr).Port
	cfg.Port = allocatedPort

	router := http.NewServeMux()

	httpSrv := &http.Server{
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	s := &Server{
		config:        cfg,
		mdService:     mdService,
		healthService: healthService,
		httpServer:    httpSrv,
		listener:      listener,
		tmpl:          tmpl,
		err404Tmpl:    err404Tmpl,
		router:        router,
		lanIP:         lanIP,
	}

	s.setupRoutes()

	return s, nil
}

func (s *Server) setupRoutes() {
	// 1. Rota de Health Check
	s.router.HandleFunc("GET /api/health", s.handleHealth)

	// 2. Rota de Favicon
	s.router.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		iconBytes, err := web.EmbeddedFS.ReadFile("static/img/icon.svg")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Header().Set("Cache-Control", "public, max-age=86400")
		_, _ = w.Write(iconBytes)
	})

	// 3. Rota de Assets Estáticos
	staticFS, err := fs.Sub(web.EmbeddedFS, "static")
	if err == nil {
		s.router.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))
	}

	// 4. Rota catch-all para visualização de documentos Markdown
	s.router.HandleFunc("GET /", s.handleDocument)
}

// Handler expõe o roteador HTTP configurado para testes de integração.
func (s *Server) Handler() http.Handler {
	return s.router
}

// Start inicia a escuta HTTP na porta vinculada.
func (s *Server) Start(ctx context.Context) error {
	errChan := make(chan error, 1)
	go func() {
		if err := s.httpServer.Serve(s.listener); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		return s.Stop(context.Background())
	case err := <-errChan:
		return err
	default:
		return nil
	}
}

// Stop encerra o servidor de forma graciosa.
func (s *Server) Stop(ctx context.Context) error {
	if s.httpServer != nil {
		shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		return s.httpServer.Shutdown(shutdownCtx)
	}
	return nil
}

// URL retorna a URL completa onde o servidor está atendendo no host local.
func (s *Server) URL() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return fmt.Sprintf("http://127.0.0.1:%d", s.config.Port)
}

// LocalURL retorna a URL de acesso local via localhost.
func (s *Server) LocalURL() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return fmt.Sprintf("http://localhost:%d", s.config.Port)
}

// LANURL retorna a URL completa acessível por outros dispositivos na rede local (LAN).
func (s *Server) LANURL() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.lanIP == "" {
		return s.LocalURL()
	}
	return fmt.Sprintf("http://%s:%d", s.lanIP, s.config.Port)
}

// LANIP retorna o endereço IP da rede local detectado para o servidor.
func (s *Server) LANIP() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.lanIP
}

// Port retorna a porta alocada pelo servidor.
func (s *Server) Port() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.config.Port
}
