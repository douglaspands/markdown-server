package browser

import (
	"context"
	"errors"
	"os/exec"
	"runtime"
	"strings"

	"github.com/douglas/markdown-server/internal/core/ports"
)

// CommandExecutor define a função para disparo de comandos no sistema operacional (para facilitar testes e mocks).
type CommandExecutor func(ctx context.Context, name string, args ...string) error

// BrowserLauncher implementa a porta ports.BrowserLauncherPort para abertura de navegadores em Windows, Linux e macOS.
type BrowserLauncher struct {
	executor CommandExecutor
	goos     string
}

// NewBrowserLauncher constrói o inicializador com o executor padrão do sistema.
func NewBrowserLauncher() *BrowserLauncher {
	return &BrowserLauncher{
		goos: runtime.GOOS,
		executor: func(ctx context.Context, name string, args ...string) error {
			cmd := exec.CommandContext(ctx, name, args...)
			return cmd.Start()
		},
	}
}

// NewCustomBrowserLauncher permite injetar um executor customizado e SO específico para testes unitários.
func NewCustomBrowserLauncher(goos string, executor CommandExecutor) *BrowserLauncher {
	if goos == "" {
		goos = runtime.GOOS
	}
	return &BrowserLauncher{
		goos:     goos,
		executor: executor,
	}
}

// OpenURL abre a URL especificada no navegador padrão de acordo com o sistema operacional.
func (b *BrowserLauncher) OpenURL(ctx context.Context, url string) error {
	if ctx != nil && ctx.Err() != nil {
		return ctx.Err()
	}

	trimmedURL := strings.TrimSpace(url)
	if trimmedURL == "" {
		return errors.New("url não pode ser vazia")
	}

	var name string
	var args []string

	switch b.goos {
	case "windows":
		name = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", trimmedURL}
	case "darwin":
		name = "open"
		args = []string{trimmedURL}
	default: // linux, freebsd, etc.
		name = "xdg-open"
		args = []string{trimmedURL}
	}

	if b.executor == nil {
		return errors.New("nenhum executor de comando configurado")
	}

	return b.executor(ctx, name, args...)
}

var _ ports.BrowserLauncherPort = (*BrowserLauncher)(nil)
