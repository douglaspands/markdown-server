package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"github.com/douglas/markdown-server/internal/adapters/in/cli"
	adapterhttp "github.com/douglas/markdown-server/internal/adapters/in/http"
	"github.com/douglas/markdown-server/internal/adapters/out/browser"
	"github.com/douglas/markdown-server/internal/adapters/out/fs"
	"github.com/douglas/markdown-server/internal/adapters/out/renderer"
	"github.com/douglas/markdown-server/internal/core/domain"
	"github.com/douglas/markdown-server/internal/core/services"
	"github.com/douglas/markdown-server/internal/version"
)

func main() {
	cli.Execute(runServer)
}

func runServer(cmd *cobra.Command, args []string, cfg domain.ServerConfig) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	return runServerWithContext(ctx, cmd, args, cfg)
}

func runServerWithContext(ctx context.Context, cmd *cobra.Command, args []string, cfg domain.ServerConfig) error {
	startTime := time.Now()

	// 1. Instanciação dos Adaptadores de Saída
	scanner, err := fs.NewFSScanner(cfg.RootDir)
	if err != nil {
		return fmt.Errorf("falha ao inicializar scanner de arquivos em %q: %w", cfg.RootDir, err)
	}

	goldmarkRenderer := renderer.NewGoldmarkRenderer()
	browserLauncher := browser.NewBrowserLauncher()

	// 2. Instanciação dos Serviços de Domínio
	mdService := services.NewMarkdownService(scanner, goldmarkRenderer)
	healthService := services.NewHealthCheckService(startTime)

	// 3. Instanciação do Adaptador de Entrada HTTP
	server, err := adapterhttp.NewServer(cfg, mdService, healthService)
	if err != nil {
		return fmt.Errorf("falha ao instanciar servidor HTTP: %w", err)
	}

	localURL := server.LocalURL()
	lanURL := server.LANURL()
	fmt.Println("=================================================================")
	fmt.Println(" " + version.Formatted("md_server"))
	fmt.Println("=================================================================")
	fmt.Printf(" Servindo diretório : %s\n", cfg.RootDir)
	fmt.Printf(" Local              : %s\n", localURL)
	fmt.Printf(" Rede Local (LAN)   : %s\n", lanURL)
	fmt.Println(" Pressione Ctrl+C para encerrar o servidor.")
	fmt.Println("=================================================================")

	// 4. Inicialização do Servidor HTTP em background
	serverErrChan := make(chan error, 1)
	go func() {
		if err := server.Start(ctx); err != nil {
			serverErrChan <- err
		}
	}()

	// 5. Abertura Automática do Navegador (Suporte a Duplo Clique no Windows/Linux)
	if cfg.AutoOpenBrowser {
		go func() {
			_ = browserLauncher.OpenURL(ctx, localURL)
		}()
	}

	select {
	case <-ctx.Done():
		fmt.Println("\nSinal de encerramento recebido. Desligando servidor graciosamente...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return server.Stop(shutdownCtx)
	case err := <-serverErrChan:
		return err
	}
}
