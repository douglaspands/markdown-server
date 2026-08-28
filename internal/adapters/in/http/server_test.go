package http_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	adapterhttp "github.com/douglas/markdown-server/internal/adapters/in/http"
	"github.com/douglas/markdown-server/internal/core/domain"
	"github.com/douglas/markdown-server/internal/core/services"
	"github.com/douglas/markdown-server/internal/testutils"
)

func setupTestServer(t *testing.T) (*adapterhttp.Server, *testutils.MockFileScanner, *testutils.MockMarkdownRenderer) {
	t.Helper()
	scanner := &testutils.MockFileScanner{
		SanitizePathFunc: func(p string) (string, error) {
			if p == "escape" || p == "../../escape" {
				return "", domain.ErrPathEscape
			}
			return p, nil
		},
		ReadFileFunc: func(ctx context.Context, p string) ([]byte, error) {
			if p == "README.md" || p == "doc.md" {
				return []byte("# Titulo\nConteudo"), nil
			}
			return nil, domain.ErrFileNotFound
		},
		ScanDirectoryFunc: func(ctx context.Context) (*domain.FileNode, error) {
			return &domain.FileNode{
				Name: "root",
				Type: domain.NodeTypeDirectory,
				Children: []*domain.FileNode{
					{Name: "doc.md", RelativePath: "doc.md", Type: domain.NodeTypeFile, IsMarkdown: true},
				},
			}, nil
		},
	}

	renderer := &testutils.MockMarkdownRenderer{
		RenderHTMLFunc: func(ctx context.Context, md []byte, p string) (string, string, error) {
			return "<h1>Titulo</h1><p>Conteudo</p>", "Titulo", nil
		},
	}

	mdService := services.NewMarkdownService(scanner, renderer)
	healthService := services.NewHealthCheckService(time.Now())

	cfg := domain.ServerConfig{
		RootDir:         ".",
		Port:            8080,
		AutoOpenBrowser: false,
	}

	server, err := adapterhttp.NewServer(cfg, mdService, healthService)
	require.NoError(t, err)

	return server, scanner, renderer
}

func TestHTTPServer(t *testing.T) {
	server, _, _ := setupTestServer(t)
	handler := server.Handler()

	t.Run("Given GET /api/health When requested Then it returns 200 OK with health JSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Header().Get("Content-Type"), "application/json")

		var status domain.HealthStatus
		err := json.Unmarshal(rec.Body.Bytes(), &status)
		require.NoError(t, err)
		assert.Equal(t, "UP", status.Status)
	})

	t.Run("Given GET /static/css/style.css When requested Then it returns 200 and CSS content", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/static/css/style.css", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "--bg-primary")
	})

	t.Run("Given GET /favicon.ico When requested Then it returns 200 and SVG icon bytes", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/favicon.ico", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "image/svg+xml", rec.Header().Get("Content-Type"))
		assert.Contains(t, rec.Body.String(), "<svg")
	})

	t.Run("Given GET / with default ExposeLAN=false When requested Then QR Code button and modal are hidden", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Header().Get("Content-Type"), "text/html")
		assert.Contains(t, rec.Body.String(), "Titulo")
		assert.Contains(t, rec.Body.String(), "Markdown Viewer")
		assert.Contains(t, rec.Body.String(), `id="sidebar-resizer"`)
		assert.NotContains(t, rec.Body.String(), `id="qrcode-btn"`)
		assert.NotContains(t, rec.Body.String(), `id="qrcode-modal"`)
	})

	t.Run("Given server with ExposeLAN=true When requested Then QR Code button and modal are rendered", func(t *testing.T) {
		lanCfg := domain.ServerConfig{
			RootDir:         ".",
			Port:            0,
			AutoOpenBrowser: false,
			ExposeLAN:       true,
		}
		lanServer, err := adapterhttp.NewServer(lanCfg, services.NewMarkdownService(&testutils.MockFileScanner{
			SanitizePathFunc: func(p string) (string, error) { return p, nil },
			ReadFileFunc: func(ctx context.Context, p string) ([]byte, error) {
				return []byte("# Doc"), nil
			},
			ScanDirectoryFunc: func(ctx context.Context) (*domain.FileNode, error) {
				return &domain.FileNode{Name: "root", Type: domain.NodeTypeDirectory}, nil
			},
		}, &testutils.MockMarkdownRenderer{
			RenderHTMLFunc: func(ctx context.Context, md []byte, p string) (string, string, error) {
				return "<h1>Doc</h1>", "Doc", nil
			},
		}), services.NewHealthCheckService(time.Now()))
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		lanServer.Handler().ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `id="qrcode-btn"`)
		assert.Contains(t, rec.Body.String(), `id="qrcode-modal"`)
		assert.Contains(t, rec.Body.String(), `data-lan-url=`)
	})

	t.Run("Given GET /missing.md When requested Then it returns 404 with error template", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/missing.md", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Contains(t, rec.Body.String(), "404")
		assert.Contains(t, rec.Body.String(), "Documento Não Encontrado")
		assert.Contains(t, rec.Body.String(), `id="sidebar-resizer"`)
	})

	t.Run("Given GET with path escape attempt When requested Then it returns 403 Forbidden", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/escape", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusForbidden, rec.Code)
	})

	t.Run("Given a server When Start and Stop are invoked Then lifecycle completes smoothly and exposes URLs", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err := server.Start(ctx)
		require.NoError(t, err)
		assert.NotEmpty(t, server.URL())
		assert.NotEmpty(t, server.LocalURL())
		assert.NotEmpty(t, server.LANURL())
		assert.Empty(t, server.LANIP()) // ExposeLAN is false by default
		assert.True(t, server.Port() > 0)

		err = server.Stop(context.Background())
		require.NoError(t, err)
	})
}
